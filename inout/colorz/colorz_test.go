package colorz

import (
	"strings"
	"testing"
)

func TestColorize(t *testing.T) {

	testText := "text"

	v := Colorize(RED, testText)
	if strings.Compare(v, RED+testText+RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestRed(t *testing.T) {

	testText := "text"

	v := Red(testText)
	if strings.Compare(v, RED+testText+RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestBlue(t *testing.T) {

	testText := "text"

	v := Blue(testText)
	if strings.Compare(v, BLUE+testText+RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestCyan(t *testing.T) {

	testText := "text"

	v := Cyan(testText)
	if strings.Compare(v, CYAN+testText+RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestGray(t *testing.T) {

	testText := "text"

	v := Gray(testText)
	if strings.Compare(v, GRAY+testText+RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestGreen(t *testing.T) {

	testText := "text"

	v := Green(testText)
	if strings.Compare(v, GREEN+testText+RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestPurple(t *testing.T) {

	testText := "text"

	v := Purple(testText)
	if strings.Compare(v, PURPLE+testText+RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestWhite(t *testing.T) {

	testText := "text"

	v := White(testText)
	if strings.Compare(v, WHITE+testText+RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestYellow(t *testing.T) {

	testText := "text"

	v := Yellow(testText)
	if strings.Compare(v, YELLOW+testText+RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}
