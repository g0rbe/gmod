package freax_test

import (
	"testing"

	"github.com/g0rbe/gmod/freax"
)

var (
	PidTestPath = "./test.pid"
)

func TestPIDCreate(t *testing.T) {

	err := freax.PIDCreate(PidTestPath, 0600)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestPIDRead(t *testing.T) {

	pid, err := freax.PIDRead(PidTestPath)
	if err != nil {
		t.Fatalf("Failed to read: %s\n", err)
	}

	t.Logf("pid: %d\n", pid)
}

func TestPIDCheck(t *testing.T) {

	exist, err := freax.PIDCheck(PidTestPath)
	if err != nil {
		t.Fatalf("Failed to check: %s\n", err)
	}

	if !exist {
		t.Fatalf("PID should exist\n")
	}
}

func TestPIDRemove(t *testing.T) {

	err := freax.PIDRemove(PidTestPath)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
