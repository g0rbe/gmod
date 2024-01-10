package dns

import "testing"

func TestQueryDNAME(t *testing.T) {

	r, err := QueryDNAME("design.ucla.edu")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("design.ucla.edu DNAME -> %s\n", r)

}

func TestTryQueryDNAME(t *testing.T) {

	r, err := TryQueryDNAME("design.ucla.edu")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("design.ucla.edu DNAME -> %s\n", r)

}

func TestIsSetDNAME(t *testing.T) {

	r, err := IsSetDNAME("design.ucla.edu")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: DNAME is not set for design.ucla.edu\n")
	}
}
