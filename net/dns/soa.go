package dns

import (
	"fmt"

	mdns "github.com/miekg/dns"
)

var TypeSOA uint16 = 6

// See more: https://www.rfc-editor.org/rfc/rfc1035.html#section-3.3.13
type SOA struct {
	Mname   string
	Rname   string
	Serial  int
	Refresh int
	Retry   int
	Expire  int
	MinTTL  int
}

func (s SOA) String() string {
	return fmt.Sprintf("%s %s %d %d %d %d %d", s.Mname, s.Rname, s.Serial, s.Refresh, s.Retry, s.Expire, s.MinTTL)
}

// QuerySOA ask the server and returns a SOA struct pointer.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Server) QuerySOA(name string) (*SOA, error) {

	rr, err := s.query(name, TypeSOA)
	if err != nil {
		return nil, err
	}

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.SOA:
			return &SOA{Mname: v.Ns, Rname: v.Mbox, Serial: int(v.Serial), Refresh: int(v.Refresh), Retry: int(v.Retry), Expire: int(v.Expire), MinTTL: int(v.Minttl)}, nil
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

	return nil, nil
}

// QuerySOA ask a random server from servers and returns a SOA struct pointer.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func (s *Servers) QuerySOA(name string) (*SOA, error) {

	return s.Get(-1).QuerySOA(name)
}

// QuerySOA ask a random server from DefaultServers and returns a SOA struct pointer.
// The answer slice will be nil in case of error.
//
// The other record types are ignored.
func QuerySOA(name string) (*SOA, error) {

	return DefaultServers.QuerySOA(name)
}

// TryQuerySOA asks the servers for type SOA. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func (s *Servers) TryQuerySOA(name string) (*SOA, error) {

	rr, err := s.TryQuery(name, TypeSOA)
	if err != nil {
		return nil, err
	}

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.SOA:
			return &SOA{Mname: v.Ns, Rname: v.Mbox, Serial: int(v.Serial), Refresh: int(v.Refresh), Retry: int(v.Retry), Expire: int(v.Expire), MinTTL: int(v.Minttl)}, nil
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

	return nil, nil
}

// TryQuerySOA asks the DefaultServers for type SOA. If any error occurred, retries with next server (except if error is NXDOMAIN).
//
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// The first used server is random. The other record types are ignored.
func TryQuerySOA(name string) (*SOA, error) {

	return DefaultServers.TryQuerySOA(name)
}

// IsSetSOA checks whether an SOA type record set for name.
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSetSOA(name string) (bool, error) {
	return s.IsSet(name, TypeSOA)
}

// IsSetSOA checks whether an SOA type record set for name using the DefaultServers.
// NXDOMAIN is not an error here, because it means "not found".
func IsSetSOA(name string) (bool, error) {
	return DefaultServers.IsSetSOA(name)
}

// TODO: Decide if the domain is registered based on the SOA record/root server
