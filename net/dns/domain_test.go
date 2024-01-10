package dns

import (
	"testing"
)

func TestIsDomain(t *testing.T) {

	// Generate a string with 400 x 'a'
	longStringFunc := func() string {
		s := make([]byte, 400)
		for i := 0; i < 400; i++ {
			s = append(s, 'a')
		}
		return string(s)
	}

	cases := []struct {
		Domain string
		Result bool
	}{
		{Domain: "elmasy.com", Result: true},
		{Domain: "elmasy.com.", Result: true},
		{Domain: ".elmasy.com", Result: false},
		{Domain: ".elmasy.com.", Result: false},
		{Domain: "aaaaaa", Result: false},
		{Domain: longStringFunc(), Result: false},
		{Domain: "", Result: false},
		{Domain: ".", Result: false},
		{Domain: "a a", Result: false},
		{Domain: "a=a", Result: false},
		{Domain: " ", Result: false},
		{Domain: "*", Result: false},
	}

	for i := range cases {
		if r := IsDomain(cases[i].Domain); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].Domain, r)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsDomain("test.elmasy.com.")
	}
}

func TestIsDomainPart(t *testing.T) {

	// Generate a string with 400 x 'a'
	longStringFunc := func() string {
		s := make([]byte, 64)
		for i := 0; i < 64; i++ {
			s = append(s, 'a')
		}
		return string(s)
	}

	cases := []struct {
		Domain string
		Result bool
	}{
		{Domain: "localhost", Result: true},
		{Domain: "elmasy.", Result: false},
		{Domain: ".elmasy", Result: false},
		{Domain: ".elmasy.com.", Result: false},
		{Domain: "aaaaaa", Result: true},
		{Domain: longStringFunc(), Result: false},
		{Domain: "", Result: false},
		{Domain: ".", Result: false},
		{Domain: "a a", Result: false},
		{Domain: "a=a", Result: false},
		{Domain: "a-a", Result: true},
		{Domain: "-aa", Result: false},
		{Domain: "aa-", Result: false},
		{Domain: "a--a", Result: false},
		{Domain: " ", Result: false},
		{Domain: "*", Result: false},
	}

	for i := range cases {
		if r := IsDomainPart(cases[i].Domain); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].Domain, r)
		}
	}
}

func BenchmarkIsDomainPart(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsDomainPart("elmasy")
	}
}

func TestGetParts(t *testing.T) {

	// 0. element = test domain
	// 1. element = tld
	// 2. element = domain
	// 3. element = sub
	cases := [][4]string{
		{"com", "com", "", ""},
		{"a.0emm.com", "com", "0emm", "a"},
		{"0emm.com", "com", "0emm", ""},
		{"amazon.co.uk", "co.uk", "amazon", ""},
		{"books.amazon.co.uk", "co.uk", "amazon", "books"},
		{"amazon.com", "com", "amazon", ""},
		{"example0.debian.net", "net", "debian", "example0"},
		{"example1.debian.org", "org", "debian", "example1"},
		{"golang.dev", "dev", "golang", ""},
		{"golang.net", "net", "golang", ""},
		{"play.golang.org", "org", "golang", "play"},
		{"gophers.in.space.museum", "space.museum", "in", "gophers"},
		{"b.c.d.0emm.com", "com", "0emm", "b.c.d"},
		{"there.is.no.such-tld", "such-tld", "no", "there.is"},
		{"foo.org", "org", "foo", ""},
		{"foo.co.uk", "co.uk", "foo", ""},
		{"foo.dyndns.org", "org", "dyndns", "foo"},
		{"www.foo.dyndns.org", "org", "dyndns", "www.foo"},
		{"foo.blogspot.co.uk", "co.uk", "blogspot", "foo"},
		{"www.foo.blogspot.co.uk", "co.uk", "blogspot", "www.foo"},
		{"test.com.test.com", "com", "test", "test.com"},
		{"s3.ca-central-1.amazonaws.com", "com", "amazonaws", "s3.ca-central-1"},
		{"www.test.r.appspot.com", "com", "appspot", "www.test.r"},
		{"test.blogspot.commmmm", "commmmm", "blogspot", "test"},
		{"test.blogspot.colu", "colu", "blogspot", "test"},
		{"test.blogspot.bgtrfesw.bgtrfesw.bgtrfesw", "bgtrfesw", "bgtrfesw", "test.blogspot.bgtrfesw"},
	}

	for i := range cases {

		parts := GetParts(cases[i][0])

		switch {

		case parts.TLD != cases[i][1]:
			t.Fatalf("FAIL: TLD failed with %s, want=%s get=%s\n", cases[i][0], cases[i][1], parts.TLD)
		case parts.Domain != cases[i][2]:
			t.Fatalf("FAIL: Domain failed with %s, want=%s get=%s\n", cases[i][0], cases[i][2], parts.Domain)
		case parts.Sub != cases[i][3]:
			t.Fatalf("FAIL: Sub failed with %s, want=%s get=%s\n", cases[i][0], cases[i][3], parts.Sub)
		}
	}
}

