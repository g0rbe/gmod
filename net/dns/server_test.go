package dns

import (
	"strings"
	"testing"
	"time"

	mdns "github.com/miekg/dns"
)

func TestNewServer(t *testing.T) {

	cases := []struct {
		Protocol string
		IP       string
		Port     string
		S        Server
		Err      string
	}{
		{Protocol: "udp", IP: "127.0.0.1", Port: "53", S: Server{Protocol: "udp", IP: "127.0.0.1", Port: "53", family: 4}, Err: ""},
		{Protocol: "tcp", IP: "127.0.0.1", Port: "53", S: Server{Protocol: "tcp", IP: "127.0.0.1", Port: "53", family: 4}, Err: ""},
		{Protocol: "tcp-tls", IP: "127.0.0.1", Port: "53", S: Server{Protocol: "tcp-tls", IP: "127.0.0.1", Port: "53", family: 4}, Err: ""},

		{Protocol: "udp", IP: "::1", Port: "53", S: Server{Protocol: "udp", IP: "::1", Port: "53", family: 6}, Err: ""},
		{Protocol: "tcp", IP: "::1", Port: "53", S: Server{Protocol: "tcp", IP: "::1", Port: "53", family: 6}, Err: ""},
		{Protocol: "tcp-tls", IP: "::1", Port: "53", S: Server{Protocol: "tcp-tls", IP: "::1", Port: "53", family: 6}, Err: ""},

		{Protocol: "udp", IP: "::1", Port: "853", S: Server{Protocol: "udp", IP: "::1", Port: "853", family: 6}, Err: ""},
		{Protocol: "tcp", IP: "::1", Port: "853", S: Server{Protocol: "tcp", IP: "::1", Port: "853", family: 6}, Err: ""},
		{Protocol: "tcp-tls", IP: "::1", Port: "853", S: Server{Protocol: "tcp-tls", IP: "::1", Port: "853", family: 6}, Err: ""},

		{Protocol: "", IP: "127.0.0.1", Port: "53", S: Server{}, Err: "protocol is empty"},
		{Protocol: "upd", IP: "127.0.0.1", Port: "53", S: Server{}, Err: "invalid protocol"},

		{Protocol: "tcp", IP: "", Port: "53", S: Server{}, Err: "ip is empty"},
		{Protocol: "tcp", IP: "127.0.0", Port: "53", S: Server{}, Err: "invalid ip"},

		{Protocol: "tcp", IP: "127.0.0.1", Port: "", S: Server{}, Err: "port is empty"},
		{Protocol: "tcp", IP: "127.0.0.1", Port: "-1", S: Server{}, Err: "invalid port"},
	}

	for i := range cases {

		s, err := NewServer(cases[i].Protocol, cases[i].IP, cases[i].Port, 2*time.Second)

		// Want error but not gor
		if cases[i].Err != "" && err == nil {
			t.Fatalf("FAIL: %s://%s:%s error wanted: %s, error got: nil\n", cases[i].Protocol, cases[i].IP, cases[i].Port, cases[i].Err)
			continue

		}

		// Check returned error
		if err != nil {

			if !strings.Contains(err.Error(), cases[i].Err) {
				t.Fatalf("FAIL: %s://%s:%s error wanted: %s, error got: %s\n", cases[i].Protocol, cases[i].IP, cases[i].Port, cases[i].Err, err)
			}

			continue
		}

		if cases[i].S.Protocol != s.Protocol {
			t.Fatalf("FAIL: %s://%s:%s proto wanted: %s, proto got: %s\n", cases[i].Protocol, cases[i].IP, cases[i].Port, cases[i].S.Protocol, s.Protocol)
		}

		if cases[i].S.IP != s.IP {
			t.Fatalf("FAIL: %s://%s:%s IP wanted: %s, IP got: %s\n", cases[i].Protocol, cases[i].IP, cases[i].Port, cases[i].S.IP, s.IP)
		}

		if cases[i].S.Port != s.Port {
			t.Fatalf("FAIL: %s://%s:%s port wanted: %s, port got: %s\n", cases[i].Protocol, cases[i].IP, cases[i].Port, cases[i].S.Port, s.Port)
		}

		if cases[i].S.family != s.family {
			t.Fatalf("FAIL: %s://%s:%s family wanted: %d, family got: %d\n", cases[i].Protocol, cases[i].IP, cases[i].Port, cases[i].S.family, s.family)
		}
	}
}

