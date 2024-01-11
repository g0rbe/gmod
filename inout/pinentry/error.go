package pinentry

import (
	"bytes"
	"fmt"
	"strconv"
)

type Error struct {
	Code    int
	Message string
}

var (
	ErrOpCancelled = NewError(83886179, "Operation cancelled") // ErrOpCancelled returned if "GETPIN", "CONFIRM" or "MESSAGE" is cancelled.
)

func NewError(code int, message string) *Error {
	return &Error{Code: code, Message: message}
}

// ParseError parses line and return an error from it.
func ParseError(line []byte) *Error {

	fields := bytes.SplitN(line, []byte{' '}, 3)
	if len(fields) != 3 {
		return &Error{-1, fmt.Sprintf("invalid line: %s", line)}
	}

	code, err := strconv.Atoi(string(fields[1]))
	if err != nil {
		return &Error{-2, fmt.Sprintf("invalid code: %s", line)}
	}

	return &Error{code, string(fields[2])}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error {

	switch e.Code {
	case ErrOpCancelled.Code:
		return ErrOpCancelled
	default:
		return nil
	}
}
