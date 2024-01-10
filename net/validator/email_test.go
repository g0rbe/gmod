package validator

import "testing"

func TestEmail(t *testing.T) {

	cases := []struct {
		URL    string
		Result bool
	}{

		{URL: "simple@example.com", Result: true},
		{URL: "very.common@example.com", Result: true},
		{URL: "disposable.style.email.with+symbol@example.com", Result: true},
		{URL: "other.email-with-hyphen@example.com", Result: true},
		{URL: "fully-qualified-domain@example.com", Result: true},
		{URL: "user.name+tag+sorting@example.com", Result: true},
		{URL: "x@example.com", Result: true},
		{URL: "example-indeed@strange-example.com", Result: true},
		{URL: "test/test@test.com", Result: true},

		{URL: "test", Result: false},
		{URL: "test.com", Result: false},
		{URL: "@test.com", Result: false},
		{URL: ".a@test.com", Result: false},
		{URL: "a.@test.com", Result: false},
		{URL: "a..a@test.com", Result: false},
		{URL: "a@test", Result: false},
		{URL: "A@b@c@example.com", Result: false},
		{URL: "a\"b(c)d,e:f;g<h>i[j\\k]l@example.com", Result: false},
		{URL: "just\"not\"right@example.com", Result: false},
		{URL: "this is\"not\\allowed@example.com", Result: false},
		{URL: "1234567890123456789012345678901234567890123456789012345678901234+x@example.com", Result: false},
	}

	for i := range cases {
		if r := Email(cases[i].URL); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].URL, r)
		}
	}
}
