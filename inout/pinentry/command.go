package pinentry

import (
	"fmt"
	"os"
)

var (
	DefaultTimeout = 30
	DefaultTitle   = os.Args[0]
	DefaultOk      = "Continue"
	DefaultCancel  = "Cancel"
)

type Command struct {
	Command string
	Value   string
}

func NewCommand(command string, v any) *Command {
	return &Command{Command: command, Value: fmt.Sprint(v)}
}

// CommandSetTimeout returns a command with SETTIMEOUT set to v.
func CommandSetTimeout(v int) *Command {
	return NewCommand("SETTIMEOUT", v)
}

// CommandSetPrompt returns a command with SETPROMPT set to v.
func CommandSetPrompt(v string) *Command {
	return NewCommand("SETPROMPT", v)
}

// CommandSetDescription returns a command with SETDESC set to v.
func CommandSetDescription(v string) *Command {
	return NewCommand("SETDESC", v)
}

// CommandSetTitle returns a command with SETTITLE set to v.
func CommandSetTitle(v string) *Command {
	return NewCommand("SETTITLE", v)
}

// CommandSetOk returns a command with SETOK set to v.
func CommandSetOk(v string) *Command {
	return NewCommand("SETOK", v)
}

// CommandSetCancel returns a command with SETCANCEL set to v.
func CommandSetCancel(v string) *Command {
	return NewCommand("SETCANCEL", v)
}

// CommandSetNotOk returns a command with SETNOTOK set to v.
func CommandSetNotOk(v string) *Command {
	return NewCommand("SETNOTOK", v)
}

// CommandSetError returns a command with SETERROR set to v.
func CommandSetError(v string) *Command {
	return NewCommand("SETERROR", v)
}

// CommandGetPin returns a command with GETPIN set to v.
func CommandGetPin() *Command {
	return NewCommand("GETPIN", "")
}

// CommandConfirm returns a command with CONFIRM set to v.
func CommandConfirm() *Command {
	return NewCommand("CONFIRM", "")
}

// CommandMessage returns a command with MESSAGE set to v.
func CommandMessage(v string) *Command {
	return NewCommand("MESSAGE", v)
}

// CommandBye returns a command with BYE set to v.
func CommandBye() *Command {
	return NewCommand("BYE", "")
}

// CommandSetRepeat returns a command with SETREPEAT set to v.
func CommandSetRepeat() *Command {
	return NewCommand("SETREPEAT", "")
}

// String returns the command and the value (if set).
// Eg.: "BYE" or "SETTIMEOUT 60"
func (c *Command) String() string {

	if c.Value == "" {
		return c.Command
	}

	return fmt.Sprintf("%s %s", c.Command, c.Value)
}
