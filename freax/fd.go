package freax

import (
	"io"

	"golang.org/x/sys/unix"
)

// FcntlGetfl returns the file access mode and the file status flags.
func FcntlGetfl(fd uintptr) (int, error) {
	return unix.FcntlInt(fd, unix.F_GETFL, 0)
}

// FcntlCheckfl returns whether the flag is set for fd.
func FcntlCheckfl(fd uintptr, flag int) (bool, error) {

	f, err := unix.FcntlInt(fd, unix.F_GETFL, 0)
	return f&flag == flag, err
}

func FDRead(fd uintptr, p []byte) (n int, err error) {

	n, err = unix.Read(int(fd), p)

	// Return io.EOF if no error returned and no byte read
	if err == nil && n == 0 {
		err = io.EOF
	}

	return
}

func FDReadAll(fd uintptr) ([]byte, error) {

	b := make([]byte, 0)

	for {

		buf := make([]byte, 512)

		n, err := FDRead(fd, buf)
		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return b, err
		}

		b = append(b, buf[:n]...)
	}
}
