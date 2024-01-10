package dns

import (
	"fmt"

	mdns "github.com/miekg/dns"
)

var TypeMX uint16 = 15

type MX struct {
	Preference int    // Priority
	Exchange   string // Server's hostname
}

func (m MX) String() string {
	return fmt.Sprintf("%d %s", m.Preference, m.Exchange)
}

// QueryMX ask the server and returns a slice of MX.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QueryMX(name string) ([]MX, error) {

	rr, err := s.query(name, TypeMX)
	if err != nil {
		return nil, err
	}

	r := make([]MX, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.MX:
			r = append(r, MX{Preference: int(v.Preference), Exchange: v.Mx})
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

// QueryMX ask a random server from servers and returns a slice of MX.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QueryMX(name string) ([]MX, error) {

	return s.Get(-1).QueryMX(name)
}

// QueryMX ask a random server from DefaultServers and returns a slice of MX.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QueryMX(name string) ([]MX, error) {

	return DefaultServers.QueryMX(name)
}

// TryQueryMX asks the servers for type MX. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQueryMX(name string) ([]MX, error) {

	rr, err := s.TryQuery(name, TypeMX)
	if err != nil {
		return nil, err
	}

	r := make([]MX, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.MX:
			r = append(r, MX{Preference: int(v.Preference), Exchange: v.Mx})
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

// TryQueryMX asks the DefaultServers for type MX. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQueryMX(name string) ([]MX, error) {

	return DefaultServers.TryQueryMX(name)
}

// IsSetMX checks whether an MX type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetMX(name string) (bool, error) {
	return s.IsSet(name, TypeMX)
}

// IsSetMX checks whether an MX type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetMX(name string) (bool, error) {
	return DefaultServers.IsSetMX(name)
}
