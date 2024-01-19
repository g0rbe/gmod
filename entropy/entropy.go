//go:build linux

// Cryptographically secure pseudorandom numbers
package entropy

import (
	"errors"

	"github.com/g0rbe/gmod/octets"
	"golang.org/x/sys/unix"
)

var (
	ErrInvalidSize = errors.New("invalid size")
)

// Flags for Getrandom()
var (
	GRND_NONBLOCK = 0x1
	GRND_RANDOM   = 0x2
	GRND_INSECURE = 0x4
)

// TODO: https://www.zx2c4.com/projects/linux-rng-5.17-5.18/

// Getrandom obtains a series of random bytes by filling buffer buf.
//
// By default Getrandom() blocks if no random bytes are available.
// If the GRND_NONBLOCK flag is set, then Getrandom() does not block in these cases, but instead immediately returns unix.EAGAIN.
//
// If GRND_RANDOM bit is set, then random bytes are drawn from the random source (i.e., the same source as the /dev/random device) instead of the urandom source.
//
// The GRND_INSECURE bit is removes the blocking and allow returning potentially "insecure" random bytes.
//
// See more: man random(4) and random(7)
func Getrandom(buf []byte, flags int) error {

	var (
		l = len(buf)
		i = 0 // Index to count number of bytes read
	)
	// Fill buf until not full
	for i < l {

		n, err := unix.Getrandom(buf[i:], flags)
		if err != nil {
			return err
		}

		// Add the number of read bytes to i index.
		i += n
	}

	return nil
}

// Bytes returns n number cryptographically secure random bytes from /dev/random.
// This function panics if size <= 0 or failed to get random bytes.
func Bytes(n int) []byte {

	if n <= 0 {
		panic("invalid argument")
	}

	buf := make([]byte, n)

	err := Getrandom(buf, 0)
	if err != nil {
		panic(err)
	}

	return buf
}

// Int returns a random int.
// It panics if failed to get random bytes.
func Int() int {

	buf := Bytes(octets.IntSize / 8)

	v, ok := octets.ReadInt(&buf, octets.IntSize)
	if !ok {
		panic("failed to read bytes")
	}

	return v
}

// Intn return a random int in the half-open interval [0,n).
// It panics if n <= 0 or failed to get random bytes.
func Intn(n int) int {

	if n <= 0 {
		panic("invalid argument")
	}

	v := Int()
	if v < 0 {
		v = 0 - v
	}

	return v % n
}

// String generates cryptographically **NOT** secure a random string from chars with length l.
// If chars nil/empty or l < 1, returns an empty string ("").
func String(chars []byte, l int) string {

	if len(chars) == 0 || l < 1 {
		return ""
	}

	r := make([]byte, 0, l)

	for i := 0; i < l; i++ {
		r = append(r, chars[Intn(len(chars))])
	}

	return string(r)
}
