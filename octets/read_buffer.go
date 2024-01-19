package octets

import (
	"time"
)

// ReadBytes removes the first n bytes from b and returns it.
// If the read failed, returns nil.
func (b *Buffer) ReadBytes(n int) []byte {

	if len(b.b) < n || n < 0 {
		return nil
	}

	v := b.b[:n]
	b.b = b.b[n:]

	return v
}

// Skip removes the first n bytes from b.
// Returns whether it was successful.
func (b *Buffer) Skip(n int) bool {
	return b.ReadBytes(n) != nil
}

// ReadUint8 removes the first byte from b and returns it as an uint8.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadUint8() (uint8, bool) {

	v := b.ReadBytes(1)
	if v == nil {
		return 0, false
	}

	return uint8(v[0]), true
}

// ReadInt8 removes the first byte from b and returns it as an int8.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadInt8() (int8, bool) {

	v := b.ReadBytes(1)
	if v == nil {
		return 0, false
	}

	return int8(v[0]), true
}

// ReadUint16 removes the first bytes from b and returns it as an uint16.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadUint16() (uint16, bool) {

	v := b.ReadBytes(2)
	if v == nil {
		return 0, false
	}

	return uint16(v[0])<<8 | uint16(v[1]), true
}

// ReadInt16 removes the first bytes from b and returns it as an int16.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadInt16() (int16, bool) {

	v := b.ReadBytes(2)
	if v == nil {
		return 0, false
	}

	return int16(v[0])<<8 | int16(v[1]), true
}

// ReadUint24 removes the first bytes from b and returns it as a uint32.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadUint24() (uint32, bool) {

	v := b.ReadBytes(3)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<16 | uint32(v[1])<<8 | uint32(v[2]), true
}

// ReadInt24 removes the first bytes from b and returns it as a int32.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadInt24() (int32, bool) {

	v := b.ReadBytes(3)
	if v == nil {
		return 0, false
	}

	return int32(v[0])<<16 | int32(v[1])<<8 | int32(v[2]), true
}

// ReadUint32 removes the first bytes from b and returns it as an uint32.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadUint32() (uint32, bool) {

	v := b.ReadBytes(4)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3]), true
}

// ReadInt32 removes the first bytes from b and returns it as an int32.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadInt32() (int32, bool) {

	v := b.ReadBytes(4)
	if v == nil {
		return 0, false
	}

	return int32(v[0])<<24 | int32(v[1])<<16 | int32(v[2])<<8 | int32(v[3]), true
}

// ReadUint64 removes the first bytes from b and returns it as an uint64.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadUint64() (uint64, bool) {

	v := b.ReadBytes(8)
	if v == nil {
		return 0, false
	}

	return uint64(v[0])<<56 | uint64(v[1])<<48 | uint64(v[2])<<40 | uint64(v[3])<<32 | uint64(v[4])<<24 | uint64(v[5])<<16 | uint64(v[6])<<8 | uint64(v[7]), true
}

// ReadInt64 removes the first bytes from b and returns it as an int64.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadInt64() (int64, bool) {

	v := b.ReadBytes(8)
	if v == nil {
		return 0, false
	}

	return int64(v[0])<<56 | int64(v[1])<<48 | int64(v[2])<<40 | int64(v[3])<<32 | int64(v[4])<<24 | int64(v[5])<<16 | int64(v[6])<<8 | int64(v[7]), true
}

// ReadInt removes the first bytes (depends on IntSize) from b and returns it as an int.
func (b *Buffer) ReadInt() (int, bool) {

	switch IntSize {
	case 32:
		l, ok := b.ReadInt32()
		return int(l), ok
	case 64:
		l, ok := b.ReadInt64()
		return int(l), ok
	default:
		panic("invalid IntSize")
	}
}

// ReadGMTUnixTime32 removes the first bytes from b and returns it as an unix time.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadGMTUnixTime32() (time.Time, bool) {

	v, ok := b.ReadUint32()
	if !ok {
		return time.Time{}, false
	}

	return time.Unix(int64(v), 0), true
}

// ReadVector reads the length of bytes then the bytes itself.
// The length type is depend on bitSize (eg.: uint8, uint16, uint24, uint32, uint64).
// Therefore, bitSize must be 8/16/24/32/64.
// If bitSize is an invalid number, this function panics.
func (b *Buffer) ReadVector(bitSize int) ([]byte, bool) {

	var n int

	switch bitSize {
	case 8:
		n8, ok := b.ReadUint8()
		if !ok {
			return []byte{}, false
		}
		n = int(n8)
	case 16:
		n16, ok := b.ReadUint16()
		if !ok {
			return []byte{}, false
		}
		n = int(n16)
	case 24:
		n24, ok := b.ReadUint24()
		if !ok {
			return []byte{}, false
		}
		n = int(n24)
	case 32:
		n32, ok := b.ReadUint32()
		if !ok {
			return []byte{}, false
		}
		n = int(n32)
	case 64:
		n64, ok := b.ReadUint64()
		if !ok {
			return []byte{}, false
		}
		n = int(n64)
	default:
		panic("invalid bitSize value")
	}

	v := b.ReadBytes(n)
	if v == nil {
		return []byte{}, false
	}

	return v, true
}