func TestNewServerStr(t *testing.T) {

	cases := []struct {
		V   string
		S   Server
		Err string
	}{
		{V: "udp://127.0.0.1:53", S: Server{Protocol: "udp", IP: "127.0.0.1", Port: "53", family: 4}, Err: ""},
		{V: "tcp://127.0.0.1:53", S: Server{Protocol: "tcp", IP: "127.0.0.1", Port: "53", family: 4}, Err: ""},
		{V: "tcp-tls://127.0.0.1:53", S: Server{Protocol: "tcp-tls", IP: "127.0.0.1", Port: "53", family: 4}, Err: ""},
		{V: "127.0.0.1:53", S: Server{Protocol: "udp", IP: "127.0.0.1", Port: "53", family: 4}, Err: ""},
		{V: "udp://127.0.0.1", S: Server{Protocol: "udp", IP: "127.0.0.1", Port: "53", family: 4}, Err: ""},
		{V: "127.0.0.1", S: Server{Protocol: "udp", IP: "127.0.0.1", Port: "53", family: 4}, Err: ""},
		{V: "8.8.8.8", S: Server{Protocol: "udp", IP: "8.8.8.8", Port: "53", family: 4}, Err: ""},

		{V: "upd://127.0.0.1:53", S: Server{}, Err: "invalid protocol"},
		{V: "udp://127.0.0:53", S: Server{}, Err: "invalid ip"},
		{V: "udp://127.0.0.1:-1", S: Server{}, Err: "invalid port"},

		{V: "udp://[::1]:53", S: Server{Protocol: "udp", IP: "::1", Port: "53", family: 6}, Err: ""},
		{V: "tcp://[::1]]:53", S: Server{Protocol: "tcp", IP: "::1", Port: "53", family: 6}, Err: ""},
		{V: "tcp-tls://[::1]:53", S: Server{Protocol: "tcp-tls", IP: "::1", Port: "53", family: 6}, Err: ""},
		{V: "[::1]:53", S: Server{Protocol: "udp", IP: "::1", Port: "53", family: 6}, Err: ""},
		{V: "udp://[::1]", S: Server{Protocol: "udp", IP: "::1", Port: "53", family: 6}, Err: ""},
		{V: "[::1]", S: Server{Protocol: "udp", IP: "::1", Port: "53", family: 6}, Err: ""},

		{V: "upd://[::1]:53", S: Server{}, Err: "invalid protocol"},
		{V: "udp://[::-1]:53", S: Server{}, Err: "invalid ip"},
		{V: "udp://[::1]:-1", S: Server{}, Err: "invalid port"},
	}

	for i := range cases {

		s, err := NewServerStr(cases[i].V, 2*time.Second)

		// Want error but not gor
		if cases[i].Err != "" && err == nil {
			t.Fatalf("FAIL: %s error wanted: %s, error got: nil\n", cases[i].V, cases[i].Err)
			continue

		}

		if err != nil {

			if !strings.Contains(err.Error(), cases[i].Err) {
				t.Fatalf("FAIL: %s error wanted: %s, error got: %s\n", cases[i].V, cases[i].Err, err)
			}

			continue
		}

		if cases[i].S.Protocol != s.Protocol {
			t.Fatalf("FAIL: %s proto wanted: %s, proto got: %s\n", cases[i].V, cases[i].S.Protocol, s.Protocol)
		}

		if cases[i].S.IP != s.IP {
			t.Fatalf("FAIL: %s IP wanted: %s, IP got: %s\n", cases[i].V, cases[i].S.IP, s.IP)
		}

		if cases[i].S.Port != s.Port {
			t.Fatalf("FAIL: %s port wanted: %s, port got: %s\n", cases[i].V, cases[i].S.Port, s.Port)
		}

		if cases[i].S.family != s.family {
			t.Fatalf("FAIL: %s family wanted: %d, family got: %d\n", cases[i].V, cases[i].S.family, s.family)
		}

	}
}

func TestServerString(t *testing.T) {

	cases := []struct {
		Protocol string
		IP       string
		Port     string
		S        string
	}{
		{Protocol: "udp", IP: "127.0.0.1", Port: "53", S: "udp://127.0.0.1:53"},
		{Protocol: "tcp", IP: "127.0.0.1", Port: "53", S: "tcp://127.0.0.1:53"},
		{Protocol: "tcp-tls", IP: "127.0.0.1", Port: "53", S: "tcp-tls://127.0.0.1:53"},

		{Protocol: "udp", IP: "::1", Port: "53", S: "udp://[::1]:53"},
		{Protocol: "tcp", IP: "::1", Port: "53", S: "tcp://[::1]:53"},
		{Protocol: "tcp-tls", IP: "::1", Port: "53", S: "tcp-tls://[::1]:53"},

		{Protocol: "udp", IP: "::1", Port: "853", S: "udp://[::1]:853"},
		{Protocol: "tcp", IP: "::1", Port: "853", S: "tcp://[::1]:853"},
		{Protocol: "tcp-tls", IP: "::1", Port: "853", S: "tcp-tls://[::1]:853"},
	}

	for i := range cases {

		s, err := NewServer(cases[i].Protocol, cases[i].IP, cases[i].Port, 2*time.Second)

		// Check returned error
		if err != nil {

			t.Fatalf("FAIL: Failed to create new server for %s://%s:%s: %s\n", cases[i].Protocol, cases[i].IP, cases[i].Port, err)

			continue
		}

		if cases[i].S != s.String() {
			t.Fatalf("FAIL: %s://%s:%s string wanted: %s, string got: %s\n", cases[i].Protocol, cases[i].IP, cases[i].Port, cases[i].S, s.String())
		}
	}
}

