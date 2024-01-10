package route

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"

	gip "github.com/g0rbe/gmod/net/ip"
)

// See "man 8 route" for more info
type Route4 struct {
	Interface *net.Interface
	Dest      net.IP
	Gateway   net.IP
	Flags     uint64
	RefCnt    uint64
	Use       uint64
	Metric    uint64
	Mask      net.IPMask
	MTU       uint64
	Window    uint64
	IRTT      uint64
}

// IsUp checks whether the UP flag is set.
func (r *Route4) IsUp() bool {
	return r.Flags&RTF_UP != 0
}

// IsGateway check whether the GATEWAY flag is set.
func (r *Route4) IsGateway() bool {
	return r.Flags&RTF_GATEWAY != 0
}

func (r Route4) String() string {

	return fmt.Sprintf("%s\t%s\t%s\t%d\t%d\t%d\t%d\t%s\t%d\t%d\t%d",
		r.Interface.Name, r.Dest.String(), r.Gateway.String(),
		r.Flags, r.RefCnt, r.Use, r.Metric, r.Mask.String(),
		r.MTU, r.Window, r.IRTT)

}

// parseIPv4 parses a little-endian represented IP/Mask and return a big-endian byte slice with length of 4.
// The resulted slice can be transformed to net.IP with net.IPv4()
// or net.IPMask with net.IPv4Mask()
func parseIPv4(ip string) ([]byte, error) {

	if len(ip) != 8 {
		return nil, fmt.Errorf("invalid length: %d", len(ip))
	}

	ipb := []byte(ip)

	b4, err := strconv.ParseUint(string(ipb[:2]), 16, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse first byte: %s", err)
	}

	b3, err := strconv.ParseUint(string(ipb[2:4]), 16, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse second byte: %s", err)
	}

	b2, err := strconv.ParseUint(string(ipb[4:6]), 16, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse third byte: %s", err)
	}

	b1, err := strconv.ParseUint(string(ipb[6:8]), 16, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse fourth byte: %s", err)
	}

	return []byte{byte(b1), byte(b2), byte(b3), byte(b4)}, nil
}

func selectInterface(s []net.Interface, name string) *net.Interface {

	for i := range s {
		if s[i].Name == name {
			return &s[i]
		}
	}

	return nil
}

// GetRoutes4 parses /proc/net/route and return a slice of the routes.
func GetRoutes4() ([]Route4, error) {

	out, err := os.ReadFile("/proc/net/route")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	routes := make([]Route4, 0)

	devs, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %s", err)
	}

	for i := range lines {

		if lines[i] == "" {
			continue
		}

		fields := strings.Fields(lines[i])

		if len(fields) != 11 {
			return nil, fmt.Errorf("invalid number of fields: %d", len(fields))
		}
		if fields[0] == "Iface" {
			continue
		}

		route := Route4{}

		// Iface
		route.Interface = selectInterface(devs, fields[0])
		if route.Interface == nil {
			return nil, fmt.Errorf("interface not found: %s", fields[0])
		}

		// Destination
		dBytes, err := parseIPv4(fields[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse Destination: %s", err)
		}
		route.Dest = net.IPv4(dBytes[0], dBytes[1], dBytes[2], dBytes[3])

		// Gateway
		gBytes, err := parseIPv4(fields[2])
		if err != nil {
			return nil, fmt.Errorf("failed to parse Gateway: %s", err)
		}
		route.Gateway = net.IPv4(gBytes[0], gBytes[1], gBytes[2], gBytes[3])

		// Flags
		route.Flags, err = strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Flags: %s", err)
		}

		// RefCnt
		route.RefCnt, err = strconv.ParseUint(fields[4], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RefCnt: %s", err)
		}

		// Use
		route.Use, err = strconv.ParseUint(fields[5], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Use: %s", err)
		}

		// Metric
		route.Metric, err = strconv.ParseUint(fields[6], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Metric: %s", err)
		}

		// Mask
		mBytes, err := parseIPv4(fields[7])
		if err != nil {
			return nil, fmt.Errorf("failed to parse Mask: %s", err)
		}
		route.Mask = net.IPv4Mask(mBytes[0], mBytes[1], mBytes[2], mBytes[3])

		// MTU
		route.MTU, err = strconv.ParseUint(fields[8], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("faield to parse MTU: %s", err)
		}

		// Window
		route.Window, err = strconv.ParseUint(fields[9], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Window: %s", err)
		}

		// IRTT
		route.IRTT, err = strconv.ParseUint(fields[10], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse IRTT: %s", err)
		}

		routes = append(routes, route)
	}

	return routes, nil
}

// getDefaultGateway4 returns the gateway route with the lowest Metric.
func getDefaultGateway4() (Route4, error) {

	routes, err := GetRoutes4()
	if err != nil {
		return Route4{}, err
	}

	gws := make([]Route4, 0)

	for i := range routes {
		if routes[i].IsGateway() {
			gws = append(gws, routes[i])
		}
	}

	if len(gws) == 0 {
		return Route4{}, fmt.Errorf("not found")
	}

	sort.Slice(gws, func(i, j int) bool { return gws[i].Metric < gws[j].Metric })

	return gws[0], nil
}

// GetRoute4 selects the route for ip.
func GetRoute4(ip net.IP) (Route4, error) {

	if !gip.IsValid4(ip) {
		return Route4{}, fmt.Errorf("not valid IPv4 address: %s", ip)
	}

	routes, err := GetRoutes4()
	if err != nil {
		return Route4{}, err
	}

	// Possible routes
	pRoutes := make([]Route4, 0)

	// First step: find the interface for the given ip, skipping route with 0.0.0.0 as destination
	for i := range routes {

		// skip destinations that equals to  0.0.0.0
		if routes[i].Dest.Equal(net.IPv4(0x0, 0x0, 0x0, 0x0)) {
			continue
		}

		net := net.IPNet{IP: routes[i].Dest, Mask: routes[i].Mask}

		if net.Contains(ip) {
			pRoutes = append(pRoutes, routes[i])
		}
	}

	// ip is a global address
	if len(pRoutes) == 0 {
		return getDefaultGateway4()
	}

	// Sort the routes based on Metric
	sort.Slice(pRoutes, func(i, j int) bool { return pRoutes[i].Metric < pRoutes[j].Metric })

	return pRoutes[0], nil
}
