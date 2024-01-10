package filio

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

var (
	ShredIterations = 3 // Default iteration number for Shred()
)

// shred overwrite file for n times with random (if random is true) or zeroes.
// Seek to the beginning of file at every iteration.
func shred(file *os.File, n int, random bool) (int64, error) {

	stat, err := file.Stat()
	if err != nil {
		return -1, fmt.Errorf("failed stat: %w", err)
	}

	var totalWritten int64

	for i := 0; i < n; i++ {

		_, err = file.Seek(0, 0)
		if err != nil {
			return -1, fmt.Errorf("failed to seek: %w", err)
		}

		var written int64

		if random {
			written, err = io.CopyN(file, rand.Reader, stat.Size())
		} else {
			written, err = io.CopyN(file, DevZero, stat.Size())
		}

		totalWritten += written

		if err != nil {
			return written, err
		}
	}

	return totalWritten, nil
}

// Shred overwrites file in path for n times.
// Writes random data if random is true, else zeroes.
func Shred(path string, n int, random bool) (int64, error) {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_SYNC, 0)
	if err != nil {
		return -1, fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

	return shred(file, n, random)
}

// ShredLock apply a nonblocking write lock with FlockWrite and overwrites file in path for n times.
// Writes random data if random is true, else zeroes.
//
// This function returns error if other lock is applied on path.
func ShredLock(path string, n int, random bool) (int64, error) {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_SYNC, 0)
	if err != nil {
		return -1, fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

	// Apply a blocking lock.
	err = FlockWrite(file.Fd())
	if err != nil {
		return -1, fmt.Errorf("failed to lock: %w", err)
	}
	// Unlock at return
	defer FlockUnlock(file.Fd())

	return shred(file, n, random)
}

// ShredLock apply a blocking write lock with FlockWriteWait and overwrites file in path for n times.
// Writes random data if random is true, else zeroes.
//
// This function blocks until other lock is applied on path.
func ShredLockWait(path string, n int, random bool) (int64, error) {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_SYNC, 0)
	if err != nil {
		return -1, fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

	// Apply a nonblocking lock.
	err = FlockWriteWait(file.Fd())
	if err != nil {
		return -1, fmt.Errorf("failed to lock: %w", err)
	}
	// Unlock at return
	defer FlockUnlockWait(file.Fd())

	return shred(file, n, random)
}
