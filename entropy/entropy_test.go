package entropy_test

import (
	"testing"

	"github.com/g0rbe/gmod/entropy"
)

func TestGetandom(t *testing.T) {

	buf := make([]byte, 8)

	err := entropy.Getrandom(buf, 0)
	if err != nil {
		t.Fatalf("Error: %s'\n", err)
	}

	t.Logf("%v\n", buf)
}

func BenchmarkGetrandom(b *testing.B) {

	buf := make([]byte, 8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		entropy.Getrandom(buf, 0)
	}
}

func TestBytes(t *testing.T) {

	buf := entropy.Bytes(64)

	t.Logf("%v\n", buf)
}

func BenchmarkBytes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		entropy.Bytes(8)
	}
}

func TestInt(t *testing.T) {

	v := entropy.Int()

	t.Logf("%d\n", v)
}

func BenchmarkInt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		entropy.Int()
	}
}

func TestIntn(t *testing.T) {

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

		switch v := entropy.Intn(5); v {
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
		entropy.Intn(5)
	}
}

func TestRandomString(t *testing.T) {

	v := entropy.String([]byte("abcdefgh"), 10)
	if len(v) != 10 {
		t.Fatalf("FAIL: invalid length: %d, want: 10", len(v))
	}
}

func BenchmarkRandomString(b *testing.B) {

	for i := 0; i < b.N; i++ {
		entropy.String([]byte("abcdefgh0"), 256)
	}
}
