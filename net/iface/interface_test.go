package iface

import (
	"net"
	"testing"
)

func TestGetIPNets(t *testing.T) {

	ifaces, err := net.Interfaces()
	if err != nil {
		t.Errorf("Failed to get interfaces: %s\n", err)
	}

	for i := range ifaces {

		nets, err := GetIPNets(&ifaces[i])
		if err != nil {
			t.Errorf("Failed to get nets for %s: %s\n", ifaces[i].Name, err)
		}

		for v := range nets {
			t.Logf("%s -> %s\n", ifaces[i].Name, nets[v].String())
		}
	}
}

func TestGetIPNets4(t *testing.T) {

	ifaces, err := net.Interfaces()
	if err != nil {
		t.Errorf("Failed to get interfaces: %s\n", err)
	}

	for i := range ifaces {

		nets, err := GetIPNets4(&ifaces[i])
		if err != nil {
			t.Errorf("Failed to get nets for %s: %s\n", ifaces[i].Name, err)
		}

		for v := range nets {
			t.Logf("%s -> %s\n", ifaces[i].Name, nets[v].String())
		}
	}
}

func TestGetIPNets6(t *testing.T) {

	ifaces, err := net.Interfaces()
	if err != nil {
		t.Errorf("Failed to get interfaces: %s\n", err)
	}

	for i := range ifaces {

		nets, err := GetIPNets6(&ifaces[i])
		if err != nil {
			t.Errorf("Failed to get nets for %s: %s\n", ifaces[i].Name, err)
		}

		for v := range nets {
			t.Logf("%s -> %s\n", ifaces[i].Name, nets[v].String())
		}
	}
}
