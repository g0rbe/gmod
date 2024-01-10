package ip

import "testing"

func TestIsValid4(t *testing.T) {

	cases := []struct {
		IP     string
		Result bool
	}{
		{IP: "0.0.0.0", Result: true},
		{IP: "255.255.255.255", Result: true},
		{IP: "0.0.0", Result: false},
		{IP: "0.0.0.256", Result: false},
	}

	for i := range cases {
		if r := IsValid4(cases[i].IP); r != cases[i].Result {
			t.Fatalf("FAIL: %s is %v", cases[i].IP, r)
		}
	}
}
