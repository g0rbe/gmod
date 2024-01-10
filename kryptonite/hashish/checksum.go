package hashish

import (
	"bytes"
	"encoding/hex"
	"errors"
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
