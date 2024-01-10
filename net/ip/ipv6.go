package ip

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"net/netip"
)

// ReservedIPv6 is a collection of reserved IPv6 addresses.
var ReservedIPv6 = []net.IPNet{
	{IP: net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}}, // ::/128, Unspecified Address
	{IP: net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}}, // ::1/128, Loopback Address
	{IP: net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0}},     // ::ffff:0:0/96, IPv4-mapped addresses
	{IP: net.IP{0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0}},     // ::ffff:0:0:0/96, IPv4 translated addresses
	{IP: net.IP{0, 100, 255, 155, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0}},   // 64:ff9b::/96, IPv4-IPv6 Translat.
	{IP: net.IP{0, 100, 255, 155, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},               // 64:ff9b:1::/48, IPv4-IPv6 Translat.
	{IP: net.IP{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0}},                 // 100::/64, Discard-Only Address Block
	{IP: net.IP{32, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                        // 2001::/32, IETF Protocol Assignments
	{IP: net.IP{32, 1, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 240, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                       // 2001:20::/28, ORCHIDv2
	{IP: net.IP{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                     // 2001:db8::/32, Documentation
	{IP: net.IP{32, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                            // 2002::/16, 6to4
	{IP: net.IP{252, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                             // fc00::/7, Unique-Local
	{IP: net.IP{254, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                         // fe80::/10, Link-Local Unicast
	{IP: net.IP{255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Mask: net.IPMask{255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},                             // ff00::/8, Multicast
}

var (
	// /128 network mask
	Mask128 = net.IPMask{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	// /127 network mask
	Mask127 = net.IPMask{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe}
)

// IsReserved6 checks if the given IPv6 address is reserved.
func IsReserved6(ip net.IP) bool {
	for i := range ReservedIPv6 {
		if ReservedIPv6[i].Contains(ip) {
			return true
		}
	}
	return false
}

// IsValid6 checks whether ip is valid IPv6 address.
// This function is a wrapper for github.com/elmasy-com/elnet/validator.IPv6
//
// Deprecated: Use github.com/elmasy-com/elnet/validator instead.
func IsValid6[T IPTypes](ip T) bool {

	switch t := any(ip).(type) {
	case net.IP:
		return t.To16() != nil
	case *net.IP:
		return t.To16() != nil
	case netip.Addr:
		return t.Is6()
	case *netip.Addr:
		return t.Is6()
	case string:
		i := net.ParseIP(t)
		if i == nil {
			return false
		}

		return i.To16() != nil
	case *string:
		i := net.ParseIP(*t)
		if i == nil {
			return false
		}

		return i.To16() != nil
	default:
		return false
	}
}

// GetRandom6 is return a random IPv6 address.
// The returned IP *can be* a reserved address.
func GetRandom6() net.IP {

	bytes := make([]byte, 16)

	rand.Read(bytes)

	return net.IP{
		bytes[0], bytes[1], bytes[2], bytes[3],
		bytes[4], bytes[5], bytes[6], bytes[7],
		bytes[8], bytes[9], bytes[10], bytes[11],
		bytes[12], bytes[13], bytes[14], bytes[15]}
}

// GetPublic6 is return a *non reserved* IPv6 address.
func GetPublic6() net.IP {

	for {
		ip := GetRandom6()

		if !IsReserved6(ip) {
			return ip
		}
	}
}

// GetList6 creates a list of IPv6 address on the given IPNet.
// Returns the first (identification address), the last (broadcast address) and a channel of usable IP addresses.
// If the mask is /128 the last and the usable channel is nil.
// If the mask is /127 the usable channel is nil.
func GetList6(n net.IPNet) (net.IP, net.IP, <-chan net.IP, error) {

	switch {
	case !IsValid6(n.IP):
		return nil, nil, nil, fmt.Errorf("invalid IPv6 address: %s", n.IP)
	case len(n.Mask) != net.IPv6len:
		return nil, nil, nil, fmt.Errorf("invalid IPv6 mask: %s", n.Mask)
	}

	first := make(net.IP, net.IPv6len)

	copy(first, n.IP.Mask(n.Mask))

	// IPv6-128 has only one IP address
	if bytes.Equal(n.Mask, Mask128) {
		return first, nil, nil, nil
	}

	last := make(net.IP, net.IPv6len)
	copy(last, first)

	// Find the last address
	ip := make(net.IP, net.IPv6len)
	copy(ip, first)

	for Increase(ip); n.Contains(ip); Increase(ip) {
		copy(last, ip)
	}

	// IPv6/127 has two IP address: first and last
	if bytes.Equal(n.Mask, Mask127) {
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
