package slitu

import (
	"errors"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

var (
	ErrAssert = errors.New("syscall.Stat_t assertion failed")
)

// GetFileOwner returns the owner user.User struct of file.
// Returns ErrAssert if syscall assertion failed or
// UnknownUserIdError if failed to get owner from the UID.
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
