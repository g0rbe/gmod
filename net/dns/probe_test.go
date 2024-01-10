package dns

import (
	"testing"
	"time"
)

func TestProbeUDP(t *testing.T) {

	e, err := Probe("udp", "1.1.1.1", "53", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: %s", err)
	}

	if !e {
		t.Fatalf("FAIL: UDP 1.1.1.1:53 should be a valid DNS server")
	}
}

func TestProbeTCP(t *testing.T) {

	e, err := Probe("tcp", "1.1.1.1", "53", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: %s", err)
	}

	if !e {
		t.Fatalf("FAIL: TCP 1.1.1.1:53 should be a valid DNS server")
	}
}

func TestProbeTCPTLS(t *testing.T) {

	e, err := Probe("tcp-tls", "1.1.1.1", "853", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: %s", err)
	}

	if !e {
		t.Fatalf("FAIL: TCP-TLS 1.1.1.1:853 should be a valid DNS server")
	}
}
