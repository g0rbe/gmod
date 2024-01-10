package dns

import (
	"fmt"
	"net"

	mdns "github.com/miekg/dns"
)

var TypeA uint16 = 1

// QueryA ask the server and returns a slice of net.IP.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QueryA(name string) ([]net.IP, error) {

	rr, err := s.query(name, TypeA)
	if err != nil {
		return nil, err
	}

	r := make([]net.IP, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.A:
			r = append(r, v.A)
		case *mdns.CNAME:
			// Ignore CNAME
			continue
		case *mdns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryA ask a random server from servers and returns a slice of net.IP.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QueryA(name string) ([]net.IP, error) {

	return s.Get(-1).QueryA(name)
}

// QueryA ask a random server from DefaultServers and returns a slice of net.IP.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QueryA(name string) ([]net.IP, error) {

	return DefaultServers.QueryA(name)
}

// TryQueryA asks the servers for type A. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQueryA(name string) ([]net.IP, error) {

	rr, err := s.TryQuery(name, TypeA)
	if err != nil {
		return nil, err
	}

	r := make([]net.IP, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.A:
			r = append(r, v.A)
		case *mdns.CNAME:
			// Ignore CNAME
			continue
		case *mdns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// TryQueryA asks the DefaultServers for type A. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQueryA(name string) ([]net.IP, error) {

	return DefaultServers.TryQueryA(name)
}

// IsSetA checks whether an A type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetA(name string) (bool, error) {
	return s.IsSet(name, TypeA)
}

// IsSetA checks whether an A type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetA(name string) (bool, error) {
	return DefaultServers.IsSetA(name)
}
