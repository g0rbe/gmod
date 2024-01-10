package dns

import (
	"fmt"

	mdns "github.com/miekg/dns"
)

type SRV struct {
	Priority int
	Weight   int
	Port     int
	Target   string
}

func (s SRV) String() string {
	return fmt.Sprintf("%d %d %d %s", s.Priority, s.Weight, s.Port, s.Target)
}

var TypeSRV uint16 = 33

// QuerySRV ask the server and returns a slice of SRV.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QuerySRV(name string) ([]SRV, error) {

	rr, err := s.query(name, TypeSRV)
	if err != nil {
		return nil, err
	}

	r := make([]SRV, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.SRV:
			r = append(r, SRV{Priority: int(v.Priority), Weight: int(v.Weight), Port: int(v.Port), Target: v.Target})
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

// QuerySRV ask a random server from servers and returns a slice of SRV.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QuerySRV(name string) ([]SRV, error) {

	return s.Get(-1).QuerySRV(name)
}

// QuerySRV ask a random server from DefaultServers and returns a slice of SRV.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QuerySRV(name string) ([]SRV, error) {

	return DefaultServers.QuerySRV(name)
}

// TryQuerySRV asks the servers for type SRV. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQuerySRV(name string) ([]SRV, error) {

	rr, err := s.TryQuery(name, TypeSRV)
	if err != nil {
		return nil, err
	}

	r := make([]SRV, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.SRV:
			r = append(r, SRV{Priority: int(v.Priority), Weight: int(v.Weight), Port: int(v.Port), Target: v.Target})
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

// TryQuerySRV asks the DefaultServers for type SRV. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQuerySRV(name string) ([]SRV, error) {

	return DefaultServers.TryQuerySRV(name)
}

// IsSetSRV checks whether an SRV type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetSRV(name string) (bool, error) {
	return s.IsSet(name, TypeSRV)
}

// IsSetSRV checks whether an SRV type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetSRV(name string) (bool, error) {
	return DefaultServers.IsSetSRV(name)
}
