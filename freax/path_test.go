package freax_test

import (
	"os"
	"testing"

	"github.com/g0rbe/gmod/freax"
)

func TestReadAll(t *testing.T) {

	b, err := freax.ReadAll("path.go")
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	t.Logf("Read bytes: %d\n", len(b))
}

func BenchmarkReadAll(b *testing.B) {

	for i := 0; i < b.N; i++ {
		freax.ReadAll("path.go")
	}
}

func TestWrite(t *testing.T) {

	t.Cleanup(func() {
		os.RemoveAll("test.txt")
	})

	err := freax.Write("test.txt", []byte{'a', 'b', 'c'}, 0666)
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}
}

func BenchmarkWrite(b *testing.B) {

	b.Cleanup(func() {
		os.RemoveAll("test.txt")
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.Write("test.txt", []byte{'a', 'b', 'c'}, 0666)
	}
}

func TestWriteSync(t *testing.T) {

	t.Cleanup(func() {
		os.RemoveAll("test.txt")
	})

	err := freax.WriteSync("test.txt", []byte{'a', 'b', 'c'}, 0666)
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}
}

func BenchmarkWriteSync(b *testing.B) {

	b.Cleanup(func() {
		os.RemoveAll("test.txt")
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.WriteSync("test.txt", []byte{'a', 'b', 'c'}, 0666)
	}
}
