package dns

import (
	"testing"
	"time"
)

func TestNewServersFromSlice(t *testing.T) {

	one, err := NewServer("udp", "8.8.8.8", "53", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: Failed to create udp://8.8.8.8:53: %s\n", err)
	}

	two, err := NewServer("udp", "1.1.1.1", "53", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: Failed to create udp://1.1.1.1:53: %s\n", err)
	}

	three, err := NewServer("udp", "9.9.9.9", "53", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: Failed to create udp://9.9.9.9:53: %s\n", err)
	}

	srvs := NewServersSlice(3, one, two, three)

	_, err = srvs.QueryA("example.com")
	if err != nil {
		t.Fatalf("FAIL: Failed to query example.com: %s\n", err)
	}
}

func TestNewServersFromIPs(t *testing.T) {

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		t.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	_, err = srvs.QueryA("example.com")
	if err != nil {
		t.Fatalf("FAIL: Failed to query example.com: %s\n", err)
	}
}

func TestNewServersAppend(t *testing.T) {

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		t.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	one, err := NewServer("udp", "8.8.8.8", "53", 2*time.Second)
	if err != nil {
		t.Fatalf("FAIL: Failed to create udp://8.8.8.8:53: %s\n", err)
	}

	srvs.Append(one)

	if len(srvs.srvs) != 4 {
		t.Fatalf("FAIL: Failed to append: length is %d\n", len(srvs.srvs))
	}

	_, err = srvs.QueryA("example.com")
	if err != nil {
		t.Fatalf("FAIL: Failed to query example.com: %s\n", err)
	}
}

func TestNewServersGet(t *testing.T) {

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		t.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	if srv := srvs.Get(-1); srv == nil {
		t.Fatalf("FAIL: Failed to get server with index -1\n")
	}

	if srv := srvs.Get(1); srv == nil {
		t.Fatalf("FAIL: Failed to get server with index 1\n")
	}

	if srv := srvs.Get(4); srv == nil {
		t.Fatalf("FAIL: Failed to get server with index 4\n")
	}

	_, err = srvs.QueryA("example.com")
	if err != nil {
		t.Fatalf("FAIL: Failed to query example.com: %s\n", err)
	}
}

func TestNewServersGetMaxRetries(t *testing.T) {

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		t.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	if srvs.GetMaxRetries() != 3 {
		t.Fatalf("FAIL: Invalid number of max retries: %d, want: 3\n", srvs.GetMaxRetries())
	}
}

func TestNewServersSetMaxRetries(t *testing.T) {

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		t.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	if srvs.GetMaxRetries() != 3 {
		t.Fatalf("FAIL: Invalid number of max retries: %d, want: 3\n", srvs.GetMaxRetries())
	}

	srvs.SetMaxRetries(10)

	if srvs.GetMaxRetries() != 10 {
		t.Fatalf("FAIL: Invalid number of max retries: %d, want: 10\n", srvs.GetMaxRetries())
	}
}

func TestNewServersTryQuery(t *testing.T) {

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		t.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	r, err := srvs.TryQuery("example.com", TypeA)
	if err != nil {
		t.Fatalf("FAIL: Failed to query example.com: %s\n", err)
	}

	if len(r) == 0 {
		t.Fatalf("FAIL: Invalid result length: %d\n", len(r))
	}
}

func TestNewServersIsSet(t *testing.T) {

	srvs, err := NewServersStr(3, 1*time.Second, "8.8.8.8", "8.8.8.8", "8.8.8.8")
	if err != nil {
		t.Fatalf("FAIL: Failed to create servers: %s\n", err)
	}

	r, err := srvs.IsSet("example.com", TypeA)
	if err != nil {
		t.Fatalf("FAIL: Failed to query example.com: %s\n", err)
	}

	if !r {
		t.Fatalf("FAIL: A record for example.com is not set!\n")
	}
}
