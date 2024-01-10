package dns

import (
	"fmt"

	"github.com/elmasy-com/slices"
	mdns "github.com/miekg/dns"
)

type Record struct {
	Type  uint16
	Value string
}

// QueryAll query every known type and returns the records.
// This function checks whether name with the type is a wildcard, and if name is a wildcard, ommit from the retuned []Record.
//
// It is possible to return records when error returned.
func (s *Servers) QueryAll(name string) ([]Record, error) {

	var (
		rr = make([]mdns.RR, 0)
	)

	/*
	 * A
	 */

	r, err := s.TryQuery(name, TypeA)
	if err != nil {

		// NXDOMAIN means, that there is no record for name
		// If server responds NOERROR with 0 answer, means that there is a record for name, but not with the given type
		return nil, err

	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeA)
		if err != nil {

			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * AAAA
	 */

	r, err = s.TryQuery(name, TypeAAAA)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeAAAA)
		if err != nil {
			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * CAA
	 */

	r, err = s.TryQuery(name, TypeCAA)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeCAA)
		if err != nil {
			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * CNAME
	 */

	r, err = s.TryQuery(name, TypeCNAME)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeCNAME)
		if err != nil {
			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * DNAME
	 */

	r, err = s.TryQuery(name, TypeDNAME)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeDNAME)
		if err != nil {
			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * MX
	 */
	r, err = s.TryQuery(name, TypeMX)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeMX)
		if err != nil {
			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * NS
	 */

	r, err = s.TryQuery(name, TypeNS)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeNS)
		if err != nil {
			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * SOA
	 */

	r, err = s.TryQuery(name, TypeSOA)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeSOA)
		if err != nil {
			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * SRV
	 */

	r, err = s.TryQuery(name, TypeSRV)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeSRV)
		if err != nil {
			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * TXT
	 */

	r, err = s.TryQuery(name, TypeTXT)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {

		// Checks whether name is a wildcard
		wc, err := s.IsWildcard(name, TypeTXT)
		if err != nil {
			// Ignore error and assume that name is a wildcard
			wc = true
		}

		// If not a wildcard domain, append the result
		if !wc {
			rr = append(rr, r...)
		}
	}

	/*
	 * Read the answers
	 */
	rs := make([]Record, 0, len(rr))

	for i := range rr {

		switch v := rr[i].(type) {
		case *mdns.A:
			rs = slices.AppendUnique(rs, Record{Type: TypeA, Value: v.A.String()})
		case *mdns.AAAA:
			rs = slices.AppendUnique(rs, Record{Type: TypeAAAA, Value: v.AAAA.String()})
		case *mdns.CAA:
			rs = slices.AppendUnique(rs, Record{Type: TypeCAA, Value: fmt.Sprintf("%d %s %s", v.Flag, v.Tag, v.Value)})
		case *mdns.CNAME:
			rs = slices.AppendUnique(rs, Record{Type: TypeCNAME, Value: v.Target})
		case *mdns.DNAME:
			rs = slices.AppendUnique(rs, Record{Type: TypeDNAME, Value: v.Target})
		case *mdns.MX:
			rs = slices.AppendUnique(rs, Record{Type: TypeMX, Value: fmt.Sprintf("%d %s", v.Preference, v.Mx)})
		case *mdns.NS:
			rs = slices.AppendUnique(rs, Record{Type: TypeNS, Value: v.Ns})
		case *mdns.SOA:
			rs = slices.AppendUnique(rs, Record{Type: TypeSOA, Value: fmt.Sprintf("%s %s %d %d %d %d %d", v.Ns, v.Mbox, v.Serial, v.Refresh, v.Retry, v.Expire, v.Minttl)})
		case *mdns.SRV:
			rs = slices.AppendUnique(rs, Record{Type: TypeSRV, Value: fmt.Sprintf("%d %d %d %s", v.Priority, v.Weight, v.Port, v.Target)})
		case *mdns.TXT:
			for ii := range v.Txt {
				rs = slices.AppendUnique(rs, Record{Type: TypeTXT, Value: v.Txt[ii]})
			}
		default:
			return rs, fmt.Errorf("unknown type: %T", v)
		}
	}

	return rs, nil
}

// QueryAll query every known type and returns the records.
// This function checks whether name with the type is a wildcard.
func QueryAll(name string) ([]Record, error) {

	return DefaultServers.QueryAll(name)
}
