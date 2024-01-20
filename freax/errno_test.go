package freax_test

import (
	"math/rand"
	"syscall"
	"testing"

	"github.com/g0rbe/gmod/freax"
)

// func TestNewErrno(t *testing.T) {

// 	err := freax.NewErrno(1000)
// 	if errors.Is(err, freax.Errno) {
// 		t.Fatalf("Error type: %T\n", err)
// 	}
// 	t.Logf("%s\n")
// }

func BenchmarkNewErrno(b *testing.B) {

	for i := 0; i < b.N; i++ {
		freax.NewErrno(rand.Intn(134))
	}
}

func TestErrnoFromSyscallErrno(t *testing.T) {

	_, err := syscall.Open("", -1, 0)
	if err == nil {
		t.Fatalf("Should return error\n")
	}

	err = freax.ErrnoFromSyscallErrno(err)

	t.Logf("%T\n", err)

}
