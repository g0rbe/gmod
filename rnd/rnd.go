package rnd

import (
	"errors"
	mrand "math/rand"

	"github.com/g0rbe/gmod/bitter"
	"golang.org/x/sys/unix"
)

var (
	ErrInvalidSize = errors.New("invalid size")
)

// Bytes returns cryptohraphically secure random bytes with length size.
// This function panics if size <= 0 or failed to get random bytes.
func Bytes(size int) []byte {

	if size <= 0 {
		panic("invalid argument")
	}

	buf := make([]byte, size)

	_, err := unix.Getrandom(buf, 0)
	if err != nil {
		panic(err)
	}

	return buf
}

// Int64 returns a random int64.
// It panics if failed to get random bytes.
func Int64() int64 {

	buf := Bytes(8)

	v, ok := bitter.ReadInt64(&buf)
	if !ok {
		panic("failed to read bytes")
	}

	return v
}

// Int64n return a random int64 int the half-open interval [0,n).
// It panics if n <= 0 or failed to get random bytes.
func Int64n(n int64) int64 {

	if n <= 0 {
		panic("invalid argument")
	}

	v := Int64()
	if v < 0 {
		v = 0 - v
	}

	return v % n
}

// Uint64 returns a random uint64.
// This function panics if failed to get random bytes.
func Uint64() uint64 {

	buf := Bytes(8)

	v, ok := bitter.ReadUint64(&buf)
	if !ok {
		panic("failed to read bytes")
	}

	return v
}

// String generates cryptographically **NOT** secure a random string from chars with length l.
// If chars nil/empty or l < 1, returns an empty string ("").
func String(chars []byte, l int) string {

	if len(chars) == 0 || l < 1 {
		return ""
	}

	r := make([]byte, 0, l)

	for i := 0; i < l; i++ {
		r = append(r, chars[mrand.Intn(len(chars))])
	}

	return string(r)
}
