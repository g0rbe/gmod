package dns

import "testing"

func TestQuerySRV(t *testing.T) {

	r, err := QuerySRV("_sip._tls.tesla.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range r {
		t.Logf("_sip._tls.tesla.com SRV -> %s\n", r[i])
	}
}

func TestTryQuerySRV(t *testing.T) {

	r, err := TryQuerySRV("_sip._tls.tesla.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	for i := range r {
		t.Logf("_sip._tls.tesla.com SRV -> %s\n", r[i])
	}
}

func TestIsSetSRV(t *testing.T) {

	r, err := IsSetSRV("_sip._tls.tesla.com")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: SRV is not set for _sip._tls.tesla.com\n")
	}
}
