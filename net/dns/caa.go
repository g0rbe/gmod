package dns

import (
	"fmt"

	mdns "github.com/miekg/dns"
)

var TypeCAA uint16 = 257

type CAA struct {
	Flag  uint8
	Tag   string
	Value string
}

func (c CAA) String() string {
	return fmt.Sprintf("%d %s %s", c.Flag, c.Tag, c.Value)
}

// QueryCAA ask the server and returns a slice of CAA.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QueryCAA(name string) ([]CAA, error) {

	rr, err := s.query(name, TypeCAA)
	if err != nil {
		return nil, err
	}

	r := make([]CAA, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.CAA:
			r = append(r, CAA{Flag: v.Flag, Tag: v.Tag, Value: v.Value})
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

// QueryCAA ask a random server from servers and returns a slice of CAA.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QueryCAA(name string) ([]CAA, error) {

	return s.Get(-1).QueryCAA(name)
}

// QueryCAA ask a random server from DefaultServers and returns a slice of CAA.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QueryCAA(name string) ([]CAA, error) {

	return DefaultServers.QueryCAA(name)
}

// TryQueryCAA asks the servers for type CAA. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQueryCAA(name string) ([]CAA, error) {

	rr, err := s.TryQuery(name, TypeCAA)
	if err != nil {
		return nil, err
	}

	r := make([]CAA, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.CAA:
			r = append(r, CAA{Flag: v.Flag, Tag: v.Tag, Value: v.Value})
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

// TryQueryCAA asks the DefaultServers for type CAA. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQueryCAA(name string) ([]CAA, error) {

	return DefaultServers.TryQueryCAA(name)
}

// IsSetCAA checks whether a CAA type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetCAA(name string) (bool, error) {
	return s.IsSet(name, TypeCAA)
}

// IsSetCAA checks whether a CAA type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetCAA(name string) (bool, error) {
	return DefaultServers.IsSetCAA(name)
}
