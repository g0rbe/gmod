package freax_test

import (
	"os"
	"testing"

	"github.com/g0rbe/gmod/freax"
)

func TestLookupUser(t *testing.T) {

	c, err := freax.LookupUser(os.Getenv("USER"))
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	t.Logf("%v\n", c)
}

func BenchmarkLookupUser(b *testing.B) {

	usr := os.Getenv("USER")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.LookupUser(usr)
	}
}

func TestLookupUserID(t *testing.T) {

	c, err := freax.LookupUserID(os.Geteuid())
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}

	t.Logf("%v\n", c)
}

func BenchmarkLookupUserID(b *testing.B) {

	euid := os.Geteuid()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		freax.LookupUserID(euid)
	}
}
