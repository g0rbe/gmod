package dns

import "testing"

func TestQuerySOA(t *testing.T) {

	r, err := QuerySOA("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("elmasy.com SOA -> %s\n", r)
}

func TestTryQuerySOA(t *testing.T) {

	r, err := TryQuerySOA("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("elmasy.com SOA -> %s\n", r)
}

func TestIsSetSOA(t *testing.T) {

	r, err := IsSetSOA("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: SOA is not set for elmasy.com\n")
	}
}
