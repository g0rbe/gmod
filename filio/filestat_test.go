package filio

import (
	"testing"
)

var (
	TestFilePath = "filestat.go"
)

func TestGetFileOwner(t *testing.T) {

	owner, err := GetFileOwner(TestFilePath)
	if err != nil {
		t.Errorf("Failed to get owner: %s\n", err)
	}

	t.Logf("%#v\n", *owner)
}

func BenchmarkGetFileOwner(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := GetFileOwner(TestFilePath)
		if err != nil {
			b.Errorf("Failed to get owner: %s\n", err)
		}
	}
}

func TestGetFileGroup(t *testing.T) {

	group, err := GetFileGroup(TestFilePath)
	if err != nil {
		t.Errorf("Failed to get group: %s\n", err)
	}

	t.Logf("%#v\n", *group)
}

func BenchmarkGetFileGroup(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := GetFileGroup(TestFilePath)
		if err != nil {
			b.Errorf("Failed to get group: %s\n", err)
		}
	}
}
