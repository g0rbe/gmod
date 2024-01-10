package dns

import (
	"errors"
	"testing"
	"time"
)

func TestQueryAll(t *testing.T) {

	TestDomain := "elmasy.com"

	rr, err := QueryAll(TestDomain)
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range rr {
		t.Logf("%s %s -> %s\n", TestDomain, TypeToString(rr[i].Type), rr[i].Value)
	}
}

func TestQueryAllInvalid(t *testing.T) {

	TestDomain := "invalid.example.com"

	rr, err := QueryAll(TestDomain)
	if err != nil && !errors.Is(err, ErrName) {
		t.Fatalf("FAIL: %s\n", err)
	}

	if len(rr) > 0 {
		t.Fatalf("FAIL: Invalid number os response: want: 0, got: %d\n", len(rr))
	}
}

func BenchmarkQueryAll(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.QueryAll("example.com")
	}
}

func BenchmarkQueryAllInvalid(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.QueryAll("invalid.example.com")
	}
}
