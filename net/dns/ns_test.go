package dns

import "testing"

func TestQueryNS(t *testing.T) {

	r, err := QueryNS("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com NS -> %s\n", r[i])
	}
}

func TestTryQueryNS(t *testing.T) {

	r, err := TryQueryNS("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com NS -> %s\n", r[i])
	}
}

func TestIsSetNS(t *testing.T) {

	r, err := IsSetNS("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: NS is not set for elmasy.com\n")
	}
}
