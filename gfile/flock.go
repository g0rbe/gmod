package gfile

import (
	"syscall"
)

// FlockGetRead returns whether read lock (F_RDLCK) could be placed on fd.
func FlockGetRead(fd uintptr) (bool, error) {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_RDLCK
	flock.Whence = 0 // SEEK_SET, start of the file
	flock.Start = 0  // Starting offset
	flock.Len = 0    // Lock all bytes

	err := syscall.FcntlFlock(fd, syscall.F_GETLK, flock)

	return flock.Type == syscall.F_UNLCK, err
}

// FlockGetRead returns whether write lock (F_WRLCK) could be placed on fd.
func FlockGetWrite(fd uintptr) (bool, error) {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_WRLCK
	flock.Whence = 0 // SEEK_SET, start of the file
	flock.Start = 0  // Starting offset
	flock.Len = 0    // Lock all bytes

	err := syscall.FcntlFlock(fd, syscall.F_GETLK, flock)

	return flock.Type == syscall.F_UNLCK, err
}

// FlockRead acquire the advisory record lock for read on fd.
// This function returns error if a conflicting lock is held by another process.
// In order to place a read lock, fd must be open for reading.
func FlockRead(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_RDLCK
	flock.Whence = 0 // SEEK_SET, start of the file
	flock.Start = 0  // Starting offset
	flock.Len = 0    // Lock all bytes

	return syscall.FcntlFlock(fd, syscall.F_SETLK, flock)
}

// FlockWrite acquire the advisory record lock for write on fd.
// This function returns error if a conflicting lock is held by another process.
// In order to place a write lock, fd must be open for writing.
func FlockWrite(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_WRLCK
	flock.Whence = 0 // SEEK_SET, start of the file
	flock.Start = 0  // Starting offset
	flock.Len = 0    // Lock all bytes

	return syscall.FcntlFlock(fd, syscall.F_SETLK, flock)
}

// FlockUnlock release the advisory record lock on fd.
// This function returns error if a conflicting lock is held by another process.
func FlockUnlock(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_UNLCK
	flock.Whence = 0 // SEEK_SET, start of the file
	flock.Start = 0  // Starting offset
	flock.Len = 0    // Lock all bytes

	return syscall.FcntlFlock(fd, syscall.F_SETLK, flock)
}

// FlockReadWait acquire the advisory record lock for read on fd.
// If a conflicting lock is held on the file, then wait for that lock to be released.
// In order to place a read lock, fd must be open for reading.
func FlockReadWait(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_RDLCK
	flock.Whence = 0 // SEEK_SET, start of the file
	flock.Start = 0  // Starting offset
	flock.Len = 0    // Lock all bytes

	return syscall.FcntlFlock(fd, syscall.F_SETLKW, flock)
}

// FlockWriteWait acquire the advisory record lock for write on fd.
// If a conflicting lock is held on the file, then wait for that lock to be released.
// In order to place a write lock, fd must be open for writing.
func FlockWriteWait(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_WRLCK
	flock.Whence = 0 // SEEK_SET, start of the file
	flock.Start = 0  // Starting offset
	flock.Len = 0    // Lock all bytes

	return syscall.FcntlFlock(fd, syscall.F_SETLKW, flock)
}

// FlockUnlockWait release the advisory record lock on fd.
// If a conflicting lock is held on the file, then wait for that lock to be released.
func FlockUnlockWait(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_UNLCK
	flock.Whence = 0 // SEEK_SET, start of the file
	flock.Start = 0  // Starting offset
	flock.Len = 0    // Lock all bytes

	return syscall.FcntlFlock(fd, syscall.F_SETLKW, flock)
}
