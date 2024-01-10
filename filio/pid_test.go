package filio

import "testing"

const TestPath = "./test.pid"

func TestPIDCreate(t *testing.T) {

	err := PIDCreate(TestPath, 0600)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestPIDRead(t *testing.T) {

	pid, err := PIDRead(TestPath)
	if err != nil {
		t.Fatalf("Failed to read: %s\n", err)
	}

	t.Logf("pid: %d\n", pid)
}

func TestPIDCheck(t *testing.T) {

	exist, err := PIDCheck(TestPath)
	if err != nil {
		t.Fatalf("Failed to check: %s\n", err)
	}

	if !exist {
		t.Fatalf("PID should exist\n")
	}
}

func TestPIDRemove(t *testing.T) {

	err := PIDRemove(TestPath)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
