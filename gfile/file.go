package gfile

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

// ContainsLine reports whether line is in path file.
// A line in the file does not contains the newline character ("\n").
// This function uses bufio.Scanner to able to handle large files.
func ContainsLine(path, line string) (bool, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == line {
			return true, nil
		}
	}

	return false, scanner.Err()
}

// Contains reports whether substr is in path file.
// A line in the file does not contains the newline character ("\n").
// This function uses bufio.Scanner to able to handle large files.
func Contains(path, substr string) (bool, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == substr {
			return true, nil
		}
	}

	return false, scanner.Err()
}

// CountLines returns the number lines in file.
func CountLines(path string) (int, error) {

	// Source: https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently

	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

// IsExists returns whether path is exists.
// If path is a symbolic link, it is dereferenced.
func IsExists(path string) (bool, error) {

	err := unix.Faccessat(unix.AT_FDCWD, path, syscall.F_OK, 0)
	if err != nil {

		if errors.Is(err, unix.ENOENT) {
			err = nil
		}

		return false, err
	}

	return true, nil

}

// IsDir returns whether path is a directory.
func IsDir(path string) (bool, error) {

	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return stat.IsDir(), nil
}
