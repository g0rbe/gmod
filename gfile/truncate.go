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

	return file.Truncate(0)
}
