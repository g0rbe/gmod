// Package colorz colorize the terminal output.
package colorz

import "fmt"

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	GRAY   = "\033[37m"
	WHITE  = "\033[97m"
)

// Colorize returns a as a string with color prefix and RESET suffix (to clear color).
func Colorize(color string, a any) string {
	return fmt.Sprint(color, a, RESET)
}

// Red returns a as a string with RED prefix and RESET suffix (to clear color).
// This function is Colorize(RED, a).
func Red(a any) string {
	return Colorize(RED, a)
}

func Green(a any) string {
	return Colorize(GREEN, a)
}

func Yellow(a any) string {
	return Colorize(YELLOW, a)
}

func Blue(a any) string {
	return Colorize(BLUE, a)
}

func Purple(a any) string {
	return Colorize(PURPLE, a)
}

func Cyan(a any) string {
	return Colorize(CYAN, a)
}

func Gray(a any) string {
	return Colorize(GRAY, a)
}

func White(a any) string {
	return Colorize(WHITE, a)
}
