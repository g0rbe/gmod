package ip

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"net/netip"
)

// ReservedIPv4 is a collection of reserved IPv4 addresses.
// Source: https://www.iana.org/assignments/iana-ipv4-special-registry/iana-ipv4-special-registry.xhtml
var ReservedIPv4 = []net.IPNet{
	{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPv4Mask(255, 0, 0, 0)},               // 0.0.0.0/8, "This" network
	{IP: net.IPv4(10, 0, 0, 0), Mask: net.IPv4Mask(255, 0, 0, 0)},              // 10.0.0.0/8, Class A private network
	{IP: net.IPv4(100, 64, 0, 0), Mask: net.IPv4Mask(255, 192, 0, 0)},          // 100.64.0.0/10, Carrier-grade NAT
	{IP: net.IPv4(127, 0, 0, 0), Mask: net.IPv4Mask(255, 0, 0, 0)},             // 127.0.0.0/8, Loopback
	{IP: net.IPv4(169, 254, 0, 0), Mask: net.IPv4Mask(255, 255, 0, 0)},         // 169.254.0.0/16, Link local
	{IP: net.IPv4(172, 16, 0, 0), Mask: net.IPv4Mask(255, 240, 0, 0)},          // 172.16.0.0/12, Class B private network
	{IP: net.IPv4(192, 0, 0, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},         // 192.0.0.0/24, IETF protocol assignments
	{IP: net.IPv4(192, 0, 2, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},         // 192.0.2.0/24, TEST-NET-1
	{IP: net.IPv4(192, 88, 99, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},       // 192.88.99.0/24, Reserved, formerly IPv6 to IPv4
	{IP: net.IPv4(192, 168, 0, 0), Mask: net.IPv4Mask(255, 255, 0, 0)},         // 192.168.0.0/24, Class C private network
	{IP: net.IPv4(198, 18, 0, 0), Mask: net.IPv4Mask(255, 254, 0, 0)},          // 198.18.0.0/15, Benchmarking
	{IP: net.IPv4(198, 51, 100, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},      // 198.51.100.0/24, TEST-NET-2
	{IP: net.IPv4(203, 0, 113, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},       // 203.0.113.0/24, TEST-NET-3
	{IP: net.IPv4(224, 0, 0, 0), Mask: net.IPv4Mask(240, 0, 0, 0)},             // 224.0.0.0/4, Multicast
	{IP: net.IPv4(233, 252, 0, 0), Mask: net.IPv4Mask(255, 255, 255, 0)},       // 233.252.0.0/24 , MCAST-TEST-NET
	{IP: net.IPv4(240, 0, 0, 0), Mask: net.IPv4Mask(240, 0, 0, 0)},             // 240.0.0.0/4, Reserved for future use
	{IP: net.IPv4(255, 255, 255, 255), Mask: net.IPv4Mask(255, 255, 255, 255)}, // 255.255.255.255/32, Broadcast
}

var (
	// /32 network mask
	Mask32 = net.IPMask{0xff, 0xff, 0xff, 0xff}
	// /31 network mask
	Mask31 = net.IPMask{0xff, 0xff, 0xff, 0xfe}
)

// IsReserved4 checks if the given IPv4 address is reserved.
func IsReserved4(ip net.IP) bool {
	for i := range ReservedIPv4 {
		if ReservedIPv4[i].Contains(ip) {
			return true
		}
	}
	return false
}

// IsValid4 checks whether ip is valid IPv4 address.
func IsValid4[T IPTypes](ip T) bool {

	switch t := any(ip).(type) {
	case net.IP:
		return t.To4() != nil
	case *net.IP:
		return t.To4() != nil
	case netip.Addr:
		return t.Is4()
	case *netip.Addr:
		return t.Is4()
	case string:
		i := net.ParseIP(t)
		if i == nil {
			return false
		}

		return i.To4() != nil
	case *string:
		i := net.ParseIP(*t)
		if i == nil {
			return false
		}

		return i.To4() != nil
	default:
		return false
	}
}

// GetRandom4 is return a random IPv4 address.
// The returned IP *can be* a reserved address.
func GetRandom4() net.IP {

	bytes := make([]byte, 4)

	rand.Read(bytes)

	return net.IPv4(bytes[0], bytes[1], bytes[2], bytes[3])
}

// GetPublic4 is return a *non reserved* IPv4 address.
func GetPublic4() net.IP {

	for {
		ip := GetRandom4()

		if !IsReserved4(ip) {
			return ip
		}
	}
}

// GetList4 creates a list of IPv4 address on the given IPNet.
// Returns the first (identification address), the last (broadcast address) and a channel of usable IP addresses.
// If the mask is /32 the last and the usable channel is nil.
// If the mask is /31 the usable channel is nil.
func GetList4(n net.IPNet) (net.IP, net.IP, <-chan net.IP, error) {

	switch {
	case !IsValid4(n.IP):
		return nil, nil, nil, fmt.Errorf("invalid IPv4 address: %s", n.IP)
	case len(n.Mask) != net.IPv4len:
		return nil, nil, nil, fmt.Errorf("invalid IPv4 mask: %s", n.Mask)
	}

	first := make(net.IP, net.IPv4len)

	copy(first, n.IP.Mask(n.Mask))

	// IPv4/32 has only one IP address
	if bytes.Equal(n.Mask, Mask32) {
		return first, nil, nil, nil
	}

	last := make(net.IP, net.IPv4len)
	copy(last, first)

	// Find the last address
	ip := make(net.IP, net.IPv4len)
	copy(ip, first)

	for Increase(ip); n.Contains(ip); Increase(ip) {
		copy(last, ip)
	}

	// IPv4/31 has two IP address: first and last
	if bytes.Equal(n.Mask, Mask31) {
		return first, last, nil, nil
	}

	/*
		Use channel because of memory allocation: for example 10.0.0.0/8 allocate too much memory at once (not to mention IPv6).
		With unbuffered channel, one can iterate over it with range.
	*/

	// Start again from the first address.
	copy(ip, first)

	// TODO: Buffered channel?
	usable := make(chan net.IP)

	go func() {
		// Increase ip until the last address
		for Increase(ip); !ip.Equal(last); Increase(ip) {
			v := make(net.IP, len(ip))
			copy(v, ip)

			usable <- v
		}
		close(usable)
	}()

	return first, last, usable, nil
}
