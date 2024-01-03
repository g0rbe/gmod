// Compute and check SHA256/SHA512 message digest
package gsha

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

const (
	Size256 = 32
	Size512 = 64
)

var ErrInvalidLength = errors.New("invalid length")

type CheckSum []byte

// FromString converts s to CheckSum.
// Returns ErrInvalidLength if the decoded data has an invalid length.
func FromString(s string) (CheckSum, error) {

	data, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}

	if len(data) != Size256 && len(data) != Size512 {
		return nil, ErrInvalidLength
	}

	return data, nil
}

// Size returns the length of the CheckSum.
func (cs CheckSum) Size() int {
	return len(cs)
}

// String returns the hexadecimal encoded string format of CheckSum.
func (cs CheckSum) String() string {
	return hex.EncodeToString(cs)
}

// Bytes returns the CheckSum as a byte slice.
func (cs CheckSum) Bytes() []byte {
	return []byte(cs)
}

func (cs CheckSum) Compare(b CheckSum) int {

	return bytes.Compare(cs, b)
}

// Data256 returns the SHA256 checksum of data.
func Data256(data []byte) CheckSum {

	s := sha256.Sum256(data)
	return s[:]
}

// File256 returns the SHA256 checksum of file.
func File256(file *os.File) (CheckSum, error) {

	h := sha256.New()

	_, err := io.Copy(h, file)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Path256 returns the SHA256 checksum of the file in path.
func Path256(path string) (CheckSum, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return File256(file)
}

// Check256 compare the SHA256 checksum of the file in path with sum.
func Check256(path string, sum CheckSum) (bool, error) {

	cs, err := Path256(path)
	if err != nil {
		return false, err
	}

	return cs.Compare(sum) == 0, nil
}

// Data512 returns the SHA512 checksum of data.
func Data512(data []byte) CheckSum {

	s := sha512.Sum256(data)
	return s[:]
}

// File512 returns the SHA512 checksum of file.
func File512(file *os.File) (CheckSum, error) {

	h := sha512.New()

	_, err := io.Copy(h, file)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Path512 returns the SHA512 checksum of the file in path.
func Path512(path string) (CheckSum, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return File512(file)
}

// Check512 compare the SHA512 checksum of the file in path with sum.
func Check512(path string, sum CheckSum) (bool, error) {

	cs, err := Path512(path)
	if err != nil {
		return false, err
	}

	return cs.Compare(sum) == 0, nil
}
