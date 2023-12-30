package slitu

import "testing"

func TestRandomString(t *testing.T) {

	v := RandomString([]byte("abcdefgh"), 10)
	if len(v) != 10 {
		t.Fatalf("FAIL: invalid length: %d, want: 10", len(v))
	}
}

func BenchmarkRandomString(b *testing.B) {

	for i := 0; i < b.N; i++ {
		RandomString([]byte("abcdefgh0"), 256)
	}
}
