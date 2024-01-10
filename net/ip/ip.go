package ip

import (
	"fmt"
	"math/rand"
	"net"
)

type IPTypes interface {
	net.IP | *net.IP | net.IPNet | *net.IPNet | net.IPAddr | *net.IPAddr | net.TCPAddr | *net.TCPAddr | net.UDPAddr | *net.UDPAddr | string | *string
}

// Convert any IPType to net.IP.
// If not an IP, returns nil.
func ConvertToIP[T IPTypes](ip T) net.IP {

	switch v := any(ip).(type) {
	case net.IP:
		return v
	case *net.IP:
		return *v
	case net.IPNet:
		return v.IP
	case *net.IPNet:
		return v.IP
	case net.IPAddr:
		return v.IP
	case *net.IPAddr:
		return v.IP
	case net.TCPAddr:
		return v.IP
	case *net.TCPAddr:
		return v.IP
	case net.UDPAddr:
		return v.IP
	case *net.UDPAddr:
		return v.IP
	case string:
		return net.ParseIP(v)
	case *string:
		return net.ParseIP(*v)
	default:
		return nil
	}
}

// IsLAN checks whether the given IP is in your LAN's subnet by iteraet over the interfaces.
// It will panics on error.
func IsLAN(ip net.IP) bool {

	if ip == nil {
		panic("ip is nil")
	}

	devs, err := net.Interfaces()
	if err != nil {
		panic(fmt.Sprintf("failed to get interfaces: %s", err))
	}

	for i := range devs {

		nets, err := devs[i].Addrs()
		if err != nil {
			panic(fmt.Sprintf("failed to get addresses: %s", err))
		}

		for v := range nets {

			net, ok := nets[v].(*net.IPNet)
			if !ok {
				continue
			}

			if net.Contains(ip) {
				return true
			}
		}
	}

	return false
}

// IsValidIP checks whether ip is valid IP address.
func IsValid[T IPTypes](ip T) bool {
	return IsValid4(ip) || IsValid6(ip)
}

// IsReservedIP checks whether ip is in the reserved address range.
func IsReserved(ip net.IP) bool {
	return IsReserved4(ip) || IsReserved6(ip)
}

// GetRandom is return a random IP address.
// The returned IP *can be* a reserved address.
// The version of the IP protocol is random.
func GetRandom() net.IP {

	n := rand.Intn(2)

	if n == 0 {
		return GetRandom4()
	} else {
		return GetRandom6()
	}
}

// GetPublic is return a *non reserved* IP address.
// The version of the IP protocol is random.
func GetPublic() net.IP {

	n := rand.Intn(2)

	if n == 0 {
		return GetPublic4()
	} else {
		return GetPublic6()
	}
}

// Increase ip to the next address.
func Increase(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// GetList creates a list of address on the given IPNet.
// Returns the first (identification address), the last (broadcast address) and a channel of usable IP addresses.
// If the mask is /32 or /128, the last and the usable channel is nil.
// If the mask is /31 or /127, the usable channel is nil.
func GetList(n net.IPNet) (net.IP, net.IP, <-chan net.IP, error) {

	switch {
	case IsValid4(n.IP):
		return GetList4(n)
	case IsValid6(n.IP):
		return GetList6(n)
	default:
		return nil, nil, nil, fmt.Errorf("invalid IP address: %s", n.IP)
	}
}
