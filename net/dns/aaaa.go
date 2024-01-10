package dns

import (
	"fmt"
	"net"

	mdns "github.com/miekg/dns"
)

var TypeAAAA uint16 = 28

// QueryAAAA ask the server and returns a slice of net.IP.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QueryAAAA(name string) ([]net.IP, error) {

	rr, err := s.query(name, TypeAAAA)
	if err != nil {
		return nil, err
	}

	r := make([]net.IP, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.AAAA:
			r = append(r, v.AAAA)
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

// QueryAAAA ask a random server from servers and returns a slice of net.IP.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QueryAAAA(name string) ([]net.IP, error) {

	return s.Get(-1).QueryAAAA(name)
}

// QueryAAAA ask a random server from DefaultServers and returns a slice of net.IP.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QueryAAAA(name string) ([]net.IP, error) {

	return DefaultServers.QueryAAAA(name)
}

// TryQueryAAAA asks the servers for type AAAA. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQueryAAAA(name string) ([]net.IP, error) {

	rr, err := s.TryQuery(name, TypeAAAA)
	if err != nil {
		return nil, err
	}

	r := make([]net.IP, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.AAAA:
			r = append(r, v.AAAA)
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

// TryQueryAAAA asks the DefaultServers for type A. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQueryAAAA(name string) ([]net.IP, error) {

	return DefaultServers.TryQueryAAAA(name)
}

// IsSetAAAA checks whether an AAAA type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetAAAA(name string) (bool, error) {
	return s.IsSet(name, TypeAAAA)
}

// IsSetAAAA checks whether an AAAA type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetAAAA(name string) (bool, error) {
	return DefaultServers.IsSetAAAA(name)
}
