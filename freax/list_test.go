package freax_test

import (
	"testing"

	"github.com/g0rbe/gmod/freax"
)

func TestListDir(t *testing.T) {

	list, err := freax.ListDir(".", true, true)
	if err != nil {
		t.Fatalf("Fail: %s\n", err)
	}

	t.Logf("%v\n", list)
}

func BenchmarkListDir(b *testing.B) {

	for i := 0; i < b.N; i++ {
		freax.ListDir(".", true, true)
	}
}

func TestListFiles(t *testing.T) {

	list, err := freax.ListFiles(".")
	if err != nil {
		t.Fatalf("Fail: %s\n", err)
	}

	t.Logf("%v\n", list)
}

func BenchmarkListFiles(b *testing.B) {

	for i := 0; i < b.N; i++ {
		freax.ListFiles(".")
	}
}
