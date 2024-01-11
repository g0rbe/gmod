package dntfy

import "testing"

func TestSend(t *testing.T) {

	err := Send("", "", "", "summary", "body")
	if err != nil {
		t.Fatalf("Fail: %s\n", err)
	}
}
