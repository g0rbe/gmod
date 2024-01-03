package gpid

import "testing"

const TestPath = "./test.pid"

func TestCreate(t *testing.T) {

	err := CreatePath(TestPath, 0600)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestRead(t *testing.T) {

	pid, err := ReadPath(TestPath)
	if err != nil {
		t.Fatalf("Failed to read: %s\n", err)
	}

	t.Logf("pid: %d\n", pid)
}

func TestCheck(t *testing.T) {

	exist, err := CheckPath(TestPath)
	if err != nil {
		t.Fatalf("Failed to check: %s\n", err)
	}

	if !exist {
		t.Fatalf("PID should exist\n")
	}
}

func TestRemove(t *testing.T) {

	err := RemovePath(TestPath)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