func BenchmarkGetParts(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetParts("test.s3.dualstack.ap-northeast-2.amazonaws.com.")
	}
}

func TestGetTLD(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{"a.", "a"},
		{".a", ""},
		{"com.", "com"},
		{".com", ""},
		{"co.uk", "co.uk"},
		{"co.uk.", "co.uk"},
		{"cromulent", "cromulent"},
		{"a.0emm.com", "com"},
		{"0emm.com", "com"},
		{"amazon.co.uk", "co.uk"},
		{"books.amazon.co.uk", "co.uk"},
		{"amazon.com", "com"},
		{"example0.debian.net", "net"},
		{"example1.debian.org", "org"},
		{"golang.dev", "dev"},
		{"golang.net", "net"},
		{"play.golang.org", "org"},
		{"gophers.in.space.museum", "space.museum"},
		{"b.c.d.0emm.com", "com"},
		{"there.is.no.such-tld", "such-tld"},
		{"foo.org", "org"},
		{"foo.co.uk", "co.uk"},
		{"foo.dyndns.org", "org"},
		{"www.foo.dyndns.org", "org"},
		{"foo.blogspot.co.uk", "co.uk"},
		{"www.foo.blogspot.co.uk", "co.uk"},
		{"test.com.test.com", "com"},
		{"test.com.", "com"},
		{"test.com.test.com.", "com"},
		{"s3.ca-central-1.amazonaws.com", "com"},
		{"www.test.r.appspot.com", "com"},
		{"test.blogspot.com", "com"},
	}

	for i := range cases {
		tld := GetTLD(cases[i][0])
		if tld != cases[i][1] {
			t.Fatalf("Case: '%s', want: '%s', got: '%s'\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetTLD(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetTLD("test.s3.dualstack.ap-northeast-2.amazonaws.com")
	}
}

func TestGetTLDIndex(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{"a.", "a."},
		{".a", ""},
		{"com.", "com."},
		{".com", ""},
		{"co.uk", "co.uk"},
		{"co.uk.", "co.uk."},
		{"cromulent", "cromulent"},
		{"a.0emm.com", "com"},
		{"0emm.com", "com"},
		{"amazon.co.uk", "co.uk"},
		{"books.amazon.co.uk", "co.uk"},
		{"amazon.com", "com"},
		{"example0.debian.net", "net"},
		{"example1.debian.org", "org"},
		{"golang.dev", "dev"},
		{"golang.net", "net"},
		{"play.golang.org", "org"},
		{"gophers.in.space.museum", "space.museum"},
		{"b.c.d.0emm.com", "com"},
		{"there.is.no.such-tld", "such-tld"},
		{"foo.org", "org"},
		{"foo.co.uk", "co.uk"},
		{"foo.dyndns.org", "org"},
		{"www.foo.dyndns.org", "org"},
		{"foo.blogspot.co.uk", "co.uk"},
		{"www.foo.blogspot.co.uk", "co.uk"},
		{"test.com.test.com", "com"},
		{"test.com.", "com."},
		{"test.com.test.com.", "com."},
		{"s3.ca-central-1.amazonaws.com", "com"},
		{"www.test.r.appspot.com", "com"},
		{"test.blogspot.com", "com"}}

	for i := range cases {
		tld := GetTLDIndex(cases[i][0])

		if tld == -1 {
			if cases[i][1] != "" {
				t.Fatalf("Case: '%s', want: '%s', index: %d\n", cases[i][0], cases[i][1], tld)
			}
			continue
		}

		if cases[i][0][tld:] != cases[i][1] {
			t.Fatalf("Case: '%s', want: '%s', got: '%s', index: %d\n", cases[i][0], cases[i][1], cases[i][0][tld:], tld)
		}
	}
}

