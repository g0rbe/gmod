// DNS queries and helper function for domain names.
//
// Based on [miekg's dns module](https://github.com/miekg/dns).
package dns

import (
	"fmt"
	"time"
)

var (

	// DefaultServers is the package default servers.
	// It is set in the init().
	// Used servers: "8.8.8.8", "1.1.1.1", "9.9.9.10", "1.0.0.1", "8.8.4.4", "149.112.112.10"
	DefaultServers Servers

	// DefaultMaxRetries is the default number of retries of failed queries.
	// Must be greater than 1, else functions will fail with ErrInvalidMaxRetries.
	DefaultMaxRetries int = 5

	// DefaultQueryTimeoutSec is the default query timeout in seconds.
	DefaultQueryTimeoutSec = 2
)

func init() {

	var err error

	DefaultServers, err = NewServersStr(DefaultMaxRetries, 2*time.Second, "8.8.8.8", "1.1.1.1", "9.9.9.10", "1.0.0.1", "8.8.4.4", "149.112.112.10")
	if err != nil {
		panic(fmt.Sprintf("Failed to set DefaultServers: %s", err))
	}

}

// IsExists checks whether a record with type A, AAAA, TXT, CNAME, MX, NS, CAA or SRV is set for name.
// NXDOMAIN is not an error here, because it means "not found".
//
// If found a setted record, this function returns without trying for the other types.
func IsExists(name string) (bool, error) {

	// A
	setA, err := IsSetA(name)
	if err != nil {
		return false, fmt.Errorf("check A failed: %w", err)
	}
	if setA {
		return true, nil
	}

	// AAAA
	setAAAA, err := IsSetAAAA(name)
	if err != nil {
		return false, fmt.Errorf("check AAAA failed: %w", err)
	}
	if setAAAA {
		return true, nil
	}

	// TXT
	setTXT, err := IsSetTXT(name)
	if err != nil {
		return false, fmt.Errorf("check TXT failed: %w", err)
	}
	if setTXT {
		return true, nil
	}

	// CNAME
	setCNAME, err := IsSetCNAME(name)
	if err != nil {
		return false, fmt.Errorf("check CNAME failed: %w", err)
	}
	if setCNAME {
		return true, nil
	}

	// MX
	setMX, err := IsSetMX(name)
	if err != nil {
		return false, fmt.Errorf("check MX failed: %w", err)
	}
	if setMX {
		return true, nil
	}

	// NS
	setNS, err := IsSetNS(name)
	if err != nil {
		return false, fmt.Errorf("check NS failed: %w", err)
	}
	if setNS {
		return true, nil
	}

	// CAA
	setCAA, err := IsSetCAA(name)
	if err != nil {
		return false, fmt.Errorf("chack CAA failed: %w", err)
	}
	if setCAA {
		return true, nil
	}

	// SRV
	setSRV, err := IsSetSRV(name)
	if err != nil {
		return false, fmt.Errorf("check SRV failed: %w", err)
	}

	return setSRV, nil
}

func TypeToString(t uint16) string {

	switch t {
	case TypeA:
		return "A"
	case TypeAAAA:
		return "AAAA"
	case TypeCAA:
		return "CAA"
	case TypeCNAME:
		return "CNAME"
	case TypeDNAME:
		return "DNAME"
	case TypeMX:
		return "MX"
	case TypeNS:
		return "NS"
	case TypeSOA:
		return "SOA"
	case TypeSRV:
		return "SRV"
	case TypeTXT:
		return "TXT"
	default:
		return "unknown"
	}
}
