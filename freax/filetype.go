package freax

type FileType uint32

const (
	_S_IFMT   FileType = 00170000 // This is a bit mask used to extract the file type code from a FileType value.
	_S_IFSOCK FileType = 0140000  // This is the file type constant of a socket.
	_S_IFLNK  FileType = 0120000  // This is the file type constant of a symbolic link.
	_S_IFREG  FileType = 0100000  // This is the file type constant of a regular file.
	_S_IFBLK  FileType = 0060000  // This is the file type constant of a block-oriented device file.
	_S_IFDIR  FileType = 0040000  // This is the file type constant of a directory file.
	_S_IFCHR  FileType = 0020000  // This is the file type constant of a character-oriented device file.
	_S_IFIFO  FileType = 0010000  // This is the file type constant of a FIFO or pipe.
)

func (t FileType) String() string {

	switch t := t & _S_IFMT; t {
	case _S_IFSOCK:
		return "s"
	case _S_IFLNK:
		return "l"
	case _S_IFREG:
		return "r"
	case _S_IFBLK:
		return "b"

	case _S_IFDIR:
		return "d"
	case _S_IFCHR:
		return "c"

	case _S_IFIFO:
		return "f"

	default:
		return "?"
	}
}
