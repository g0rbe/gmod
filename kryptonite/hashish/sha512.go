package hashish

import (
	"crypto/sha512"
	"io"
	"os"
)

const (
	Size512 = 64
)

// Data512 returns the SHA512 checksum of data.
func Data512(data []byte) CheckSum {

	s := sha512.Sum512(data)
	return s[:]
}

// File512 returns the SHA512 checksum of the file in path.
func File512(path string) (CheckSum, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	h := sha512.New()

	_, err = io.Copy(h, file)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// File512Check compare the SHA512 checksum of the file in path with sum.
func File512Check(path string, sum CheckSum) (bool, error) {

	cs, err := File512(path)
	if err != nil {
		return false, err
	}

	return cs.Compare(sum) == 0, nil
}
