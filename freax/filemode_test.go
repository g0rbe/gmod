package freax_test

import (
	"io/fs"
	"os"
	"testing"

	"github.com/g0rbe/gmod/freax"
)

const (
	TEST_FILE = "testfile.txt"
	TEST_DIR  = "testdir"
)

func TestGetMode(t *testing.T) {

	m, err := freax.FileModeGetPath(".")
	if err != nil {
		t.Fatalf("Failed to get mode of \".\": %s\n", err)
	}

	t.Logf("Mode of \".\": %b\n", m)
	t.Logf("S_IFMT     : %b\n", freax.S_IFDIR|freax.S_IROTH|freax.S_IXOTH)

}

func TestFileMode(t *testing.T) {

	mode := freax.FileModeSet(0, freax.S_IFDIR)

	if !freax.IsDir(mode) {
		t.Fatalf("Must be dir\n")
	}
	t.Logf("+S_IFDIR: %b\n", mode)

	if mode = freax.FileModeSet(mode, freax.S_IRUSR); !freax.FileModeIsSet(mode, freax.S_IRUSR) {
		t.Fatalf("Mode must have S_IRUSR\n")
	}
	t.Logf("+S_IRUSR: %b\n", mode)

	if mode = freax.FileModeToggle(mode, freax.S_IRUSR); freax.FileModeIsSet(mode, freax.S_IRUSR) {
		t.Fatalf("Mode must not have S_IRUSR\n")
	}
	t.Logf("~S_IRUSR: %b\n", mode)

	if mode = freax.FileModeToggle(mode, freax.S_IRUSR); !freax.FileModeIsSet(mode, freax.S_IRUSR) {
		t.Fatalf("Mode must have S_IRUSR\n")
	}
	t.Logf("~S_IRUSR: %b\n", mode)

	if mode = freax.FileModeClear(mode, freax.S_IFDIR); freax.IsDir(mode) {
		t.Fatalf("Should not be dir\n")
	}
	t.Logf("-S_IFDIR: %b\n", mode)

	if mode = freax.FileModeClear(mode, freax.S_IFDIR); freax.IsDir(mode) {
		t.Fatalf("Should not be dir\n")
	}
	t.Logf("-S_IFDIR: %b\n", mode)

}

func TestPath(t *testing.T) {

	defer os.Remove(TEST_FILE)

	file, err := os.OpenFile(TEST_FILE, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("Failed to open %s: %s\n", TEST_FILE, err)
	}
	file.Close()

	isDir, err := freax.IsDirPath(TEST_FILE)
	if err != nil {
		t.Fatalf("Failed to check is dir: %s\n", err)
	}
	if isDir {
		t.Fatalf("Must not be a dir")
	}

	if err = freax.FileModeSetPath(TEST_FILE, freax.S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	}

	if isSet, err := freax.FileModeIsSetPath(TEST_FILE, freax.S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	} else if !isSet {
		t.Fatalf("S_IROTH must be set\n")
	}

	if err = freax.FileModeSetPath(TEST_FILE, freax.S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	}

	if isSet, err := freax.FileModeIsSetPath(TEST_FILE, freax.S_IROTH); err != nil {
		t.Fatalf("Failed to set S_IROTH: %s\n", err)
	} else if !isSet {
		t.Fatalf("S_IROTH must be set\n")
	}

}

func BenchmarkGetPath(b *testing.B) {

	defer os.Remove(TEST_FILE)

	file, err := os.OpenFile(TEST_FILE, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		b.Fatalf("Failed to open %s: %s\n", TEST_FILE, err)
	}
	file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.FileModeGetPath(TEST_FILE)
	}
}

func BenchmarkSetPath(b *testing.B) {

	defer os.Remove(TEST_FILE)

	file, err := os.OpenFile(TEST_FILE, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		b.Fatalf("Failed to open %s: %s\n", TEST_FILE, err)
	}
	file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.FileModeSetPath(TEST_FILE, freax.S_IXOTH)
	}
}

func TestIsDir(t *testing.T) {

	t.Logf("Perms: %b\n", 0070)

	t.Logf("fs.ModeDir: %b %o %x\n", fs.ModeDir, fs.ModeDir, fs.ModeDir)
	t.Logf("S_IFDIR: %b\n", freax.S_IFDIR)

	testName := "testdir"

	err := os.MkdirAll(testName, 0o600)
	if err != nil {
		t.Fatalf("Failed to create %s: %s\n", testName, err)
	}

	defer func() {
		err = os.Remove(testName)
		if err != nil {
			t.Fatalf("Failed to remove %s: %s\n", testName, err)
		}
	}()

	ok, err := freax.IsDirPath(testName)
	if err != nil {
		t.Fatalf("FAIL: failed to check directory: %s\n", err)
	}

	if !ok {
		m, err := freax.FileModeGetPath(testName)
		if err != nil {
			t.Fatalf("FAIL: failed to get mode: %s\n", err)
		}
		t.Fatalf("FAIL: %s should be a directory: %b %v\n", testName, m, freax.IsDir(m))
	}

}
