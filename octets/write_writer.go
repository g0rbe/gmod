package octets

import "io"

// WriteWriterBytes appends v at the end of in.
func WriteWriterBytes(in io.Writer, bytes ...byte) error {

	_, err := in.Write(bytes)

	return err
}

// WriteWriterUint8 appends v at the end of in.
func WriteWriterUint8(in io.Writer, v uint8) error {
	return WriteWriterBytes(in, byte(v))
}

// WriteWriterUint16 appends v at the end of in.
func WriteWriterUint16(in io.Writer, v uint16) error {
	return WriteWriterBytes(in, byte(v>>8), byte(v))
}

// WriteWriterUint24 appends v at the end of in.
func WriteWriterUint24(in io.Writer, v uint32) error {
	return WriteWriterBytes(in, byte(v>>16), byte(v>>8), byte(v))
}

// WriteWriterUint32 appends v at the end of in.
func WriteWriterUint32(in io.Writer, v uint32) error {
	return WriteWriterBytes(in, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}
