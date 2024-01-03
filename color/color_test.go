package color

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Printf("%s\n", Red("red"))
	fmt.Printf("%s\n", Green("green"))
	fmt.Printf("%s\n", Yellow("yellow"))
	fmt.Printf("%s\n", Blue("blue"))
	fmt.Printf("%s\n", Purple("purple"))
	fmt.Printf("%s\n", Cyan("cyan"))
	fmt.Printf("%s\n", Gray("gray"))
	fmt.Printf("%s\n", White("white"))
}
