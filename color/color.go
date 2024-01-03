package color

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

func Red(s string) string {
	return fmt.Sprint(RED, s, RESET)
}

func Green(s string) string {
	return fmt.Sprint(GREEN, s, RESET)
}

func Yellow(s string) string {
	return fmt.Sprint(YELLOW, s, RESET)
}

func Blue(s string) string {
	return fmt.Sprint(BLUE, s, RESET)
}

func Purple(s string) string {
	return fmt.Sprint(PURPLE, s, RESET)
}

func Cyan(s string) string {
	return fmt.Sprint(CYAN, s, RESET)
}

func Gray(s string) string {
	return fmt.Sprint(GRAY, s, RESET)
}

func White(s string) string {
	return fmt.Sprint(WHITE, s, RESET)
}