func BenchmarkGetTLDIndex(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetTLDIndex("test.s3.dualstack.ap-northeast-2.amazonaws.com")
	}
}

func TestGetDomain(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{".cromulent", ""},
		{"a.0emm.com", "0emm.com"}, // a.0emm.com is a TLD as per publicsuffix
		{"0emm.com", "0emm.com"},   // 0emm.com is not a TLD, only *.0emm.com
		{"amazon.co.uk", "amazon.co.uk"},
		{"books.amazon.co.uk", "amazon.co.uk"},
		{"amazon.com", "amazon.com"},
		{"example0.debian.net", "debian.net"},
		{"example1.debian.org", "debian.org"},
		{"golang.dev", "golang.dev"},
		{"golang.net", "golang.net"},
		{"play.golang.org", "golang.org"},
		{"gophers.in.space.museum", "in.space.museum"},
		{"b.c.d.0emm.com", "0emm.com"},
		{"there.is.no.such-tld", "no.such-tld"},
		{"foo.org", "foo.org"},
		{"foo.co.uk", "foo.co.uk"},
		{"foo.dyndns.org", "dyndns.org"},
		{"www.foo.dyndns.org", "dyndns.org"},
		{"foo.blogspot.co.uk", "blogspot.co.uk"},
		{"www.foo.blogspot.co.uk", "blogspot.co.uk"},
		{"test.com.test.com", "test.com"},
		{"test.com.", "test.com"},
		{"test.com.test.com.", "test.com"},
		{"s3.ca-central-1.amazonaws.com", "amazonaws.com"},
		{"www.test.r.appspot.com", "appspot.com"},
		{"test.blogspot.com", "blogspot.com"},
	}

	for i := range cases {
		tld := GetDomain(cases[i][0])
		if tld != cases[i][1] {
			t.Fatalf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetDomain(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetDomain("test.s3.dualstack.ap-northeast-2.amazonaws.com.")
	}
}

func TestGetDomainIndex(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{".cromulent", ""},
		{"a.0emm.com", "0emm.com"}, // a.0emm.com is a TLD as per publicsuffix
		{"0emm.com", "0emm.com"},   // 0emm.com is not a TLD, only *.0emm.com
		{"amazon.co.uk", "amazon.co.uk"},
		{"books.amazon.co.uk", "amazon.co.uk"},
		{"amazon.com", "amazon.com"},
		{"example0.debian.net", "debian.net"},
		{"example1.debian.org", "debian.org"},
		{"golang.dev", "golang.dev"},
		{"golang.net", "golang.net"},
		{"play.golang.org", "golang.org"},
		{"gophers.in.space.museum", "in.space.museum"},
		{"b.c.d.0emm.com", "0emm.com"},
		{"there.is.no.such-tld", "no.such-tld"},
		{"foo.org", "foo.org"},
		{"foo.co.uk", "foo.co.uk"},
		{"foo.dyndns.org", "dyndns.org"},
		{"www.foo.dyndns.org", "dyndns.org"},
		{"foo.blogspot.co.uk", "blogspot.co.uk"},
		{"www.foo.blogspot.co.uk", "blogspot.co.uk"},
		{"test.com.test.com", "test.com"},
		{"test.com.", "test.com."},
		{"test.com.test.com.", "test.com."},
		{"s3.ca-central-1.amazonaws.com", "amazonaws.com"},
		{"www.test.r.appspot.com", "appspot.com"},
		{"test.blogspot.com", "blogspot.com"},
	}

	for i := range cases {
		tld := GetDomainIndex(cases[i][0])

		if tld == -1 {
			if cases[i][1] != "" {
				t.Fatalf("Case: '%s', want: '%s', index: %d\n", cases[i][0], cases[i][1], tld)
			}
			continue
		}

		if cases[i][0][tld:] != cases[i][1] {
			t.Fatalf("Case: '%s', want: '%s', got: '%s', index: %d\n", cases[i][0], cases[i][1], cases[i][0][tld:], tld)
		}
	}
}

