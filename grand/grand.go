package grand

import (
	"errors"
	"math/rand"

	crand "crypto/rand"
)

var (
	ErrInvalidSize = errors.New("invalid size")
)

// String generates cryptographically **NOT** secure a random string from chars with length l.
// If chars nil/empty or l < 1, returns an empty string ("").
func String(chars []byte, l int) string {

	if len(chars) == 0 || l < 1 {
		return ""
	}

	r := make([]byte, 0, l)

	for i := 0; i < l; i++ {
		r = append(r, chars[rand.Intn(len(chars))])
	}

	return string(r)
}

// Bytes returns cryptohraphically secure random bytes with with length size.
// Returns ErrInvalidSize if size is less than one (size < 1).
func Bytes(size int) ([]byte, error) {

	if size < 1 {
		return nil, ErrInvalidSize
	}

	v := make([]byte, size)

	_, err := crand.Read(v)

	return v, err
}

// MustBytes returns cryptohraphically secure random bytes with with length size.
// This function panics if any failed to get random bytes.
func MustBytes(size int) []byte {

	v, err := Bytes(size)
	if err != nil {
		panic(err)
	}

	return v
}
