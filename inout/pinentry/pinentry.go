// Provides basic API for pinentry
//
// BUG(g0rbe): If the password does not match, pinentry will freeze and the only option is to Cancel when ErrOpCancelled is returned.
package pinentry

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"

	"github.com/g0rbe/gmod/inout"
)

var (
	DefaultPath string = "/usr/bin/pinentry"
	LineOk             = []byte{'O', 'K'}
	PrefixOk           = []byte{'O', 'K', ' '}
	PrefixErr          = []byte{'E', 'R', 'R', ' '}
	PrefixData         = []byte{'D', ' '}
)

// writeCommand writes command to w with a newline character ('\n').
func writeCommand(w io.Writer, c *Command) error {

	// Write command with a newline character
	_, err := fmt.Fprintf(w, "%s\n", c)

	return err
}

// readLines returns the interesting lines.
// OK line (line prefixed with "OK ") is skipped.
// ERR line (prefixed with "ERR ") is returned returns Error error.
func readLines(r io.Reader) ([][]byte, error) {

	var lines [][]byte

	for {

		line, err := inout.ReadLine(r)
		if err != nil {
			return lines, err
		}

		// Return without appending OK line
		if bytes.Equal(line, LineOk) {
			return lines, nil
		}

		// Return ERR line as Error
		if bytes.HasPrefix(line, PrefixErr) {
			return lines, ParseError(line)
		}

		// Skip greeting line
		if bytes.Equal(line, []byte("OK Pleased to meet you")) {
			continue
		}

		// // Skip OK lines with message
		// if bytes.HasPrefix(line, PrefixOk) {
		// 	continue
		// }

		lines = append(lines, line)
	}

}

func Run(commands ...*Command) ([]byte, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, "pinentry")

	input, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	defer input.Close()

	output, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer output.Close()

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	defer cmd.Cancel()

	var (
		result []byte
	)

	// Write commands
	for i := range commands {

		// SKip unsupported commands
		if commands[i] == nil {
			continue
		}

		err = writeCommand(input, commands[i])
		if err != nil {
			return nil, fmt.Errorf("failed to write %s: %w", commands[i], err)
		}

		lines, err := readLines(output)
		if err != nil {

			// EOF if not an error if command is BYE.
			if commands[i].Command == "BYE" && errors.Is(err, io.EOF) {
				break
			}

			return nil, fmt.Errorf("failed to read %s: %w", commands[i], err)
		}

		for j := range lines {

			switch commands[i].Command {
			case "GETPIN":
				if bytes.HasPrefix([]byte(lines[j]), PrefixData) {
					result = bytes.TrimPrefix(lines[j], PrefixData)
				}
			}
		}
	}

	return result, cmd.Wait()
}

// Message displays a message with pinentry.
//
// The Timeout is set to DefaultTimeout.
// The Title is set to DefaultTitle.
// The Ok button is set to DefaultOk.
// The Cancel button is set to DefaultCancel.
func Message(prompt, description, message string) error {

	_, err := Run(
		CommandSetTimeout(DefaultTimeout),
		CommandSetTitle(DefaultTitle),
		CommandSetPrompt(prompt),
		CommandSetDescription(description),
		CommandSetOk(DefaultOk),
		CommandSetCancel(DefaultCancel),
		CommandMessage(message),
		CommandBye(),
	)

	return err
}

// Confirm displays a confirm with pinentry.
// Returns true if Ok button used.
// Returns false and nil is Cancel button used.
//
// The Timeout is set to DefaultTimeout.
// The Title is set to DefaultTitle.
func Confirm(prompt, description, ok, cancel string) (bool, error) {

	_, err := Run(
		CommandSetTimeout(DefaultTimeout),
		CommandSetTitle(DefaultTitle),
		CommandSetPrompt(prompt),
		CommandSetDescription(description),
		CommandSetOk(ok),
		CommandSetCancel(cancel),
		CommandConfirm(),
		CommandBye(),
	)

	if err != nil {
		return true, nil
	}

	if errors.Is(err, ErrOpCancelled) {
		return false, nil
	}

	return false, err
}

// Password reads a password with pinentry.
//
// The Timeout is set to DefaultTimeout.
// The Title is set to DefaultTitle.
// The Ok button is set to DefaultOk.
// The Cancel button is set to DefaultCancel.
func Password(prompt, description string) ([]byte, error) {

	pass, err := Run(
		CommandSetTimeout(DefaultTimeout),
		CommandSetTitle(DefaultTitle),
		CommandSetPrompt(prompt),
		CommandSetDescription(description),
		CommandSetOk(DefaultOk),
		CommandSetCancel(DefaultCancel),
		CommandGetPin(),
		CommandBye(),
	)

	return pass, err
}

// PasswordRepeat reads a password with pinentry with SETREPEAT.
//
// The Timeout is set to DefaultTimeout.
// The Title is set to DefaultTitle.
// The Ok button is set to DefaultOk.
// The Cancel button is set to DefaultCancel.
func PasswordRepeat(prompt, description string) ([]byte, error) {

	pass, err := Run(
		CommandSetTimeout(DefaultTimeout),
		CommandSetTitle(DefaultTitle),
		CommandSetPrompt(prompt),
		CommandSetDescription(description),
		CommandSetOk(DefaultOk),
		CommandSetCancel(DefaultCancel),
		CommandSetRepeat(),
		CommandGetPin(),
		CommandBye(),
	)

	return pass, err
}
