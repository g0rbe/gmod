package gmod

import "math/rand"

// RandomString generates a random string from chars with length l.
// If chars nil/empty or l < 1, returns an empty string ("").
//
// RandomString is not cryptographically secure, because uses math/rand!
func RandomString(chars []byte, l int) string {

	if len(chars) == 0 || l < 1 {
		return ""
	}

	r := make([]byte, 0, l)

	for i := 0; i < l; i++ {
		r = append(r, chars[rand.Intn(len(chars))])
	}

	return string(r)
}
