package dns

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	mdns "github.com/miekg/dns"
)

type Servers struct {
	srvs       []Server
	maxRetries int
	m          *sync.Mutex
}

// NewServersSlice creates a new Servers from srvs.
// Retries is the maximum number retries allowed in TryQuery function.
func NewServersSlice(retries int, srvs ...Server) Servers {

	var sr Servers
	sr.srvs = make([]Server, 0, len(srvs))

	sr.srvs = append(sr.srvs, srvs...)

	sr.m = new(sync.Mutex)
	sr.maxRetries = retries

	return sr
}

// NewServersStr creates a new Servers from srvs.
// The protocol must be "udp", "tcp" or "tcp-tls". The ddefault protocol is "udp" and the default port is "53".
// Retries is the maximum number retries allowed in TryQuery function.
// Timeout is a cumulative timeout for dial.
func NewServersStr(retries int, timeout time.Duration, s ...string) (Servers, error) {

	if len(s) == 0 {
		return Servers{}, fmt.Errorf("servers is empty")
	}

	var srvs Servers
	srvs.srvs = make([]Server, 0, len(s))
	srvs.m = new(sync.Mutex)
	srvs.maxRetries = retries

	for i := range s {

		srv, err := NewServerStr(s[i], timeout)
		if err != nil {
			return srvs, fmt.Errorf("failed to create new server from %s: %w", s[i], err)
		}

		srvs.srvs = append(srvs.srvs, srv)
	}

	return srvs, nil
}

// Append add Server srv to the Servers.
func (s *Servers) Append(srv Server) {

	s.m.Lock()
	defer s.m.Unlock()

	s.srvs = append(s.srvs, srv)
}

// Get returns a DNS server to use.
// If index is between the servers range, returns the selected server.
// If index is not between the servers range, returns a random one.
//
// Set index to -1 to get a random one.
func (s *Servers) Get(index int) *Server {

	switch l := len(s.srvs); l {
	case 0:
		// No server is configured
		panic("no servers configured")
	case 1:
		// Only ne server
		return &s.srvs[0]
	default:

		if index >= 0 && index < l {
			// If index is between the servers range, return the selected server
			return &s.srvs[index]
		} else {
			// If not in range, return a random server
			return &s.srvs[rand.Intn(l)]
		}
	}
}

// GetMaxRetries returns the maximum number retries configured in the Servers.
func (s *Servers) GetMaxRetries() int {

	return s.maxRetries
}

// SetMaxRetries sets the maximum number retries in the Servers.
func (s *Servers) SetMaxRetries(n int) {

	s.m.Lock()
	s.maxRetries = n
	s.m.Unlock()
}

// TryQuery asks the servers for type t. If any error occurred, retries with an other server (except if error is NXDOMAIN).
// Returns the Answer section.
// In case of error, the answer will be nil and return ErrX or any unknown error.
//
// NOTE: The first used server is random.
func (s *Servers) TryQuery(name string, t uint16) ([]mdns.RR, error) {

	var (
		err        = ErrInvalidMaxRetries
		rr         []mdns.RR
		maxRetries = s.maxRetries - 1
	)

	for i := -1; i < maxRetries; i++ {

		rr, err = s.Get(i).query(name, t)
		if err == nil || errors.Is(err, ErrName) {
			break
		}
	}

	return rr, err
}

// IsSet checks whether a record with type t is set for name.
// This function retries the query in case of error (**NOT** errors like NXDOMAIN) up to n times (configured when created the Servers).
// NXDOMAIN is not an error here, because it means "not found".
func (s *Servers) IsSet(name string, t uint16) (bool, error) {

	rr, err := s.TryQuery(name, t)

	return len(rr) != 0, err
}
