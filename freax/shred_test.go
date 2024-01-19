package freax_test

import (
	"os"
	"testing"

	"github.com/g0rbe/gmod/freax"
)

func TestShred(t *testing.T) {

	testfile := "testfile.shred"

	err := os.WriteFile(testfile, make([]byte, 1024), 0600)
	if err != nil {
		t.Fatalf("Failed to write %s: %s\n", testfile, err)
	}

	written, err := freax.Shred(testfile, 3, true)
	if err != nil {
		t.Fatalf("Failed to shred with random: %s\n", err)
	}

	t.Logf("Written bytes: %d\n", written)

	written, err = freax.Shred(testfile, 1, false)
	if err != nil {
		t.Fatalf("Failed to shred with zero: %s\n", err)
	}

	t.Logf("Written bytes: %d\n", written)

	err = os.Remove(testfile)
	if err != nil {
		t.Fatalf("Failed to remove %s: %s\n", testfile, err)
	}
}
