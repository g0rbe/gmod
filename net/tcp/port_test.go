package tcp

import "testing"

func TestPort(t *testing.T) {

	{
		cases := []struct {
			Port   string
			Result bool
		}{
			{Port: "-1", Result: false},
			{Port: "0", Result: true},
			{Port: "80", Result: true},
			{Port: "65535", Result: true},
			{Port: "65536", Result: false},
		}

		for i := range cases {
			if r := Port(cases[i].Port); r != cases[i].Result {
				t.Fatalf("FAIL: %s is %v", cases[i].Port, r)
			}
		}
	}

	{
		cases := []struct {
			Port   int
			Result bool
		}{
			{Port: -1, Result: false},
			{Port: 0, Result: true},
			{Port: 80, Result: true},
			{Port: 65535, Result: true},
			{Port: 65536, Result: false},
		}

		for i := range cases {
			if r := Port(cases[i].Port); r != cases[i].Result {
				t.Fatalf("FAIL: %d is %v", cases[i].Port, r)
			}
		}
	}

}
