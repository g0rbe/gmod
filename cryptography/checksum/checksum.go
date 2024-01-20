// Compute and check SHA256/SHA512 message digest
package checksum

import (
	"crypto/sha256"
	"crypto/sha512"
	"io"
	"os"
)

// Data256 returns the SHA256 checksum of data.
func Data256(data []byte) []byte {

	s := sha256.Sum256(data)
	return s[:]
}

// Data512 returns the SHA512 checksum of data.
func Data512(data []byte) []byte {

	s := sha512.Sum512(data)
	return s[:]
}

// Path256 returns the SHA256 checksum of the file in path.
// Uses io.Copy() to be usable on large files.
func Path256(path string) ([]byte, error) {

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

// Path512 returns the SHA512 checksum of the file in path.
func Path512(path string) ([]byte, error) {

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
