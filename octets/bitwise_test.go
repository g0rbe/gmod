package octets_test

import (
	"testing"

	"github.com/g0rbe/gmod/octets"
)

func TestBitwiseSet(t *testing.T) {

	var v byte = 0b00000000

	v = octets.BitwiseSet(v, 0b00000001)

	v = octets.BitwiseSet(v, 0b1000000)
	v = octets.BitwiseSet(v, 0b1000000)

	if v != 0b01000001 {
		t.Fatalf("Want 0b01000001, got 0b%08b\n", v)
	}

	t.Logf("0b%08b\n", v)
}

func TestBitwiseClear(t *testing.T) {

	var v byte = 0b01000001

	v = octets.BitwiseClear(v, 0b00000001)

	v = octets.BitwiseClear(v, 0b1000000)
	v = octets.BitwiseClear(v, 0b1000000)

	if v != 0b00000000 {
		t.Fatalf("Want 0b00000000, got 0b%08b\n", v)
	}

	t.Logf("0b%08b\n", v)
}

func TestBitwiseToggle(t *testing.T) {

	var v byte = 0b00000000

	v = octets.BitwiseToggle(v, 0b00000001)

	v = octets.BitwiseToggle(v, 0b00010000)
	v = octets.BitwiseToggle(v, 0b00010000)

	v = octets.BitwiseToggle(v, 0b1000000)

	if v != 0b01000001 {
		t.Fatalf("Want 0b01000001, got 0b%08b\n", v)
	}

	t.Logf("0b%08b\n", v)
}

func TestBitwiseCheck(t *testing.T) {

	var v byte = 0b01000001

	if !octets.BitwiseCheck(v, 0b01000001) {
		t.Fatalf("Bits 0b01000001 should be set, got: 0b%08b\n", v)
	}

	if octets.BitwiseCheck(v, 0b10000010) {
		t.Fatalf("Bits 0b10000010 should not be set, got: 0b%08b\n", v)
	}
}
