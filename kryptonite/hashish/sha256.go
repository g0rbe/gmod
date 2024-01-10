package hashish

import (
	"crypto/sha256"
	"io"
	"os"
)

const (
	Size256 = 32
)

// Data256 returns the SHA256 checksum of data.
func Data256(data []byte) CheckSum {

	s := sha256.Sum256(data)
	return s[:]
}

// File256 returns the SHA256 checksum of the file in path.
// Uses io.Copy() to be usable on large files.
func File256(path string) (CheckSum, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	h := sha256.New()

	_, err = io.Copy(h, file)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// File256Check compare the SHA256 checksum of the file in path with sum.
func File256Check(path string, sum CheckSum) (bool, error) {

	cs, err := File256(path)
	if err != nil {
		return false, err
	}

	return cs.Compare(sum) == 0, nil
}