func BenchmarkGetDomainIndex(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetDomainIndex("test.s3.dualstack.ap-northeast-2.amazonaws.com.")
	}
}

func TestGetSub(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := [][2]string{
		{"", ""},
		{".", ""},
		{".cromulent", ""},
		{"a.0emm.com", "a"}, // a.0emm.com is a TLD as per publicsuffix
		{"0emm.com", ""},    // 0emm.com is not a TLD, only *.0emm.com
		{"amazon.co.uk", ""},
		{"books.amazon.co.uk", "books"},
		{"amazon.com", ""},
		{"example0.debian.net", "example0"},
		{"example1.debian.org", "example1"},
		{"golang.dev", ""},
		{"golang.net", ""},
		{"play.golang.org", "play"},
		{"gophers.in.space.museum", "gophers"},
		{"b.c.d.0emm.com", "b.c.d"},
		{"there.is.no.such-tld", "there.is"},
		{"foo.org", ""},
		{"foo.co.uk", ""},
		{"foo.dyndns.org", "foo"},
		{"www.foo.dyndns.org", "www.foo"},
		{"foo.blogspot.co.uk", "foo"},
		{"www.foo.blogspot.co.uk", "www.foo"},
		{"test.com.test.com", "test.com"},
		{"test.com.", ""},
		{"test.com.test.com.", "test.com"},
		{"s3.ca-central-1.amazonaws.com", "s3.ca-central-1"},
		{"www.test.r.appspot.com", "www.test.r"},
		{"test.blogspot.com", "test"},
	}

	for i := range cases {
		tld := GetSub(cases[i][0])
		if tld != cases[i][1] {
			t.Fatalf("Case: %s, want: %s, got: %s\n", cases[i][0], cases[i][1], tld)
		}
	}
}

func BenchmarkGetSub(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetSub("test.s3.dualstack.ap-northeast-2.amazonaws.com.")
	}
}

func TestHasSub(t *testing.T) {

	// 1. element = test string
	// 2. element = wanted result
	cases := []struct {
		Value  string
		Result bool
	}{
		{"", false},
		{".", false},
		{".cromulent", false},
		{"a.0emm.com", true}, // a.0emm.com is a TLD as per publicsuffix
		{"0emm.com", false},  // 0emm.com is not a TLD, only *.0emm.com
		{"amazon.co.uk", false},
		{"books.amazon.co.uk", true},
		{"amazon.com", false},
		{"example0.debian.net", true},
		{"example1.debian.org", true},
		{"golang.dev", false},
		{"golang.net", false},
		{"play.golang.org", true},
		{"gophers.in.space.museum", true},
		{"b.c.d.0emm.com", true},
		{"there.is.no.such-tld", true},
		{"foo.org", false},
		{"foo.co.uk", false},
		{"foo.dyndns.org", true},
		{"www.foo.dyndns.org", true},
		{"foo.blogspot.co.uk", true},
		{"www.foo.blogspot.co.uk", true},
		{"test.com.test.com", true},
		{"test.com.", false},
		{"test.com.test.com.", true},
		{"s3.ca-central-1.amazonaws.com", true},
		{"www.test.r.appspot.com", true},
		{"test.blogspot.com", true},
	}

	for i := range cases {
		r := HasSub(cases[i].Value)
		if r != cases[i].Result {
			t.Fatalf("Case: %s, want: %v, got: %v\n", cases[i].Value, cases[i].Result, r)
		}
	}
}

func BenchmarkHasSub(b *testing.B) {

	for i := 0; i < b.N; i++ {
		HasSub("test.s3.dualstack.ap-northeast-2.amazonaws.com.")
	}
}
