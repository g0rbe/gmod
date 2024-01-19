package freax

import (
	"fmt"
	"os"
)

type FileMode uint32

// From linux/stat.h
// See more: https://www.gnu.org/software/libc/manual/html_node/Permission-Bits.html
// See more: https://www.gnu.org/software/libc/manual/html_node/Testing-File-Type.html

const (
	S_IFMT   FileMode = 00170000 // This is a bit mask used to extract the file type code from a mode value.
	S_IFSOCK FileMode = 0140000  // This is the file type constant of a socket.
	S_IFLNK  FileMode = 0120000  // This is the file type constant of a symbolic link.
	S_IFREG  FileMode = 0100000  // This is the file type constant of a regular file.
	S_IFBLK  FileMode = 0060000  // This is the file type constant of a block-oriented device file.
	S_IFDIR  FileMode = 0040000  // This is the file type constant of a directory file.
	S_IFCHR  FileMode = 0020000  // This is the file type constant of a character-oriented device file.
	S_IFIFO  FileMode = 0010000  // This is the file type constant of a FIFO or pipe.
	S_ISUID  FileMode = 0004000  // This is the set-user-ID on execute bit.
	S_ISGID  FileMode = 0002000  // This is the set-group-ID on execute bit.
	S_ISVTX  FileMode = 0001000  // This is the sticky bit.
	S_IRWXU  FileMode = 0000700  // This is equivalent to (S_IRUSR | S_IWUSR | S_IXUSR).
	S_IRUSR  FileMode = 00000400 // Read permission bit for the owner of the file.
	S_IWUSR  FileMode = 00000200 // Write permission bit for the owner of the file.
	S_IXUSR  FileMode = 00000100 // Execute (for ordinary files) or search (for directories) permission bit for the owner of the file.
	S_IRWXG  FileMode = 00000070 // This is equivalent to (S_IRGRP | S_IWGRP | S_IXGRP).
	S_IRGRP  FileMode = 00000040 // Read permission bit for the group owner of the file.
	S_IWGRP  FileMode = 00000020 // Write permission bit for the group owner of the file.
	S_IXGRP  FileMode = 00000010 // Execute or search permission bit for the group owner of the file.
	S_IRWXO  FileMode = 00000007 // This is equivalent to (S_IROTH | S_IWOTH | S_IXOTH).
	S_IROTH  FileMode = 00000004 // Read permission bit for other users.
	S_IWOTH  FileMode = 00000002 // Write permission bit for other users.
	S_IXOTH  FileMode = 00000001 // Execute or search permission bit for other users.

)

// FileModeSet sets m bit in mode.
func FileModeSet(mode, m FileMode) FileMode {
	return mode | m
}

// FileModeClear clears m bit in mode.
func FileModeClear(mode, m FileMode) FileMode {
	return mode ^ (mode & m)
}

// FileModeClear clears m bit in mode.
func FileModeToggle(mode, m FileMode) FileMode {
	return mode ^ m
}

// FileModeIsSet check whether m bit is set in mode.
func FileModeIsSet(mode, m FileMode) bool {
	return mode&m == m
}

// FileModeIsLnk returns whether the file is a symbolic link.
func FileModeIsLnk(m FileMode) bool {

	// S_ISLNK(m) (((m) & S_IFMT) == S_IFLNK)

	return m&S_IFMT == S_IFLNK
}

// FileModeIsReg returns whether the file is a regular file.
func FileModeIsReg(m FileMode) bool {

	// S_ISREG(m) (((m) & S_IFMT) == S_IFREG)

	return m&S_IFMT == S_IFREG
}

// FileModeIsDir returns whether the file is a directory.
func IsDir(m FileMode) bool {

	// S_ISDIR(m) (((m) & S_IFMT) == S_IFDIR)

	return m&S_IFMT == S_IFDIR
}

// FileModeIsChr returns whether the file is a character special file (a device like a terminal).
func FileModeIsChr(m FileMode) bool {

	// S_ISCHR(m) (((m) & S_IFMT) == S_IFCHR)

	return m&S_IFMT == S_IFCHR
}

// FileModeIsBlk returns whether the file is a block special file (a device like a disk).
func FileModeIsBlk(m FileMode) bool {

	// S_ISBLK(m) (((m) & S_IFMT) == S_IFBLK)

	return m&S_IFMT == S_IFBLK
}

// FileModeIsFifo returns whether the file is a FIFO special file, or a pipe.
func FileModeIsFifo(m FileMode) bool {

	// S_ISFIFO(m) (((m) & S_IFMT) == S_IFIFO)

	return m&S_IFMT == S_IFIFO
}

// FileModeIsSock returns  whether the file is a socket.
func FileModeIsSock(m FileMode) bool {

	// S_ISSOCK(m) (((m) & S_IFMT) == S_IFSOCK)

	return m&S_IFMT == S_IFSOCK
}

// FileModeGetPath returns the mode of the file in path.
// This function (and every *Path functions) does not follows symbolic link.
func FileModeGetPath(path string) (FileMode, error) {

	stat, err := Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("failed to stat: %w", err)
	}

	return FileMode(stat.Mode), nil
}

// FileModeSetPath sets the mode bit m for file in path.
func FileModeSetPath(path string, m FileMode) error {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

	stat, err := Fstat(file)
	if err != nil {
		return fmt.Errorf("dtat failed: %w", err)
	}

	return file.Chmod(os.FileMode(FileModeSet(stat.Mode, m)))
}

// FileModeUnsetPath unsets the mode bit m for file in path.
func FileModeUnsetPath(path string, m FileMode) error {

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

	stat, err := Fstat(file)
	if err != nil {
		return fmt.Errorf("dtat failed: %w", err)
	}

	return file.Chmod(os.FileMode(FileModeSet(stat.Mode, m)))
}

// FileModeIsSetPath returns whether mode m of file is set in path.
// This function does not follow symbolic link.
func FileModeIsSetPath(path string, m FileMode) (bool, error) {

	stat, err := Lstat(path)
	if err != nil {
		return false, fmt.Errorf("failed to stat: %w", err)
	}

	return FileModeIsSet(stat.Mode, m), nil
}

// IsLnkPath returns whether the file in path is a symbolic link.
func IsLnkPath(path string) (bool, error) {

	m, err := FileModeGetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFLNK, nil
}

// IsRegPath returns whether the file in path is a regular file.
func IsRegPath(path string) (bool, error) {

	m, err := FileModeGetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFREG, nil
}

// IsDirPath returns whether the file in path is a directory.
func IsDirPath(path string) (bool, error) {

	m, err := FileModeGetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFDIR, nil
}

// IsChrPath returns whether the file in path is a character special file (a device like a terminal).
func IsChrPath(path string) (bool, error) {

	m, err := FileModeGetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFCHR, nil
}

// IsBlkPath returns whether the file in path is a block special file (a device like a disk).
func IsBlkPath(path string) (bool, error) {

	m, err := FileModeGetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFBLK, nil
}

// IsFifoPath returns whether the file in path is a FIFO special file, or a pipe.
func IsFifoPath(path string) (bool, error) {

	m, err := FileModeGetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFIFO, nil
}

// IsSockPath returns whether the file in path is a socket.
func IsSockPath(path string) (bool, error) {

	m, err := FileModeGetPath(path)
	if err != nil {
		return false, fmt.Errorf("failed to get mode: %w", err)
	}

	return m&S_IFMT == S_IFSOCK, nil
}
