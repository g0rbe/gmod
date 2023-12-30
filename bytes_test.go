package slitu

import "testing"

func TestIsDigit(t *testing.T) {

	v := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for i := range v {
		if !IsDigit(v[i]) {
			t.Fatalf("FAIL: %c is not a digit\n", v[i])
		}
	}
}

func TestIsHexa(t *testing.T) {

	v := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'a', 'b', 'c', 'd', 'e', 'f'}

	for i := range v {
		if !IsHexa(v[i]) {
			t.Fatalf("FAIL: %c is not a hexa\n", v[i])
		}
	}
}

func TestIsLowerLetter(t *testing.T) {

	v := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	for i := range v {
		if !IsLowerLetter(v[i]) {
			t.Fatalf("FAIL: %c is not a lowercase letter\n", v[i])
		}
	}
}

func TestIsUpperLetter(t *testing.T) {

	v := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	for i := range v {
		if !IsUpperLetter(v[i]) {
			t.Fatalf("FAIL: %c is not an uppercase letter\n", v[i])
		}
	}
}
