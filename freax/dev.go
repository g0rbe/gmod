//go:build linux

package freax

import (
	"fmt"
	"io"
	"os"
)

var (
	DevZero io.Reader
	DevNull io.Writer
)

func init() {

	var err error

	DevZero, err = os.OpenFile("/dev/zero", os.O_RDONLY, 0)
	if err != nil {
		panic(fmt.Sprintf("Failed to open /dev/zero: %s", err))
	}

	DevNull, err = os.OpenFile("/dev/null", os.O_WRONLY, 0)
	if err != nil {
		panic(fmt.Sprintf("Failed to open /dev/null: %s", err))
	}
}
