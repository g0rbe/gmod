package octets

// WriteBytes appends bytes at the end of b.
func WriteBytes(b *[]byte, bytes ...byte) {
	*b = append(*b, bytes...)
}

// WriteUint8 appends v at the end of b.
func WriteUint8(b *[]byte, v uint8) {
	WriteBytes(b, byte(v))
}

// WriteInt8 appends v at the end of b.
func WriteInt8(b *[]byte, v int8) {
	WriteBytes(b, byte(v))
}

// WriteUint16 appends v at the end of b.
func WriteUint16(b *[]byte, v uint16) {
	WriteBytes(b, byte(v>>8), byte(v))
}

// WriteInt16 appends v at the end of b.
func WriteInt16(b *[]byte, v int16) {
	WriteBytes(b, byte(v>>8), byte(v))
}

// WriteUint24 appends v at the end of b.
func WriteUint24(b *[]byte, v uint32) {
	WriteBytes(b, byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt24 appends v at the end of b.
func WriteInt24(b *[]byte, v int32) {
	WriteBytes(b, byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint32 appends v at the end of b.
func WriteUint32(b *[]byte, v uint32) {
	WriteBytes(b, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt32 appends v at the end of b.
func WriteInt32(b *[]byte, v int32) {
	WriteBytes(b, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteUint64 appends v at the end of b.
func WriteUint64(b *[]byte, v uint64) {
	WriteBytes(b, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt64 appends v at the end of b.
func WriteInt64(b *[]byte, v int64) {
	WriteBytes(b, byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

// WriteInt append v at the end of b with size bitSize.
// bitSize must be 8, 16, 24, 32 or 64.
func WriteInt(b *[]byte, v int, bitSize int) {
	switch bitSize {
	case 8:
		WriteUint8(b, uint8(v))
	case 16:
		WriteUint16(b, uint16(v))
	case 24:
		WriteUint24(b, uint32(v))
	case 32:
		WriteUint32(b, uint32(v))
	case 64:
		WriteUint64(b, uint64(v))
	default:
		panic("invalid bitSize value")
	}
}
