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

	s, err := freax.Stat("stat.go")
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	t.Logf("%#v\n", s)
}

func BenchmarkStat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		freax.Stat("stat.go")
	}
}

func TestFstat(t *testing.T) {

	file, err := os.OpenFile("stat.go", os.O_RDONLY, 0)
	if err != nil {
		t.Fatalf("Failed to open stat.go: %s\n", err)
	}

	t.Cleanup(func() {
		file.Close()
	})

	s, err := freax.Fstat(int(file.Fd()))
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	t.Logf("%#v\n", s)
}

func BenchmarkFstat(b *testing.B) {

	file, err := os.OpenFile("stat.go", os.O_RDONLY, 0)
	if err != nil {
		b.Fatalf("Failed to open stat.go: %s\n", err)
	}

	_fd := int(file.Fd())

	b.Cleanup(func() {
		file.Close()
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.Fstat(_fd)
	}
}
func TestLstat(t *testing.T) {

	s, err := freax.Lstat("stat.go")
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	t.Logf("%#v\n", s)
}

func BenchmarkLstat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		freax.Lstat("stat.go")
	}
}
