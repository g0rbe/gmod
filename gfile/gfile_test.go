package gfile

import (
	"math/rand"
	"os"
	"testing"
)

func TestIsLineInFile(t *testing.T) {

	file, err := os.OpenFile("testfile", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("Failed to create testfile: %s\n", err)
	}
	defer os.Remove("testfile")
	defer file.Close()

	if _, err := file.WriteString("a\nb\nc\nd\n"); err != nil {
		t.Fatalf("Failed to write to testfile: %s\n", err)
	}

	exist, err := IsLineInFile("testfile", "d")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	if !exist {
		t.Fatalf("\"d\" should exist in testfile\n")
	}
}

func TestCountFileLines(t *testing.T) {

	randLen := rand.Intn(10240)

	file, err := os.OpenFile("testfile", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		t.Fatalf("Failed to create test file: %s\n", err)
	}
	defer os.Remove("testfile")
	defer file.Close()

	for i := 0; i < randLen; i++ {
		file.Write([]byte{'a', '\n'})
	}

	n, err := CountFileLines("testfile")
	if err != nil {
		t.Fatalf("Failed to count: %s\n", err)
	}

	if n != randLen {
		t.Fatalf("Invalid result: want: %d, got: %d\n", randLen, n)
	}
}

func BenchmarkIsLineinFile(b *testing.B) {

	file, err := os.OpenFile("testfile", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		b.Fatalf("Failed to create testfile: %s\n", err)
	}
	defer os.Remove("testfile")
	defer file.Close()

	if _, err := file.WriteString("a\nb\nc\nd\n"); err != nil {
		b.Fatalf("Failed to write to testfile: %s\n", err)
	}

	file.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsLineInFile("testfile", "d")
	}

}

func TestIsExists(t *testing.T) {

	testName := "testfile.txt"

	file, err := os.OpenFile(testName, os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		t.Fatalf("Failed to create %s: %s\n", testName, err)
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("Failed to close %s: %s\n", testName, err)
	}

	ok, err := IsExists(testName)
	if err != nil {
		t.Fatalf("FAIL: failed to check if exist: %s\n", err)
	}

	if !ok {
		t.Fatalf("FAIL: %s should exists!\n", testName)
	}

	err = os.Remove(testName)
	if err != nil {
		t.Fatalf("Failed to remove %s: %s\n", testName, err)
	}
}

func BenchmarkIsExistsTrue(b *testing.B) {

	testName := "testfile.txt"

	file, err := os.OpenFile(testName, os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		b.Fatalf("Failed to create %s: %s\n", testName, err)
	}

	err = file.Close()
	if err != nil {
		b.Fatalf("Failed to close %s: %s\n", testName, err)
	}

	clean := func() {
		err = os.Remove(testName)
		if err != nil {
			b.Fatalf("Failed to remove %s: %s\n", testName, err)
		}
	}
	b.Cleanup(clean)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsExists(testName)
	}

}

func BenchmarkIsExistsFalse(b *testing.B) {

	testName := "testfile.txt"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsExists(testName)
	}

}

func TestIsDir(t *testing.T) {

	testName := "testdir"

	err := os.Mkdir(testName, 0600)
	if err != nil {
		t.Fatalf("Failed to create %s: %s\n", testName, err)
	}

	ok, err := IsDir(testName)
	if err != nil {
		t.Fatalf("FAIL: failed to check directory: %s\n", err)
	}

	if !ok {
		t.Fatalf("FAIL: %s should be a directory!\n", testName)
	}

	err = os.Remove(testName)
	if err != nil {
		t.Fatalf("Failed to remove %s: %s\n", testName, err)
	}
}

func BenchmarkIsDirTrue(b *testing.B) {

	testName := "testdir"

	err := os.Mkdir(testName, 0600)
	if err != nil {
		b.Fatalf("Failed to create %s: %s\n", testName, err)
	}

	clean := func() {
		err = os.Remove(testName)
		if err != nil {
			b.Fatalf("Failed to remove %s: %s\n", testName, err)
		}
	}
	b.Cleanup(clean)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsDir(testName)
	}

}

func BenchmarkIsDirFalse(b *testing.B) {

	testName := "testdir"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsDir(testName)
	}

}
