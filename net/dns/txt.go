package dns

import (
	"fmt"

	mdns "github.com/miekg/dns"
)

var TypeTXT uint16 = 16

// QueryTXT ask the server and returns a slice of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QueryTXT(name string) ([]string, error) {

	rr, err := s.query(name, TypeTXT)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.TXT:
			r = append(r, v.Txt...)
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

// QueryTXT ask a random server from servers and returns a slice of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QueryTXT(name string) ([]string, error) {

	return s.Get(-1).QueryTXT(name)
}

// QueryTXT ask a random server from DefaultServers and returns a slice of string.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QueryTXT(name string) ([]string, error) {

	return DefaultServers.QueryTXT(name)
}

// TryQueryTXT asks the servers for type TXT. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQueryTXT(name string) ([]string, error) {

	rr, err := s.TryQuery(name, TypeTXT)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.TXT:
			r = append(r, v.Txt...)
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

// TryQueryTXT asks the DefaultServers for type TXT. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQueryTXT(name string) ([]string, error) {

	return DefaultServers.TryQueryTXT(name)
}

// IsSetTXT checks whether an TXT type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetTXT(name string) (bool, error) {
	return s.IsSet(name, TypeTXT)
}

// IsSetTXT checks whether an TXT type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetTXT(name string) (bool, error) {
	return DefaultServers.IsSetTXT(name)
}
