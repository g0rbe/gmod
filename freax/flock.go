package freax

import (
	"syscall"

	"golang.org/x/sys/unix"
)

var (
	F_RDLCK int16 = unix.F_RDLCK // Place a read lock on a file with FcntlFlock
	F_WRLCK int16 = unix.F_WRLCK // Place a write lock on a file with FcntlFlock
	F_UNLCK int16 = unix.F_UNLCK // Release the lock on a file with FcntlFlock
)

var (
	F_SETLK  = unix.F_SETLK  // Acquire or release a lock
	F_SETLKW = unix.F_SETLKW // As for F_SETLK, but if a conflicting lock is held on the file, then wait for that lock to be released
	F_GETLK  = unix.F_GETLK  // Return details about the current lock
)

// FlockGetRead returns whether read lock (F_RDLCK) could be placed on fd.
// Every byte will be checked for lock.
func FlockGetRead(fd uintptr) (bool, error) {

	flock := new(syscall.Flock_t)

	flock.Type = F_RDLCK

	err := syscall.FcntlFlock(fd, syscall.F_GETLK, flock)

	return flock.Type == syscall.F_UNLCK, err
}

// FlockGetRead returns whether write lock (F_WRLCK) could be placed on fd.
// Every byte will be checked for lock.
func FlockGetWrite(fd uintptr) (bool, error) {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_WRLCK

	err := syscall.FcntlFlock(fd, syscall.F_GETLK, flock)

	return flock.Type == syscall.F_UNLCK, err
}

// FlockSetRead acquire the advisory record lock for read on fd.
// This function returns error if a conflicting lock is held by another process.
// In order to place a read lock, fd must be open for reading.
func FlockSetRead(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_RDLCK

	return syscall.FcntlFlock(fd, syscall.F_SETLK, flock)
}

// FlockSetWrite acquire the advisory record lock for write on fd.
// This function returns error if a conflicting lock is held by another process.
// In order to place a write lock, fd must be open for writing.
func FlockSetWrite(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_WRLCK

	return syscall.FcntlFlock(fd, syscall.F_SETLK, flock)
}

// FlockUnlock release the advisory record lock on fd.
// This function returns error if a conflicting lock is held by another process.
func FlockUnlock(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_UNLCK

	return syscall.FcntlFlock(fd, syscall.F_SETLK, flock)
}

// FlockSetReadWait acquire the advisory record lock for read on fd.
// If a conflicting lock is held on the file, then wait for that lock to be released.
// In order to place a read lock, fd must be open for reading.
func FlockSetReadWait(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_RDLCK

	return syscall.FcntlFlock(fd, syscall.F_SETLKW, flock)
}

// FlockSetWriteWait acquire the advisory record lock for write on fd.
// If a conflicting lock is held on the file, then wait for that lock to be released.
// In order to place a write lock, fd must be open for writing.
func FlockSetWriteWait(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_WRLCK

	return syscall.FcntlFlock(fd, syscall.F_SETLKW, flock)
}

// FlockUnlockWait release the advisory record lock on fd.
// If a conflicting lock is held on the file, then wait for that lock to be released.
func FlockUnlockWait(fd uintptr) error {

	flock := new(syscall.Flock_t)

	flock.Type = syscall.F_UNLCK

	return syscall.FcntlFlock(fd, syscall.F_SETLKW, flock)
}
