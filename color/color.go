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

func Red(a any) string {
	return fmt.Sprint(RED, a, RESET)
}

func Green(a any) string {
	return fmt.Sprint(GREEN, a, RESET)
}

func Yellow(a any) string {
	return fmt.Sprint(YELLOW, a, RESET)
}

func Blue(a any) string {
	return fmt.Sprint(BLUE, a, RESET)
}

func Purple(a any) string {
	return fmt.Sprint(PURPLE, a, RESET)
}

func Cyan(a any) string {
	return fmt.Sprint(CYAN, a, RESET)
}

func Gray(a any) string {
	return fmt.Sprint(GRAY, a, RESET)
}

func White(a any) string {
	return fmt.Sprint(WHITE, a, RESET)
}
