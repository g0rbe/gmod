package validator

import (
	"net/url"
	"strings"

	"github.com/g0rbe/gmod/bitter"
	"github.com/g0rbe/gmod/net/dns"
	"github.com/g0rbe/gmod/net/ip"
	"github.com/g0rbe/gmod/net/tcp"
)

// isSpecialScheme returns whether scheme s is special.
// See: https://url.spec.whatwg.org/#special-scheme
func isSpecialScheme(s string) bool {

	switch s {
	case "ftp":
		return true
	case "file":
		return true
	case "http":
		return true
	case "https":
		return true
	case "ws":
		return true
	case "wss":
		return true
	default:
		return false
	}
}

// Check if string s contains only valid URL code points.
// See: https://url.spec.whatwg.org/#url-code-points
func checkURLCodePoints(s string) bool {

	l := len(s)

	for i := 0; i < l; i++ {

		switch {
		case bitter.IsDigit(s[i]):
			continue
		case bitter.IsLowerLetter(s[i]):
			continue
		case bitter.IsUpperLetter(s[i]):
			continue
		case s[i] == '!':
			continue
		case s[i] == '$':
			continue
		case s[i] == '&':
			continue
		case s[i] == 39:
			// ' character
			continue
		case s[i] == '(':
			continue
		case s[i] == ')':
			continue
		case s[i] == '*':
			continue
		case s[i] == '+':
			continue
		case s[i] == ',':
			continue
		case s[i] == '-':
			continue
		case s[i] == '.':
			continue
		case s[i] == '/':
			continue
		case s[i] == ':':
			continue
		case s[i] == ';':
			continue
		case s[i] == '=':
			continue
		case s[i] == '?':
			continue
		case s[i] == '@':
			continue
		case s[i] == '_':
			continue
		case s[i] == '~':
			continue
		case s[i] == '%':
			// The "%" character are allowed only when 2 hexacdecimal character is following

			// Check remaining length
			if l-i < 3 {
				return false
			}

			if !bitter.IsHexa(s[i+1]) {
				return false
			}

			if !bitter.IsHexa(s[i+2]) {
				return false
			}

			i += 2
			continue
		default:
			return false
		}
	}

	return true
}

// validOptionalPort reports whether port is either an empty string
// or matches /^:\d*$/
func validOptionalPort(port string) bool {

	/*
	 * Based on: https://cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/net/url/url.go;l=770
	 */

	if port == "" {
		return true
	}
	if port[0] != ':' {
		return false
	}
	for _, b := range port[1:] {
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}

// splitHostPort separates host and port. If the port is not valid, it returns
// the entire input as host, and it doesn't check the validity of the host.
// Unlike net.SplitHostPort, but per RFC 3986, it requires ports to be numeric.
func splitHostPort(hostPort string) (host, port string) {

	/*
	 * Based on: https://cs.opensource.google/go/go/+/refs/tags/go1.20.7:src/net/url/url.go;l=1159
	 */

	host = hostPort

	colon := strings.LastIndexByte(host, ':')
	if colon != -1 && validOptionalPort(host[colon:]) {
		host, port = host[:colon], host[colon+1:]
	}

	// If host is starts with "[" and ends with "]" must be a valid IPv6 address
	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") && ip.IsValid6(host[1:len(host)-1]) {
		host = host[1 : len(host)-1]
	}

	return
}

func checkValidHost(h string) bool {

	if h == "" {
		return true
	}

	host, port := splitHostPort(h)

	// Hostname must be set if port is set
	// Eg.: ":3000" is invalid
	if host == "" && port != "" {
		return false
	}

	// Port must be a valid port number
	if port != "" && tcp.Port(port) {
		return false
	}

	if dns.IsDomainPart(host) {
		return true
	}

	if dns.IsDomain(host) {
		return true
	}

	if ip.IsValid(host) {
		return true
	}

	return false
}

// Returns whether v is a valid, absolute URL.
func URL(v string) bool {

	r, err := url.Parse(v)
	if err != nil {
		return false
	}

	/*
	 * Check scheme
	 */

	if !checkURLCodePoints(r.Scheme) {
		return false
	}

	// URL is absolute
	if !r.IsAbs() {
		return false
	}

	/*
	 * Check host
	 */

	// URL with special scheme must contain Host
	if isSpecialScheme(r.Scheme) && r.Host == "" {
		return false
	}

	// Host is not a valid domain, ip and not empty
	if !checkValidHost(r.Host) {
		return false
	}

	// If userinfo is set, the Host must be set
	if r.User.String() != "" && r.Host == "" {
		return false
	}

	/*
	 * Check path: path can be anything
	 * See: https://url.spec.whatwg.org/#url-path-segment
	 */

	if !checkURLCodePoints(r.Path) {
		return false
	}
	/*
	 * Check query
	 * See: https://url.spec.whatwg.org/#url-query-string
	 */

	qs := r.Query()

	for key := range qs {

		if !checkURLCodePoints(key) {
			return false
		}

		if !checkURLCodePoints(qs.Get(key)) {
			return false
		}
	}

	/*
	 * Check fragment
	 * See: https://url.spec.whatwg.org/#url-fragment-string
	 */

	return checkURLCodePoints(r.Fragment)
}
