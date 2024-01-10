package dns

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidMaxRetries = errors.New("invalid MaxRetries")
	ErrTruncated         = errors.New("message is truncated")
)

var (
	ErrNoError        = errors.New("NOERROR")  // NOERROR
	ErrFormat         = errors.New("FORMERR")  // FORMERR
	ErrServerFailure  = errors.New("SERVFAIL") // SERVFAIL
	ErrName           = errors.New("NXDOMAIN") // NXDOMAIN
	ErrNotImplemented = errors.New("NOTIMP")   // NOTIMP
	ErrRefused        = errors.New("REFUSED")  // REFUSED
)

// RcodeToError returns the error associated with the rcode.
// If rcode is unknown, returns rcode.
func RcodeToError(rcode int) error {

	switch rcode {
	case 0:
		return ErrNoError
	case 1:
		return ErrFormat
	case 2:
		return ErrServerFailure
	case 3:
		return ErrName
	case 4:
		return ErrNotImplemented
	case 5:
		return ErrRefused
	default:
		return fmt.Errorf("%d", rcode)
	}
}
