// Send desktop notifications with notify-send
package dntfy

import (
	"fmt"
	"os"
	"os/exec"
)

// Send sends a desktop notification via the "notify-send" executable.
// urgency must be empty (""), "low", "normal" or "critical".
// summary and body are required.
//
// If add is empty, os.Args[0] is used.
// If urgency is empty, "normal" is used.
// icon is omitted if empty.
func Send(app string, urgency string, icon string, summary, body string) error {

	if urgency != "" && urgency != "low" && urgency != "normal" && urgency != "critical" {
		return fmt.Errorf("invalid urgency: %s", urgency)
	}

	if summary == "" {
		return fmt.Errorf("missing summary")
	}

	if body == "" {
		return fmt.Errorf("missing body")
	}

	args := make([]string, 0)

	if app == "" {
		app = os.Args[0]
	}
	args = append(args, "--app-name", app)

	if urgency == "" {
		urgency = "normal"
	}
	args = append(args, "--urgency", urgency)

	if icon != "" {
		args = append(args, "--icon", icon)
	}

	args = append(args, summary, body)

	out, err := exec.Command("notify-send", args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s (%s)", out, err)
	}

	if len(out) != 0 {
		return fmt.Errorf(string(out))
	}

	return nil
}
