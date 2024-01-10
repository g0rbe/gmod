package dns

import (
	"fmt"
	"net/url"
	"time"

	"github.com/elmasy-com/elnet/validator"
	mdns "github.com/miekg/dns"
)

type Server struct {
	Protocol string // Protocol name, must be "udp", "tcp" or "tcp-tls"
	IP       string // Ip address
	Port     string // Destination port
	family   int    // IP address family, must be "4" for IPv4 or "6" for IPv6
	client   *mdns.Client
}

// NewServer creates a new Server.
//
// The protocol must be "udp", "tcp" or "tcp-tls".
func NewServer(protocol string, ip string, port string, timeout time.Duration) (Server, error) {

	srv := Server{}

	if protocol == "" {
		return srv, fmt.Errorf("protocol is empty")
	}

	if protocol != "udp" && protocol != "tcp" && protocol != "tcp-tls" {
		return srv, fmt.Errorf("invalid protocol: %s", protocol)
	}

	srv.Protocol = protocol

	if ip == "" {
		return srv, fmt.Errorf("ip is empty")
	}

	if validator.IPv4(ip) {
		srv.IP = ip
		srv.family = 4
	} else if validator.IPv6(ip) {
		srv.IP = ip
		srv.family = 6
	} else {
		return srv, fmt.Errorf("invalid ip: %s", ip)
	}

	if port == "" {
		return srv, fmt.Errorf("port is empty")
	}

	if !validator.Port(port) {
		return srv, fmt.Errorf("invalid port: %s", port)
	}

	srv.Port = port

	srv.client = new(mdns.Client)
	srv.client.Net = srv.Protocol
	srv.client.Timeout = timeout

	return srv, nil
}

// NewServerStr creates a Server from string s.
//
// String format:
//
//	[protocol://]ip[:port]
//
// If protocol is missing, defaults to "udp". Valid protocols are: "udp", "tcp" and "tcp-tls".
//
// If port is missing, defaults to "53".
//
// Valid strings:
//   - udp://127.0.0.1:53 -> UDP query to 127.0.0.1 on port 53 (IPv4)
//   - udp://[::1]:53 -> UDP query to ::1 on port 53 (IPv6)
//   - udp://127.0.0.1 -> UDP query to 127.0.0.1 on port 53
//   - 127.0.0.1:53 -> UDP query to 127.0.0.1 on port 53
//   - 127.0.0.1 -> UDP query to 127.0.0.1 on port 53
func NewServerStr(s string, timeout time.Duration) (Server, error) {

	// The given server string is only an IPv4 address.
	if validator.IPv4(s) {
		return Server{Protocol: "udp", IP: s, Port: "53", family: 4, client: &mdns.Client{Net: "udp", Timeout: timeout}}, nil
	}

	// The given server string is only an IPv6 address.
	if validator.IPv6(s) {
		return Server{Protocol: "udp", IP: s, Port: "53", family: 6, client: &mdns.Client{Net: "udp", Timeout: timeout}}, nil
	}

	r, err := url.Parse(s)
	if err != nil {
		return Server{}, err
	}

	sr := Server{Protocol: r.Scheme, IP: r.Hostname(), Port: r.Port()}

	if sr.Protocol == "" {
		sr.Protocol = "udp"
	}

	if sr.Protocol != "udp" && sr.Protocol != "tcp" && sr.Protocol != "tcp-tls" {
		return Server{}, fmt.Errorf("invalid protocol: %s", sr.Protocol)
	}

	if validator.IPv4(sr.IP) {
		sr.family = 4
	} else if validator.IPv6(sr.IP) {
		sr.family = 6
	} else {
		return Server{}, fmt.Errorf("invalid ip: %s", sr.IP)
	}

	if sr.Port == "" {
		sr.Port = "53"
	}

	if !validator.Port(sr.Port) {
		return Server{}, fmt.Errorf("invalid port: %s", sr.Port)
	}

	sr.client = new(mdns.Client)
	sr.client.Net = sr.Protocol
	sr.client.Timeout = timeout

	return sr, nil
}

// String format and return the Server struct.
func (s Server) String() string {

	switch s.family {
	case 4:
		return fmt.Sprintf("%s://%s:%s", s.Protocol, s.IP, s.Port)
	case 6:
		return fmt.Sprintf("%s://[%s]:%s", s.Protocol, s.IP, s.Port)
	default:
		return fmt.Sprintf("%s://%s:%s", s.Protocol, s.IP, s.Port)
	}
}

// Server returns the server address string for github.com/miekg/dns.Client.Exchange.
// If the underlying IP is an IPv6 address, add brackets around the string (eg.: "[::1]:53")
func (s *Server) Server() string {

	switch s.family {
	case 4:
		return fmt.Sprintf("%s:%s", s.IP, s.Port)
	case 6:
		return fmt.Sprintf("[%s]:%s", s.IP, s.Port)
	default:
		return fmt.Sprintf("%s:%s", s.IP, s.Port)
	}
}

func (s *Server) ToTCP() *Server {

	if s.Protocol != "udp" {
		return nil
	}

	srv := new(Server)

	srv.Protocol = "tcp"
	srv.IP = s.IP
	srv.Port = s.Port
	srv.family = s.family
	srv.client = new(mdns.Client)
	srv.client.Net = "tcp"
	srv.client.Timeout = s.client.Timeout

	return srv
}

// Generic query for type t to server s.
// Returns the Answer section.
// In case of error, the answer will be nil and return ErrX or any unknown error.
// If the returned messsage is truncated, create a TCP server from s and retry the query.
func (s *Server) query(name string, t uint16) ([]mdns.RR, error) {

	msg := new(mdns.Msg)
	msg.SetQuestion(mdns.Fqdn(name), t)

	in, _, err := s.client.Exchange(msg, s.Server())
	if err != nil {
		return nil, err
	}

	if in.Truncated {

		tcpS := s.ToTCP()
		if tcpS == nil {
			return nil, ErrTruncated
		}

		return tcpS.query(name, t)
	}

	if in.Rcode == 0 {
		return in.Answer, nil
	}

	return nil, RcodeToError(in.Rcode)
}
