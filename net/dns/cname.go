package dns

import (
	"fmt"

	mdns "github.com/miekg/dns"
)

var TypeCNAME uint16 = 5

// QueryCNAME ask the server and returns a slice of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QueryCNAME(name string) ([]string, error) {

	rr, err := s.query(name, TypeCNAME)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.CNAME:
			r = append(r, v.Target)
		case *mdns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// QueryCNAME ask a random server from servers and returns a slice of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QueryCNAME(name string) ([]string, error) {

	return s.Get(-1).QueryCNAME(name)
}

// QueryCNAME ask a random server from DefaultServers and returns a slice of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QueryCNAME(name string) ([]string, error) {

	return DefaultServers.QueryCNAME(name)
}

// TryQueryCNAME asks the servers for type CNAME. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQueryCNAME(name string) ([]string, error) {

	rr, err := s.TryQuery(name, TypeCNAME)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.CNAME:
			r = append(r, v.Target)
		case *mdns.DNAME:
			// Ignore DNAME
			continue
		default:
			return nil, fmt.Errorf("unknown type: %T", v)
		}
	}

	return r, nil
}

// TryQueryCNAME asks the DefaultServers for type CNAME. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQueryCNAME(name string) ([]string, error) {

	return DefaultServers.TryQueryCNAME(name)
}

// IsSetCNAME checks whether an CNAME type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetCNAME(name string) (bool, error) {
	return s.IsSet(name, TypeA)
}

// IsSetCNAME checks whether an CNAME type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetCNAME(name string) (bool, error) {
	return DefaultServers.IsSetCNAME(name)
}
