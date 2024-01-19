package freax_test

import (
	"os"
	"testing"

	"github.com/g0rbe/gmod/freax"
)

var (
	TestFileInfoPath = "fileinfo.go"
)

func TestStat(t *testing.T) {

	s, err := freax.Stat("fileinfo.go")
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	t.Logf("%#v\n", s)
}

func BenchmarkStat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		freax.Stat("fileinfo.go")
	}
}

func TestFstat(t *testing.T) {

	file, err := os.OpenFile("fileinfo.go", os.O_RDONLY, 0)
	if err != nil {
		t.Fatalf("Failed to open fileinfo.go: %s\n", err)
	}

	t.Cleanup(func() {
		file.Close()
	})

	s, err := freax.Fstat(file)
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	t.Logf("%#v\n", s)
}

func BenchmarkFstat(b *testing.B) {

	file, err := os.OpenFile("fileinfo.go", os.O_RDONLY, 0)
	if err != nil {
		b.Fatalf("Failed to open fileinfo.go: %s\n", err)
	}

	b.Cleanup(func() {
		file.Close()
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.Fstat(file)
	}
}
func TestLstat(t *testing.T) {

	s, err := freax.Lstat("fileinfo.go")
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	t.Logf("%#v\n", s)
}

func BenchmarkLstat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		freax.Lstat("fileinfo.go")
	}
}
