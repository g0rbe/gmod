package freax

import (
	"sync"

	"golang.org/x/sys/unix"
)

// Flags to OpenFile wrapping those of the underlying system. Not all flags may be implemented on a given system.
const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY = unix.O_RDONLY // open the file read-only.
	O_WRONLY = unix.O_WRONLY // open the file write-only.
	O_RDWR   = unix.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND = unix.O_APPEND // append data to the file when writing.
	O_CREATE = unix.O_CREAT  // create a new file if none exists.
	O_EXCL   = unix.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   = unix.O_SYNC   // open for synchronous I/O.
	O_TRUNC  = unix.O_TRUNC  // truncate regular writable file when opened.
)

// Seek whence values.
const (
	SEEK_SET = unix.SEEK_SET // seek relative to the origin of the file
	SEEK_CUR = unix.SEEK_CUR // seek relative to the current offset
	SEEK_END = unix.SEEK_END // seek relative to the end
)

type File struct {
	fd int
	m  *sync.RWMutex
}

func OpenFile(pathname string, flags int, mode uint32) (*File, error) {

	fd, err := unix.Open(pathname, flags, mode)
	if err != nil {
		return nil, err
	}

	return &File{fd: fd, m: new(sync.RWMutex)}, nil
}

func (f *File) Read(p []byte) (n int, err error) {
	return FDRead(f.Fd(), p)
}

func (f *File) ReadByte() (byte, error) {

	b := make([]byte, 1)

	_, err := f.Read(b)

	return b[0], err
}

// func (f *File) ReadFrom(r io.Reader) (n int64, err error) {

// }

func (f *File) Write(p []byte) (n int, err error) {
	return unix.Write(f.fd, p)
}

func (f *File) WriteByte(c byte) error {

	_, err := unix.Write(f.fd, []byte{c})
	return err
}

func (f *File) WriteString(s string) (n int, err error) {
	return f.Write([]byte(s))
}

func (f *File) Seek(offset int64, whence int) (int64, error) {
	return unix.Seek(f.fd, offset, whence)
}

func (f *File) Close() error {

	if f == nil {
		return nil
	}

	return unix.Close(f.fd)
}

func (f *File) Fd() uintptr {
	return uintptr(f.fd)
}

func (f *File) Stat() (*FileInfo, error) {

	return Fstat(f.fd)
}

func (f *File) Sync() error {
	return unix.Fsync(f.fd)
}

// Lock locks the fd with FcntlFlock().
// Blocks until the lock is not released.
func (f *File) Lock() error {

	fl, err := FcntlGetfl(uintptr(f.fd))
	if err != nil {
		return err
	}

	if fl&O_RDONLY == O_RDONLY {
		return FlockSetReadWait(uintptr(f.fd))
	}

	return FlockSetWriteWait(uintptr(f.fd))
}

// Unlock unlocks the lockon fd.
// Blocks until the lock is not released.
func (f *File) Unlock() error {

	return FlockUnlockWait(uintptr(f.fd))
}
