// Query local network intefaces.
package iface

import (
	"fmt"
	"net"

	"github.com/g0rbe/gmod/net/ip"
)

// Interface flags from net/if.h and uapi/linux/if.h.
// See more at "man 7 netdevice"
const (
	// Interface is up.
	IFF_UP = 0x1

	// Broadcast address valid.
	IFF_BROADCAST = 0x2

	// Turn on debugging.
	IFF_DEBUG = 0x4

	// Is a loopback net.
	IFF_LOOPBACK = 0x8

	// Interface is point-to-point link.
	IFF_POINTOPOINT = 0x10

	// Avoid use of trailers.
	IFF_NOTRAILERS = 0x20

	// Resources allocated.
	IFF_RUNNING = 0x40

	// No address resolution protocol.
	IFF_NOARP = 0x80

	// Receive all packets.
	IFF_PROMISC = 0x100

	// Receive all multicast packets.
	IFF_ALLMULTI = 0x200

	// Master of a load balancer.
	IFF_MASTER = 0x400

	// Slave of a load balancer.
	IFF_SLAVE = 0x800

	// Supports multicast.
	IFF_MULTICAST = 0x1000

	// Can set media type.
	IFF_PORTSEL = 0x2000

	// Auto media select active.
	IFF_AUTOMEDIA = 0x4000

	// Dialup device with changing addresses.
	IFF_DYNAMIC = 0x8000

	// Driver signals L1 up
	IFF_LOWER_UP = 0x10000

	// Driver signals dormant
	IFF_DORMANT = 0x20000

	// Echo sent packets
	IFF_ECHO = 0x40000
)

// GetIPNets returns every net.IPNet address associated with iface.
func GetIPNets(iface *net.Interface) ([]*net.IPNet, error) {

	if iface == nil {
		return nil, fmt.Errorf("iface is nil")
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	res := make([]*net.IPNet, 0)

	for i := range addrs {

		if v, ok := addrs[i].(*net.IPNet); !ok {
			continue
		} else {
			res = append(res, v)
		}
	}

	return res, nil
}

// GetIPNets4 returns every IPv4 address associated with iface.
func GetIPNets4(iface *net.Interface) ([]*net.IPNet, error) {

	if iface == nil {
		return nil, fmt.Errorf("iface is nil")
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	res := make([]*net.IPNet, 0)

	for i := range addrs {

		if v, ok := addrs[i].(*net.IPNet); !ok {
			continue
		} else {
			if !ip.IsValid4(v) {
				continue
			}
			res = append(res, v)
		}
	}

	return res, nil
}

// GetIPNets6 returns every IPv6 address associated with iface.
func GetIPNets6(iface *net.Interface) ([]*net.IPNet, error) {

	if iface == nil {
		return nil, fmt.Errorf("iface is nil")
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	res := make([]*net.IPNet, 0)

	for i := range addrs {

		if v, ok := addrs[i].(*net.IPNet); !ok {
			continue
		} else {
			if !ip.IsValid6(v) {
				continue
			}
			res = append(res, v)
		}
	}

	return res, nil
}
