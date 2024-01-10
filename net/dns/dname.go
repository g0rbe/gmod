package dns

import (
	"fmt"

	mdns "github.com/miekg/dns"
)

var TypeDNAME uint16 = 39

// QueryDNAME ask the server and returns the target string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QueryDNAME(name string) (string, error) {

	rr, err := s.query(name, TypeDNAME)
	if err != nil {
		return "", err
	}

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.DNAME:
			return v.Target, nil
		case *mdns.CNAME:
			// Ignore CNAME
			continue
		default:
			return "", fmt.Errorf("unknown type: %T", v)
		}
	}

	return "", nil
}

// QueryDNAME ask a random server from servers and returns the target string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QueryDNAME(name string) (string, error) {

	return s.Get(-1).QueryDNAME(name)
}

// QueryDNAME ask a random server from DefaultServers and returns the target of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QueryDNAME(name string) (string, error) {

	return DefaultServers.QueryDNAME(name)
}

// TryQueryDNAME asks the servers for type DNAME. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQueryDNAME(name string) (string, error) {

	rr, err := s.TryQuery(name, TypeDNAME)
	if err != nil {
		return "", err
	}

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.DNAME:
			return v.Target, nil
		case *mdns.CNAME:
			// Ignore DNAME
			continue
		default:
			return "", fmt.Errorf("unknown type: %T", v)
		}
	}

	return "", nil
}

// TryQueryDNAME asks the DefaultServers for type DNAME. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQueryDNAME(name string) (string, error) {

	return DefaultServers.TryQueryDNAME(name)
}

// IsSetDNAME checks whether an DNAME type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetDNAME(name string) (bool, error) {
	return s.IsSet(name, TypeDNAME)
}

// IsSetCNAME checks whether an DNAME type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetDNAME(name string) (bool, error) {
	return DefaultServers.IsSetDNAME(name)
}
