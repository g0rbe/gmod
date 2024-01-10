package dns

import (
	"testing"
	"time"
)

func TestQueryAAAA(t *testing.T) {

	r, err := QueryAAAA("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: No AAAA record for elmasy.com\n")
	}

	for i := range r {
		t.Logf("elmasy.com AAAA -> %s\n", r[i])
	}
}

func TestTryQueryAAAA(t *testing.T) {

	r, err := TryQueryAAAA("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: No AAAA record for elmasy.com\n")
	}

	for i := range r {
		t.Logf("elmasy.com AAAA -> %s\n", r[i])
	}
}

func TestIsSetAAAA(t *testing.T) {

	r, err := IsSetAAAA("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if r != true {
		t.Fatalf("TestAIsSet failed: elmasy.com is not set!\n")
	}
}

func BenchmarkQueryAAAA(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.QueryAAAA("example.com")
	}
}

func BenchmarkTryQueryAAAA(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.TryQueryAAAA("example.com")
	}
}

func BenchmarkTryQueryAAAAInvalid(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.TryQueryAAAA("invalid.example.com")
	}
}

func BenchmarkIsSetAAAA(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.IsSetAAAA("example.com")
	}
}
