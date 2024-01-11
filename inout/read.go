package inout

import "io"

// ReadByte reads a single byte from r.
// EOF is returned as an error.
func ReadByte(r io.Reader) (byte, error) {

	c := make([]byte, 1)

	_, err := r.Read(c)

	return c[0], err
}

// ReadByte reads bytes from r until a newline character ('\n') or EOF.
// EOF is returned as an error.
func ReadLine(r io.Reader) ([]byte, error) {

	var (
		v   []byte
		c   byte
		err error
	)

	for {

		c, err = ReadByte(r)
		if err != nil || c == '\n' {
			break
		}

		v = append(v, c)
	}

	return v, err
}
