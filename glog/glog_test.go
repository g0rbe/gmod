package glog

import (
	"os"
	"strings"
	"testing"
)

func ExamplePrintf() {

	SetOutDest(os.Stdout)

	SetPrintfPrefix("Printf: ")

	Printf("Hello, %s!", "World")
}

func TestPrintf(t *testing.T) {

	prefix := "prefix "
	format := "format %s"
	arg := "arg"
	result := "prefix format arg"

	b := new(strings.Builder)

	SetOutDest(b)
	SetPrintfPrefix(prefix)

	Printf(format, arg)

	if b.String() != result {
		t.Fatalf("want %s, got %s\n", result, b.String())
	}
}

func TestErrorf(t *testing.T) {

	prefix := "prefix "
	format := "format %s"
	arg := "arg"
	result := "prefix format arg"

	b := new(strings.Builder)

	SetErrorDest(b)
	SetErrorfPrefix(prefix)

	Errorf(format, arg)

	if b.String() != result {
		t.Fatalf("want %s, got %s\n", result, b.String())
	}
}

func TestDebugf(t *testing.T) {

	prefix := "prefix "
	format := "format %s"
	arg := "arg"
	result := "prefix format arg"

	b := new(strings.Builder)

	SetOutDest(b)
	SetDebugfPrefix(prefix)

	// Debug is false, so it should not print anything
	Debugf(format, arg)

	if b.Len() != 0 {
		t.Fatalf("want an empty string, got %s\n", b.String())
	}

	// Enable verbose logging
	SetDebug(true)

	Debugf(format, arg)

	if b.String() != result {
		t.Fatalf("want %s, got %s\n", result, b.String())
	}
}

// func TestFatalf(t *testing.T) {

// 	prefix := "prefix "
// 	format := "format %s"
// 	arg := "arg"
// 	result := "prefix format arg"

// 	b := new(strings.Builder)

// 	SetErrorDest(b)
// 	SetFatalfPrefix(prefix)

// 	Fatalf(-1, format, arg)

// 	if b.String() != result {
// 		t.Fatalf("want %s, got %s\n", result, b.String())
// 	}
// }
