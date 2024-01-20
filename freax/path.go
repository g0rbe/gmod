package freax

import (
	"os"
)

func ReadAll(name string) ([]byte, error) {

	file, err := OpenFile(name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return FDReadAll(file.Fd())
}

// Write writes data to the file in path.
// This function truncates the file before writing.
// If path does not exist, create a new file with permission perm (before umask).
func Write(name string, data []byte, perm uint32) error {

	file, err := OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)

	return err
}

// WriteSync writes and syncronise data to the file in path.
// This function truncates the file before writing.
// If path does not exist, create a new file with permission perm (before umask).
func WriteSync(name string, data []byte, perm uint32) error {

	file, err := OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_SYNC, perm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)

	return err
}
