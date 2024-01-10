package arp

import (
	"net"
	"testing"
)

func TestGetARPCache(t *testing.T) {

	cache, err := GetARPCache(net.IPv4(10, 0, 0, 1), nil)
	if err != nil {
		t.Errorf("Failed to get cache: %s\n", err)
		return
	}

	t.Logf("%s\n", cache)
}
