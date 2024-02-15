package freax

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"syscall"
)

var (
	ErrEmpty      = errors.New("empty")
	ErrInvalidPid = errors.New("invalid pid")
)

// PIDCreate creates the pid file in path with mode perm (before umask).
// The file in path should not exist.
func PIDCreate(path string, perm uint32) error {

	// Use O_EXCL to make sure that the file not exist
	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, fs.FileMode(perm))
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", path, err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d", os.Getpid()))
	if err != nil {
		return fmt.Errorf("failed to write %s: %w", path, err)
	}

	return nil
}

// PIDRead returns the stored pid in path.
// Remove the newline character ('\n') from the content if exist.
//
// Returns -1 if error occurred.
// Returns ErrEmpty if file is empty.
// Returns ErrInvalidPid if the content is an invalid (eg.: not a number).
func PIDRead(path string) (int, error) {

	out, err := os.ReadFile(path)
	if err != nil {
		return -1, fmt.Errorf("failed to open %s: %w", path, err)
	}

	// Remove newline
	out = bytes.TrimSuffix(out, []byte{'\n'})

	if len(out) == 0 {
		return -1, ErrEmpty
	}

	pid, err := strconv.Atoi(string(out))
	if err != nil {
		return -1, fmt.Errorf("%w: %w", ErrInvalidPid, err)
	}

	return pid, nil
}

// CheckPath returns true if the stored PID in path is running.
// Remove the newline character ('\n') from the content if exist.
//
// Returns ErrEmpty if lock file is empty.
// Returns ErrInvalidPid if the content is an invalid (eg.: not a number).
func PIDCheck(path string) (bool, error) {

	pid, err := PIDRead(path)
	if err != nil {
		return false, err
	}

	err = syscall.Kill(pid, 0)

	return err == nil, err
}

// PIDRemove deletes the pid file in path.
// Simply calls os.Remove(path).
func PIDRemove(path string) error {
	return os.Remove(path)
}
