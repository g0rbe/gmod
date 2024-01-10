package dns

import "testing"

func TestQueryCNAME(t *testing.T) {

	r, err := QueryCNAME("autodiscover.elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: No CNAME record for autodiscover.elmasy.com\n")
	}

	for i := range r {
		t.Logf("autodiscover.elmasy.com CNAME -> %s\n", r[i])
	}
}

func TestTryQueryCNAME(t *testing.T) {

	r, err := TryQueryCNAME("autodiscover.elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range r {
		t.Logf("autodiscover.elmasy.com CNAME -> %s\n", r[i])
	}
}

func TestIsSetCNAME(t *testing.T) {

	r, err := IsSetCNAME("autodiscover.elmasy.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: CNAME is not set for autodiscover.elmasy.com\n")
	}
}
