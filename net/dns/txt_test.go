package dns

import (
	"testing"
	"time"
)

func TestQueryTXT(t *testing.T) {

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		t.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	r, err := srvs.QueryTXT("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: invalid result length: %d\n", len(r))
	}

	for i := range r {
		t.Logf("elmasy.com TXT -> %s\n", r[i])
	}
}

func TestQueryTXTTruncated(t *testing.T) {

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		t.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	r, err := srvs.QueryTXT("dkim._domainkey.danielgorbe.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: invalid result length: %d\n", len(r))
	}

	for i := range r {
		t.Logf("elmasy.com TXT -> %s\n", r[i])
	}
}

func TestTryQueryTXT(t *testing.T) {

	r, err := TryQueryTXT("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: invalid result length: %d\n", len(r))
	}

	for i := range r {
		t.Logf("elmasy.com TXT -> %s\n", r[i])
	}
}

func TestTryQueryTXTTruncated(t *testing.T) {

	r, err := TryQueryTXT("dkim._domainkey.danielgorbe.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: invalid result length: %d\n", len(r))
	}

	for i := range r {
		t.Logf("elmasy.com TXT -> %s\n", r[i])
	}
}

func TestIsSetTXT(t *testing.T) {

	r, err := IsSetTXT("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: TXT is not set for elmasy.com\n")
	}
}

func TestIsSetTXTTruncated(t *testing.T) {

	r, err := IsSetTXT("dkim._domainkey.danielgorbe.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: TXT is not set for elmasy.com\n")
	}
}
