package octets

type BitwiseType interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | uint | uintptr | int
}

// BitwiseSet sets y bits of x.
// Returns the modified value.
func BitwiseSet[T BitwiseType](x, y T) T {
	return x | y
}

// BitwiseClear unsets y bits of x.
// Returns the modified value.
func BitwiseClear[T BitwiseType](x, y T) T {
	return x ^ (x & y)
}

// BitwiseToggle toggles y bits of x.
// Returns the modified value.
func BitwiseToggle[T BitwiseType](x, y T) T {
	return x ^ y
}

// BitwiseCheck checks whether y bits is set of x.
// Returns the modified value.
func BitwiseCheck[T BitwiseType](x, y T) bool {
	return x&y == y
}
