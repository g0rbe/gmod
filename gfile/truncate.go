package gfile

import (
	"fmt"
	"os"
)

// Truncate overwrites the file with shred for ShredIterations times and truncates to 0.
func Truncate(file *os.File) error {

	_, err := shred(file, ShredIterations, true)
	if err != nil {
		return fmt.Errorf("failed to shred: %w", err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("failed to seek: %w", err)
	}

	return file.Truncate(0)
}

// TruncatePath overwrites the file at path with shred for ShredIterations times and truncates to 0.
func TruncatePath(path string) error {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_SYNC, 0)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", path, err)
	}
	defer file.Close()

	_, err = shred(file, ShredIterations, true)
	if err != nil {
		return fmt.Errorf("failed to shred: %w", err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("failed to seek: %w", err)
	}

	return file.Truncate(0)
}
