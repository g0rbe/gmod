package dns

import (
	"errors"
	"testing"
	"time"
)

func TestQueryA(t *testing.T) {

	TestDomain := "elmasy.com"

	r, err := QueryA(TestDomain)
	if err != nil {
		t.Fatalf("FAIL: Failed to find A records for %s: %s\n", TestDomain, err)
	}

	for i := range r {
		t.Logf("%s A -> %s\n", TestDomain, r[i])
	}
}

func TestQueryAInvalid(t *testing.T) {

	TestDomain := "invalid.elmasy.com"

	r, err := QueryA(TestDomain)
	if err == nil {
		t.Fatalf("FAIL: error must be NXDOMAIN, got nil\n")
	}
	if !errors.Is(err, ErrName) {
		t.Fatalf("FAIL: error want: %s, error got: %s\n", ErrName, err)
	}

	if len(r) != 0 {
		t.Fatalf("FAIL: invalid result length: %d\n", len(r))
	}
}

func TestQueryALenZero(t *testing.T) {

	TestDomain := "_dmarc.elmasy.com"

	r, err := DefaultServers.QueryA(TestDomain)
	if err != nil {
		t.Fatalf("FAIL:Failed to finf A record for %s: %s\n", TestDomain, err)
	}

	if len(r) != 0 {
		t.Fatalf("FAIL: invalid result length: %d\n", len(r))
	}
}

func TestTryQueryA(t *testing.T) {

	TestDomain := "elmasy.com"

	r, err := TryQueryA(TestDomain)
	if err != nil {
		t.Fatalf("FAIL: Failed to find A record for %s: %s\n", TestDomain, err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: No A record for %s\n", TestDomain)
	}

	for i := range r {
		t.Logf("%s A -> %s\n", TestDomain, r[i])
	}
}

func TestTryQueryAInvalid(t *testing.T) {

	TestDomain := "invalid.elmasy.com"

	r, err := TryQueryA(TestDomain)
	if err == nil {
		t.Fatalf("FAIL: error is nil for %s\n", TestDomain)
	}

	if !errors.Is(err, ErrName) {
		t.Fatalf("FAIL: error for %s: %s, want: %s\n", TestDomain, err, ErrName)
	}

	if len(r) != 0 {
		t.Fatalf("FAIL: invalid result length: %d\n", len(r))
	}
}

func TestQueryARetryInvalidMaxRetries(t *testing.T) {

	TestDomain := "elmasy.com"
	DefaultServers.SetMaxRetries(0)

	r, err := TryQueryA(TestDomain)
	if err == nil {
		t.Fatalf("FAIL: err is nil, want: %s\n", ErrInvalidMaxRetries)
	}

	if !errors.Is(err, ErrInvalidMaxRetries) {
		t.Fatalf("FAIL: error want: %s, error got: %s\n", ErrInvalidMaxRetries, err)
	}

	if len(r) != 0 {
		t.Fatalf("FAIL: A record found for %s with 0 max retries\n", TestDomain)
	}

	DefaultServers.SetMaxRetries(5)
}

func TestIsSetA(t *testing.T) {

	TestDomain := "elmasy.com"

	r, err := IsSetA(TestDomain)
	if err != nil {
		t.Fatalf("FAIL: Failed to find A record for %s: %s\n", TestDomain, err)
	}

	if r != true {
		t.Fatalf("FAIL: A record for %s is not set!\n", TestDomain)
	}
}

func BenchmarkQueryA(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.QueryA("example.com")
	}
}

func BenchmarkQueryAInvalid(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.QueryA("invalid.example.com")
	}
}

func BenchmarkTryQueryA(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.TryQueryA("example.com")
	}
}

func BenchmarkTryQueryAInvalid(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.TryQueryA("invalid.example.com")
	}
}

func BenchmarkIsSetA(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.IsSetA("example.com")
	}
}

func BenchmarkIsSetAInvalid(b *testing.B) {

	// Sleep 2 sec to not overflow the DNS server
	time.Sleep(2 * time.Second)

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		b.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srvs.IsSetA("invalid.example.com")
	}
}
