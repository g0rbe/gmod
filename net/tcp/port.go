package tcp

import "strconv"

type PortTypes interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint | uintptr | string | *string
}

// Port returns whether v is a valid port.
func Port[T PortTypes](v T) bool {

	switch t := any(v).(type) {
	case int8:
		if t >= 0 {
			return true
		}
	case uint8:
		return true
	case int16:
		if t >= 0 {
			return true
		}
	case uint16:
		return true
	case int32:
		if t >= 0 && t <= 65535 {
			return true
		}
	case uint32:
		if t <= 65535 {
			return true
		}
	case int64:
		if t >= 0 && t <= 65535 {
			return true
		}
	case uint64:
		if t <= 65535 {
			return true
		}
	case int:
		if t >= 0 && t <= 65535 {
			return true
		}
	case uint:
		if t <= 65535 {
			return true
		}
	case uintptr:
		if t <= 65535 {
			return true
		}
	case string:
		n, err := strconv.Atoi(t)
		if err != nil {
			return false
		}

		if n >= 0 && n <= 65535 {
			return true
		}
	case *string:
		n, err := strconv.Atoi(*t)
		if err != nil {
			return false
		}

		if n >= 0 && n <= 65535 {
			return true
		}
	}

	return false
}
