package octets

import (
	"fmt"
	"io"
)

// ReadReaderBytes reads n byte from in.
// At the end of the file, io.EOF is returned.
func ReadReaderBytes(in io.Reader, n int) ([]byte, error) {

	if n < 0 {
		return []byte{}, fmt.Errorf("number is less than zero")
	}

	b := make([]byte, n)

	_, err := in.Read(b)

	return b, err
}

// ReadReaderSkip skips n bytes in in.
// At the end of the file, io.EOF is returned.
func SkipReader(in io.Reader, n int) error {

	_, err := ReadReaderBytes(in, n)

	return err
}

// ReadReaderUint8 reads a byte from in and returns is as an uint8.
// At the end of the file, io.EOF is returned.
func ReadReaderUint8(in io.Reader) (uint8, error) {

	v, err := ReadReaderBytes(in, 1)

	return uint8(v[0]), err
}

// ReadReaderUint16 reads bytes from in and returns it as an uint16.
// At the end of the file io.EOF returned.
func ReadReaderUint16(in io.Reader) (uint16, error) {

	v, err := ReadReaderBytes(in, 2)
	if err != nil {
		return 0, err
	}

	return uint16(v[0])<<8 | uint16(v[1]), err
}

// ReadReaderUint24 reads bytes from in and returns it as an uint32.
// At the end of the file, io.EOF is returned.
func ReadReaderUint24(in io.Reader) (uint32, error) {

	v, err := ReadReaderBytes(in, 3)
	if err != nil {
		return 0, err
	}

	return uint32(v[0])<<16 | uint32(v[1])<<8 | uint32(v[2]), err
}

// ReadReaderUint32 reads bytes from in and returns it as an uint32.
// At the end of the file io.EOF returned.
func ReadReaderUint32(in io.Reader) (uint32, error) {

	v, err := ReadReaderBytes(in, 4)
	if err != nil {
		return 0, err
	}

	return uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3]), err
}
