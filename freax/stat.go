package freax

import (
	"errors"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

var (
	ErrAssert = errors.New("syscall.Stat_t assertion failed")
)

// FileInfo describes information about a file.
// See struct stat at stat(3) manual for more info.
type FileInfo struct {
	Dev     uint64     // ID of device containing file
	Ino     uint64     // Inode number
	Mode    FileMode   // File type and mode
	Nlink   uint64     // Number of hard links
	User    user.User  // Owner user. If invalid uid found (that can happen), only the Uid field is set.
	Group   user.Group // Owner group. If invalid gid found (that can happen), only the Gid field is set.
	Rdev    uint64     // Device ID (if special file)
	Size    int64      // Total size, in bytes
	Blksize int64      // Block size for filesystem I/O
	Blocks  int64      // Number of 512 B blocks allocated
	Atim    time.Time  // Time of last access
	Mtim    time.Time  // Time of last modification
	Ctim    time.Time  // Time of last status change
}

func fileInfoFromStatT(s *unix.Stat_t) *FileInfo {

	if s == nil {
		return nil
	}

	// User
	// In case of error, set User to a user.User with only the Uid field set.
	uidStr := strconv.Itoa(int(s.Uid))
	usr, err := user.LookupId(uidStr)
	if err != nil {
		usr = &user.User{Uid: uidStr}
	}

	// Group
	// In case of error, set Group to a user.Group with only the Gid field set.
	gidStr := strconv.Itoa(int(s.Gid))
	grp, err := user.LookupGroupId(gidStr)
	if err != nil {
		grp = &user.Group{Gid: gidStr}
	}

	i := new(FileInfo)

	i.Dev = s.Dev
	i.Ino = s.Ino
	i.Mode = FileMode(s.Mode)
	i.Nlink = s.Nlink
	i.User = *usr
	i.Group = *grp
	i.Rdev = s.Rdev
	i.Size = s.Size
	i.Blksize = s.Blksize
	i.Blocks = s.Blocks
	i.Atim = time.Unix(s.Atim.Sec, s.Atim.Nsec)
	i.Mtim = time.Unix(s.Mtim.Sec, s.Mtim.Nsec)
	i.Ctim = time.Unix(s.Ctim.Sec, s.Ctim.Nsec)

	return i
}

// Stat returns information about a file specified by path.
//
// No permissions are required on the file itself, execute (search) permission is required on all of the directories in pathname that lead to the file.
// See stat(2) manual for more info.
func Stat(path string) (*FileInfo, error) {

	s := new(unix.Stat_t)
	err := unix.Stat(path, s)
	if err != nil {
		return nil, err
	}

	return fileInfoFromStatT(s), nil
}

// Fstat returns information about a file specified by file.
func Fstat(fd int) (*FileInfo, error) {

	s := new(unix.Stat_t)

	err := unix.Fstat(fd, s)
	if err != nil {
		return nil, err
	}

	return fileInfoFromStatT(s), nil
}

// Lstat returns information about a file specified by path.
// If path is a symbolic link, then it returns information about the link itself, not the file that the link refers to.
//
// No permissions are required on the file itself, execute (search) permission is required on all of the directories in pathname that lead to the file.
// See stat(2) manual for more info.
func Lstat(path string) (*FileInfo, error) {

	s := new(unix.Stat_t)
	err := unix.Lstat(path, s)
	if err != nil {
		return nil, err
	}

	return fileInfoFromStatT(s), nil
}

// GetFileOwner returns the owner user.User struct of file.
// Returns ErrAssert if syscall assertion failed or
// UnknownUserIdError if failed to get owner from the UID.
//
// Deprecated: Use Stat(), Fstat() or Lstat() instead.
func GetFileOwner(file string) (*user.User, error) {

	stat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	sysStat, ok := stat.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, ErrAssert
	}

	return user.LookupId(strconv.FormatUint(uint64(sysStat.Uid), 10))
}

// GetFileGroup returns the user.Group struct of file.
// Returns ErrAssert if syscall assertion failed or
// UnknownGroupIdError if failed to get owner from the GID.
//
// Deprecated: Use Stat(), Fstat() or Lstat() instead.
func GetFileGroup(file string) (*user.Group, error) {

	stat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	sysStat, ok := stat.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, ErrAssert
	}

	return user.LookupGroupId(strconv.FormatUint(uint64(sysStat.Gid), 10))
}
