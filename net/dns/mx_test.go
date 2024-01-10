package dns

import "testing"

func TestQueryMX(t *testing.T) {

	r, err := QueryMX("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com MX -> %s\n", r[i])
	}
}

func TestTryQueryMX(t *testing.T) {

	r, err := TryQueryMX("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range r {
		t.Logf("elmasy.com MX -> %s\n", r[i])
	}
}

func TestIsSetMX(t *testing.T) {

	r, err := IsSetMX("elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: MX is not set for elmasy.com\n")
	}
}
