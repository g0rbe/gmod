package inout

import (
	"fmt"
	"strings"
)

// ReadString print prompt p and read a string from stdin.
// Returns the input string without the newline suffix ("\n").
func ReadString(p string) (string, error) {

	_, err := fmt.Print(p)
	if err != nil {
		return "", fmt.Errorf("failed to print prompt: %w", err)
	}

	var v string

	_, err = fmt.Scan(&v)

	return strings.TrimSuffix(v, "\n"), err
}
