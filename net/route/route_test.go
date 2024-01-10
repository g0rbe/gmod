package route

import (
	"net"
	"testing"
)

func TestParseIPv4(t *testing.T) {

	ipBytes, err := parseIPv4("00FFFFFF")
	if err != nil {
		t.Errorf("Parsing failed: %s\n", err)
	}

	ip := net.IPv4(ipBytes[0], ipBytes[1], ipBytes[2], ipBytes[3])
	if err != nil {
		t.Errorf("Failed to convert byte array to net.IP: %s\n", err)
	}

	if ip.String() != "255.255.255.0" {
		t.Errorf("Invalid result: %s\n", ip.String())
	}
}

func TestGetRoutes4(t *testing.T) {

	routes, err := GetRoutes4()
	if err != nil {
		t.Errorf("Parsing failed: %s\n", err)
	}

	for i := range routes {
		t.Logf("%s\n", routes[i])
	}
}

func TestGetRoute4(t *testing.T) {

	ip := net.ParseIP("1.1.1.1")
	if ip == nil {
		t.Errorf("Failed to parse ip\n")
	}

	route, err := GetRoute4(ip)
	if err != nil {
		t.Errorf("Failed to select route: %s\n", err)
	}

	t.Logf("%s\n", route)
}
