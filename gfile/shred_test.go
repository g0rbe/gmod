package gfile

import (
	"os"
	"testing"
)

func TestShred(t *testing.T) {

	testfile := "testfile.shred"

	err := os.WriteFile(testfile, make([]byte, 1024), 0600)
	if err != nil {
		t.Fatalf("Failed to write %s: %s\n", testfile, err)
	}

	written, err := Shred(testfile, 3, true, false)
	if err != nil {
		t.Fatalf("Failed to shred with random: %s\n", err)
	}

	t.Logf("Written bytes: %d\n", written)

	written, err = Shred(testfile, 1, false, true)
	if err != nil {
		t.Fatalf("Failed to shred with zero: %s\n", err)
	}

	t.Logf("Written bytes: %d\n", written)

	exist, err := IsExists(testfile)
	if err != nil {
		t.Fatalf("Failed to check if %s is exist: %s\n", testfile, err)
	}

	if exist {
		t.Fatalf("%s should not exist\n", testfile)
	}
}
