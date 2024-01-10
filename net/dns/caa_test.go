package dns

import "testing"

func TestQueryCAA(t *testing.T) {

	r, err := QueryCAA("github.com")
	if err != nil {
		t.Fatalf("FAIL: Faile to query github.com CAA: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: No CAA record for github.com\n")
	}

	for i := range r {
		t.Logf("github.com CAA -> %s\n", r[i])
	}
}

func TestTryQueryCAA(t *testing.T) {

	r, err := TryQueryCAA("github.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: No CAA record for github.com\n")
	}

	for i := range r {
		t.Logf("github.com CAA -> %s\n", r[i])
	}
}

func TestIsSetCAA(t *testing.T) {

	r, err := IsSetCAA("github.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: CAA is not set for github.com\n")
	}
}
