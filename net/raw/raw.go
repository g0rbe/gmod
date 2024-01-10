package raw

import (
	"encoding/binary"
)

func Htons(i uint16) uint16 {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, i)
	return uint16(b[0])<<8 | uint16(b[1])
}
