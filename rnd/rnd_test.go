package rnd

import (
	crand "crypto/rand"
	mrand "math/rand"
	"testing"

	"github.com/g0rbe/gmod/bitter"
)

func TestBytes(t *testing.T) {

	buf := Bytes(8)

	t.Logf("%v\n", buf)
}

func BenchmarkBytes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Bytes(8)
	}
}

func BenchmarkCryptoBytes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		bitter.ReadReaderBytes(crand.Reader, 8)
	}
}

func TestInt64(t *testing.T) {

	v := Int64()

	t.Logf("%d\n", v)
}

func BenchmarkInt64(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Int64()
	}
}

func TestInt64n(t *testing.T) {

	var (
		total = 10000000
	)

	var (
		zero  int
		one   int
		two   int
		three int
		four  int
	)

	for i := 0; i < total; i++ {

		switch v := Int64n(5); v {
		case 0:
			zero++
		case 1:
			one++
		case 2:
			two++
		case 3:
			three++
		case 4:
			four++
		default:
			t.Fatalf("Not in range [0;5): %d", v)
		}
	}

	if zero+one+two+three+four != total {
		t.Fatalf("Invalid total: want %d, got %d\n", total, zero+one+two+three+four)
	}

	t.Logf("0: %d, 1: %d, 2: %d, 3: %d, 4: %d, total: %d\n", zero, one, two, three, four, zero+one+two+three+four)
}

func BenchmarkInt64n(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Int64n(5)
	}
}

func BenchmarkMathInt64n(b *testing.B) {

	for i := 0; i < b.N; i++ {
		mrand.Int63n(5)
	}
}
func TestRandomString(t *testing.T) {

	v := String([]byte("abcdefgh"), 10)
	if len(v) != 10 {
		t.Fatalf("FAIL: invalid length: %d, want: 10", len(v))
	}
}

func BenchmarkRandomString(b *testing.B) {

	for i := 0; i < b.N; i++ {
		String([]byte("abcdefgh0"), 256)
	}
}