func TestServerServer(t *testing.T) {

	cases := []struct {
		Protocol string
		IP       string
		Port     string
		S        string
	}{
		{Protocol: "udp", IP: "127.0.0.1", Port: "53", S: "127.0.0.1:53"},
		{Protocol: "tcp", IP: "127.0.0.1", Port: "53", S: "127.0.0.1:53"},
		{Protocol: "tcp-tls", IP: "127.0.0.1", Port: "53", S: "127.0.0.1:53"},

		{Protocol: "udp", IP: "::1", Port: "53", S: "[::1]:53"},
		{Protocol: "tcp", IP: "::1", Port: "53", S: "[::1]:53"},
		{Protocol: "tcp-tls", IP: "::1", Port: "53", S: "[::1]:53"},

		{Protocol: "udp", IP: "::1", Port: "853", S: "[::1]:853"},
		{Protocol: "tcp", IP: "::1", Port: "853", S: "[::1]:853"},
		{Protocol: "tcp-tls", IP: "::1", Port: "853", S: "[::1]:853"},
	}

	for i := range cases {

		s, err := NewServer(cases[i].Protocol, cases[i].IP, cases[i].Port, 2*time.Second)

		// Check returned error
		if err != nil {

			t.Fatalf("FAIL: Failed to create new server for %s://%s:%s: %s\n", cases[i].Protocol, cases[i].IP, cases[i].Port, err)

			continue
		}

		if cases[i].S != s.Server() {
			t.Fatalf("FAIL: %s://%s:%s server wanted: %s, server got: %s\n", cases[i].Protocol, cases[i].IP, cases[i].Port, cases[i].S, s.Server())
		}
	}
}

func TestServerQueryAUDP(t *testing.T) {

	s, err := NewServer("udp", "1.1.1.1", "53", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: Failed to create server: %s\n", err)
	}

	rr, err := s.query("example.com", TypeA)
	if err != nil {
		t.Fatalf("FAIL: Failed to get A record for example.com: %s\n", err)
	}

	if len(rr) == 0 {
		t.Fatalf("FAIL: Failed to get A record for example.com: no record\n")
	}

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.A:
			t.Logf("example.com A -> %s", v.A)
		case *mdns.CNAME:
			t.Logf("example.com CNAME -> %s", v.Target)
		case *mdns.DNAME:
			t.Logf("example.com DNAME -> %s", v.Target)
		default:
			t.Fatalf("FAIL: Invalid type found for example.com A record: %T\n", v)
		}
	}
}

func TestServerQueryATCP(t *testing.T) {

	s, err := NewServer("tcp", "1.1.1.1", "53", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: Failed to create server: %s\n", err)
	}

	rr, err := s.query("example.com", TypeA)
	if err != nil {
		t.Fatalf("FAIL: Failed to get A record for example.com: %s\n", err)
	}

	if len(rr) == 0 {
		t.Fatalf("FAIL: Failed to get A record for example.com: no record\n")
	}

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.A:
			t.Logf("example.com A -> %s", v.A)
		case *mdns.CNAME:
			t.Logf("example.com CNAME -> %s", v.Target)
		case *mdns.DNAME:
			t.Logf("example.com DNAME -> %s", v.Target)
		default:
			t.Fatalf("FAIL: Invalid type found for example.com A record: %T\n", v)
		}
	}
}

func TestServerQueryATCPTLS(t *testing.T) {

	s, err := NewServer("tcp-tls", "1.1.1.1", "853", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: Failed to create server: %s\n", err)
	}

	rr, err := s.query("example.com", TypeA)
	if err != nil {
		t.Fatalf("FAIL: Failed to get A record for example.com: %s\n", err)
	}

	if len(rr) == 0 {
		t.Fatalf("FAIL: Failed to get A record for example.com: no record\n")
	}

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.A:
			t.Logf("example.com A -> %s", v.A)
		case *mdns.CNAME:
			t.Logf("example.com CNAME -> %s", v.Target)
		case *mdns.DNAME:
			t.Logf("example.com DNAME -> %s", v.Target)
		default:
			t.Fatalf("FAIL: Invalid type found for example.com A record: %T\n", v)
		}
	}
}
