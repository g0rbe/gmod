package dns

import (
	"fmt"

	mdns "github.com/miekg/dns"
)

var TypeNS uint16 = 2

// QueryNS ask the server and returns a slice of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QueryNS(name string) ([]string, error) {

	rr, err := s.query(name, TypeNS)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.NS:
			r = append(r, v.Ns)
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

// QueryNS ask a random server from servers and returns a slice of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QueryNS(name string) ([]string, error) {

	return s.Get(-1).QueryNS(name)
}

// QueryNS ask a random server from DefaultServers and returns a slice of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QueryNS(name string) ([]string, error) {

	return DefaultServers.QueryNS(name)
}

// TryQueryNS asks the servers for type NS. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQueryNS(name string) ([]string, error) {

	rr, err := s.TryQuery(name, TypeNS)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.NS:
			r = append(r, v.Ns)
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

// TryQueryNS asks the DefaultServers for type NS. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQueryNS(name string) ([]string, error) {

	return DefaultServers.TryQueryNS(name)
}

// IsSetNS checks whether an NS type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetNS(name string) (bool, error) {
	return s.IsSet(name, TypeNS)
}

// IsSetNS checks whether an NS type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetNS(name string) (bool, error) {
	return DefaultServers.IsSetNS(name)
}
