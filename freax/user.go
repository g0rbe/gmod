package freax

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PasswdFile = "/etc/passwd"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// Passwd stores the user account information.
//
// See passwd(5) for more info.
type User struct {
	Username string // Login name
	Password string // Optional encrypted password
	UID      int    // Numerical user ID
	GID      int    // Numerical group ID
	Gecos    string // User name or comment field
	Home     string // User  home directory
	Shell    string // Optional user command interpreter
}

// userParsePasswdLine parses a line of /etc/passwd and returns a struct from it.
func userParsePasswdLine(line string) (*User, error) {

	fields := strings.Split(line, ":")
	if len(fields) != 7 {
		return nil, fmt.Errorf("invalid line")
	}

	var usr = new(User)
	var err error

	usr.Username = fields[0]
	usr.Password = fields[1]

	if fields[2] != "" {
		if usr.UID, err = strconv.Atoi(fields[2]); err != nil {
			return nil, fmt.Errorf("failed to convert UID: %s", err)
		}
	}

	if fields[3] != "" {
		if usr.GID, err = strconv.Atoi(fields[3]); err != nil {
			return nil, fmt.Errorf("failed to convert Gid: %s", err)
		}
	}

	usr.Gecos = fields[4]
	usr.Home = fields[5]
	usr.Shell = fields[6]

	return usr, nil
}

// LookupUser search for user with username username.
// If the user cannot be found, returns ErrUserNotFound.
//
// This function parses the "/etc/passwd" file.
func LookupUser(username string) (*User, error) {

	if len(username) == 0 {
		return nil, fmt.Errorf("username is empty")
	}

	file, err := os.OpenFile(PasswdFile, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", PasswdFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		usr, err := userParsePasswdLine(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to read line: %w", err)
		}

		if usr.Username == username {
			return usr, nil
		}
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", PasswdFile, err)
	}

	return nil, ErrUserNotFound
}

// LookupUserID search for user with ID uid.
// If the user cannot be found, returns ErrUserNotFound.
//
// This function parses the "/etc/passwd" file.
func LookupUserID(uid int) (*User, error) {

	file, err := os.OpenFile(PasswdFile, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", PasswdFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		usr, err := userParsePasswdLine(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to read line: %w", err)
		}

		if usr.UID == uid {
			return usr, nil
		}
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", PasswdFile, err)
	}

	return nil, ErrUserNotFound
}
