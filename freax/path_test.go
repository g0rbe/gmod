package freax_test

import (
	"fmt"
	"path"
	"testing"

	"github.com/g0rbe/gmod/freax"
)

func ExamplePathJoin() {

	fmt.Printf("%s\n", freax.PathJoin("test", "path"))
	fmt.Printf("%s\n", freax.PathJoin("test", "", "path"))
	fmt.Printf("%s\n", freax.PathJoin("test/", "", "path"))

	// Output:
	// test/path
	// test/path
	// test/path
}

func TestPathJoin(t *testing.T) {

	p := freax.PathJoin("test", "", "path/", "join")
	if p != "test/path/join" {
		t.Fatalf("Want \"test/path/join\", got \"%s\"\n", p)
	}
}

func BenchmarkPathJoin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		freax.PathJoin("test", "", "path/", "join")
	}
}

func BenchmarkStdPathJoin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		path.Join("test", "", "path/", "join")
	}
}
