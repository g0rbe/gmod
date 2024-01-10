package dns

import (
	"testing"
	"time"
)

func TestIsWildcard(t *testing.T) {

	cases := []struct {
		D string
		R bool
	}{
		{D: "example.com", R: false},
		{D: "www.example.com", R: false},
		{D: "test.cyberdivine.ch", R: true},
		{D: "test.classicbikes.ch", R: true},
	}

	for i := range cases {

		r, err := IsWildcard(cases[i].D, TypeA)
		if err != nil {
			t.Fatalf("FAIL: failed to check if %s is a wildcard: %s\n", cases[i].D, err)
		}

		if r != cases[i].R {
			t.Fatalf("FAIL: Result for %s: %v, want: %v \n", cases[i].D, r, cases[i].R)
		}
	}
}

func BenchmarkIsWildcard(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.IsWildcard("test.classicbikes.ch", TypeA)
	}
}

func BenchmarkIsWildcardInvalid(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.IsWildcard("www.example.com", TypeA)
	}
}
