package bitter

// ReadBytes removes the first n bytes from b and returns it.
// If the read failed, returns nil.
func ReadBytes(b *[]byte, n int) []byte {

	if len(*b) < n || n < 0 {
		return nil
	}

	v := (*b)[:n]
	*b = (*b)[n:]

	return v
}

// Skip removes the first n bytes from b.
// Returns whether it was successful.
func Skip(b *[]byte, n int) bool {
	return ReadBytes(b, n) != nil
}

// ReadUint8 removes the first byte from b and returns it as an uint8.
// The bool indicates whether the read was successful.
func ReadUint8(b *[]byte) (uint8, bool) {

	v := ReadBytes(b, 1)
	if v == nil {
		return 0, false
	}

	return uint8(v[0]), true
}

// ReadInt8 removes the first byte from b and returns it as an int8.
// The bool indicates whether the read was successful.
func ReadInt8(b *[]byte) (int8, bool) {

	v := ReadBytes(b, 1)
	if v == nil {
		return 0, false
	}

	return int8(v[0]), true
}

// ReadUint16 removes the first 2 bytes from b and returns it as an uint16.
// The bool indicates whether the read was successful.
func ReadUint16(b *[]byte) (uint16, bool) {

	v := ReadBytes(b, 2)
	if v == nil {
		return 0, false
	}

	return uint16(v[0])<<8 | uint16(v[1]), true
}

// ReadInt16 removes the first 2 bytes from b and returns it as an int16.
// The bool indicates whether the read was successful.
func ReadInt16(b *[]byte) (int16, bool) {

	v := ReadBytes(b, 2)
	if v == nil {
		return 0, false
	}

	return int16(v[0])<<8 | int16(v[1]), true
}

// ReadUint24 removes the first 3 bytes from b and returns it as a uint32.
// The bool indicates whether the read was successful.
func ReadUint24(b *[]byte) (uint32, bool) {

	v := ReadBytes(b, 3)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<16 | uint32(v[1])<<8 | uint32(v[2]), true
}

// ReadInt24 removes the first 3 bytes from b and returns it as a int32.
// The bool indicates whether the read was successful.
func ReadInt24(b *[]byte) (int32, bool) {

	v := ReadBytes(b, 3)
	if v == nil {
		return 0, false
	}

	return int32(v[0])<<16 | int32(v[1])<<8 | int32(v[2]), true
}

// ReadUint32 removes the first 4 bytes from b and returns it as an uint32.
// The bool indicates whether the read was successful.
func ReadUint32(b *[]byte) (uint32, bool) {

	v := ReadBytes(b, 4)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3]), true
}

// ReadInt32 removes the first 4 bytes from b and returns it as an int32.
// The bool indicates whether the read was successful.
func ReadInt32(b *[]byte) (int32, bool) {

	v := ReadBytes(b, 4)
	if v == nil {
		return 0, false
	}

	return int32(v[0])<<24 | int32(v[1])<<16 | int32(v[2])<<8 | int32(v[3]), true
}

// ReadUint64 removes the first 8 bytes from b and returns it as an uint64.
// The bool indicates whether the read was successful.
func ReadUint64(b *[]byte) (uint64, bool) {

	v := ReadBytes(b, 8)
	if v == nil {
		return 0, false
	}

	return uint64(v[0])<<56 | uint64(v[1])<<48 | uint64(v[2])<<40 | uint64(v[3])<<32 | uint64(v[4])<<24 | uint64(v[5])<<16 | uint64(v[6])<<8 | uint64(v[7]), true
}

// ReadInt64 removes the first 8 bytes from b and returns it as an int64.
// The bool indicates whether the read was successful.
func ReadInt64(b *[]byte) (int64, bool) {

	v := ReadBytes(b, 8)
	if v == nil {
		return 0, false
	}

	return int64(v[0])<<56 | int64(v[1])<<48 | int64(v[2])<<40 | int64(v[3])<<32 | int64(v[4])<<24 | int64(v[5])<<16 | int64(v[6])<<8 | int64(v[7]), true
}

// ReadInt removes the first bytes (depends on bitSize) from b and returns it as an int.
// bitSize must be 8, 16, 24, 32 or 64.
func ReadInt(b *[]byte, bitSize int) (int, bool) {

	switch bitSize {
	case 8:
		l, ok := ReadUint8(b)
		return int(l), ok
	case 16:
		l, ok := ReadUint16(b)
		return int(l), ok
	case 24:
		l, ok := ReadUint24(b)
		return int(l), ok
	case 32:
		l, ok := ReadUint32(b)
		return int(l), ok
	case 64:
		l, ok := ReadUint64(b)
		return int(l), ok
	default:
		panic("invalid bitSize value")
	}
}
