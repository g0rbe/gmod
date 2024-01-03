package gfile

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// Shred overwrites file in path for n times.
// Writes random data if random is true, else zeroes.
// Removes the file at the end if remove is true.
func Shred(path string, n int, random bool, remove bool) (int64, error) {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_SYNC, 0)
	if err != nil {
		return -1, fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

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

	if remove {
		file.Close()
		err = os.Remove(path)
	}

	return totalWritten, err
}
