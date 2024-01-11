package pinentry

import (
	"errors"
	"testing"
)

func TestParse(t *testing.T) {

	errLine := []byte("ERR 83886179 Operation cancelled <Pinentry>")

	err := ParseError(errLine)
	if err == nil {
		t.Fatalf("Error is nil\n")
	}

	if !errors.Is(err, ErrOpCancelled) {
		t.Fatalf("Error should be %s, got %s\n", ErrOpCancelled, err)
	}

	t.Logf("%s\n", err)
}
