package validator

import "testing"

func TestCheckURLCodePoints(t *testing.T) {

	cases := []struct {
		V string // Test case
		R bool   // Expected result
	}{
		{V: "abcd", R: true},
		{V: "abcd@", R: true},
		{V: "abcd#", R: false},
		{V: "abcd%%vv", R: false},
		{V: "abcd%a", R: false},
		{V: "abcd%ab", R: true},
	}

	for i := range cases {

		if r := checkURLCodePoints(cases[i].V); r != cases[i].R {
			t.Fatalf("FAIL: \"%s\" is %v", cases[i].V, r)
		}
	}
}

func TestURL(t *testing.T) {

	cases := []struct {
		V string // Test case
		R bool   // Expected result
	}{
		{V: "http://", R: false},
		{V: "https://exa%23mple.org", R: false},
		{V: "foo://exa[mple.org", R: false},
		{V: "https://1.2.3.4.5/", R: false},
		//{V: "https://test.42", R: false},
		//{V: "https://127.0.0x0.1", R: false},
		{V: "https://255.255.4000.1", R: false},
		{V: "https://[::1", R: false},
		{V: "https://[:1]", R: false},
		{V: "https://[1:2:3:4:5:6:7:8:9]", R: false},
		{V: "https://[1::1::1]", R: false},
		{V: "https://[1:2:3!:4]", R: false},
		{V: "https://[1:2:3]", R: false},
		{V: "https://[1:1:1:1:1:1:1:127.0.0.1]", R: false},
		{V: "https://[ffff::.0.0.1]", R: false},
		{V: "https://[ffff::127.0.0.4000]", R: false},
		{V: "https://example.org/>", R: false},
		{V: " https://example.org ", R: false},
		{V: "https://example.org/%s", R: false},
		{V: "https://example.org\\path\\to\\file", R: false},
		{V: "https://user:pass@", R: false},
		{V: "https://#fragment", R: false},
		{V: "https://:443", R: false},
		{V: "https://example.org:70000", R: false},
		{V: "https://example.org:7z", R: false},
		{V: "https://example.org:8080 ", R: false},

		{V: "file://c:", R: true},
		{V: "https://example.com", R: true},
		{V: "https://example.com:443", R: true},
		{V: "https://example.com:443//", R: true},
		{V: "https://example.com:443/path", R: true},
		{V: "https://example.com:443/.", R: true},
		{V: "https://example.com:443/..", R: true},
		{V: "https://example.com:443/path?a=b", R: true},
		{V: "https://example.com:443/path?a=b#main", R: true},
		{V: "https://user@example.com:443/path?a=b#main", R: true},
		{V: "https://user:pass@example.com:443/path?a=b#main", R: true},
		{V: "https://localhost:8080", R: true},
		{V: "https://[2001:db8::8a2e:370:7334]:8080", R: true},

		// From: https://url.spec.whatwg.org/#example-url-parsing
		{V: "https:example.org", R: false},
		{V: "https://////example.com///", R: false},
		{V: "https://example.com/././foo", R: true},
		{V: "file:///C|/demo", R: false},
		{V: "file://localhost/", R: true},
		{V: "https://user:password@example.org/", R: true},
		{V: "https://example.org/foo bar", R: false},
		{V: "https://EXAMPLE.com/../x", R: true},
		{V: "https://ex ample.org/", R: false},
		{V: "example", R: false},
		{V: "https://example.com:demo", R: false},
		{V: "http://[www.example.com]/", R: false},
		{V: "https://example.org//", R: true},
		{V: "https://example.com/[]?[]#[]", R: false},
		{V: "https://example.com/a?[]#[]", R: false},
		{V: "https://example.com/a?b=c#[]", R: false},
		{V: "https://example/%?%#%", R: false},
		{V: "https://example/#%25", R: false},
		{V: "https://example/%25?%25#%25", R: false},
	}

	for i := range cases {

		if r := URL(cases[i].V); r != cases[i].R {
			t.Fatalf("FAIL: \"%s\" is %v", cases[i].V, r)
		}
	}
}
