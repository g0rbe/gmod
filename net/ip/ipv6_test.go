package ip

import "testing"

func TestIsValid6(t *testing.T) {

	cases := []struct {
		IP     string
		Result bool
	}{
		{IP: "::", Result: true},
		{IP: "0::", Result: true},
		{IP: "::x", Result: false},
		{IP: ":", Result: false},
	}

	for i := range cases {
		if r := IsValid6(cases[i].IP); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].IP, r)
		}
	}
}
