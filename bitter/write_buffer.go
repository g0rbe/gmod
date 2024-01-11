package bitter

import (
	"crypto/rand"
	"time"
)

// WriteBytes appends bytes at the end of b.
func (b *Buffer) WriteBytes(bytes ...byte) {
	b.b = append(b.b, bytes...)
}

// WriteUint8 appends v at the end of b.
func (b *Buffer) WriteUint8(v uint8) {
	b.WriteBytes(byte(v))
}

// WriteInt8 appends v at the end of b.
func (b *Buffer) WriteInt8(v int8) {
	b.WriteBytes(byte(v))
}

// WriteUint16 appends v at the end of b.
func (b *Buffer) WriteUint16(v uint16) {
	b.WriteBytes(byte(v>>8), byte(v))
}

// WriteInt16 appends v at the end of b.
func (b *Buffer) WriteInt16(v int16) {
	b.WriteBytes(byte(v>>8), byte(v))
}

// WriteUint24 appends v at the end of b.
func (b *Buffer) WriteUint24(v uint32) {
	b.WriteBytes(byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt24 appends v at the end of b.
func (b *Buffer) WriteInt24(v int32) {
	b.WriteBytes(byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint32 appends v at the end of b.
func (b *Buffer) WriteUint32(v uint32) {
	b.WriteBytes(byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt32 appends v at the end of b.
func (b *Buffer) WriteInt32(v int32) {
	b.WriteBytes(byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint64 appends v at the end of b.
func (b *Buffer) WriteUint64(v uint64) {
	b.WriteBytes(byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt64 appends v at the end of b.
func (b *Buffer) WriteInt64(v int64) {
	b.WriteBytes(byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt append v at the end of b with size bitSize.
// The number of bytes depends on IntSize.
func (b *Buffer) WriteInt(v int) {
	switch IntSize {
	case 32:
		b.WriteUint32(uint32(v))
	case 64:
		b.WriteUint64(uint64(v))
	default:
		panic("invalid IntSize")
	}
}

// WriteGMTUnixTime32 appends time to b in 32 bit unix format.
func (b *Buffer) WriteGMTUnixTime32(time time.Time) {
	b.WriteUint32(uint32(time.Unix()))
}

// WriteRandom appends n random bytes to b.
// If failed to read random, this function panics.
// This function gets random numbers from `crypto/rand` package.
func (b *Buffer) WriteRandom(n int) {
	v := make([]byte, n)

	_, err := rand.Read(v)
	if err != nil {
		panic("failed to read rand: " + err.Error())
	}

	b.WriteBytes(v...)
}

// WriteVector appends the length of bytes then the bytes itself.
// The length type is depend on bitSize (eg.: uint8, uint16, uint24, uint32, uint64).
// Therefore, bitSize must be 8/16/24/32/64.
// If bitSize is an invalid number, this function panics.
func (b *Buffer) WriteVector(v []byte, bitSize int) {

	switch bitSize {
	case 8:
		b.WriteUint8(uint8(len(v)))
	case 16:
		b.WriteUint16(uint16(len(v)))
	case 24:
		b.WriteUint24(uint32(len(v)))
	case 32:
		b.WriteUint32(uint32(len(v)))
	case 64:
		b.WriteUint64(uint64(len(v)))
	default:
		panic("invalid bitSize value")
	}

	b.WriteBytes(v...)
}
