// A simple package for logging.
//
// The getters and setters are thread-safe functions using the package level *sync.RWMutex.
package glog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	Debug bool = false // Debug controls the write to debug logs with Debugf()

	OutDest   io.Writer = os.Stdout // Output destination to write to.
	ErrorDest io.Writer = os.Stderr // Error destination to write to.

	PrintfPrefix string // Prefix string for Printf().
	ErrorfPrefix string // Prefix string for Errorf().

	DebugfPrefix string // Prefix string for Debugf().
	FatalfPrefix string // Prefix string for Fatalf().

	m *sync.RWMutex
)

func init() {

	m = new(sync.RWMutex)
}

// GetDebug returns whether debug logging is enabled.
func GetDebug() bool {

	m.RLock()
	defer m.RUnlock()

	return Debug
}

// SetDebug sets debug logging to v.
func SetDebug(v bool) {

	m.Lock()

	Debug = v

	m.Unlock()
}

// SetOutDest sets the OutDest.
func SetOutDest(w io.Writer) {
	m.Lock()
	OutDest = w
	m.Unlock()
}

// GetOutDest returns the OutDest.
func GetOutDest() io.Writer {
	m.RLock()
	defer m.RUnlock()
	return OutDest
}

// SetErrorDest sets the ErrorDest.
func SetErrorDest(w io.Writer) {
	m.Lock()
	ErrorDest = w
	m.Unlock()
}

// GetErrorDest returns the ErrorDest.
func GetErrorDest() io.Writer {
	m.RLock()
	defer m.RUnlock()
	return ErrorDest
}

// GetPrintfPrefix returns the prefix string of Printf().
func GetPrintfPrefix() string {
	m.RLock()
	defer m.RUnlock()
	return PrintfPrefix
}

// SetPrintfPrefix sets the prefix string of Printf().
func SetPrintfPrefix(prefix string) {
	m.Lock()
	PrintfPrefix = prefix
	m.Unlock()
}

// GetErrorfPrefix returns the prefix string of Errorf().
func GetErrorfPrefix() string {
	m.RLock()
	defer m.RUnlock()
	return ErrorfPrefix
}

// SetPrintfPrefix sets the prefix string of Errorf().
func SetErrorfPrefix(prefix string) {
	m.Lock()
	ErrorfPrefix = prefix
	m.Unlock()
}

// GetDebugfPrefix returns the prefix string of Debugf().
func GetDebugfPrefix() string {
	m.RLock()
	defer m.RUnlock()
	return DebugfPrefix
}

// SetDebugfPrefix sets the prefix string of Debugf().
func SetDebugfPrefix(prefix string) {
	m.Lock()
	DebugfPrefix = prefix
	m.Unlock()
}

// GetFatalfPrefix returns the prefix string of Fatalf().
func GetFatalfPrefix() string {
	m.RLock()
	defer m.RUnlock()
	return FatalfPrefix
}

// SetFatalfPrefix sets the prefix string of Fatalf().
func SetFatalfPrefix(prefix string) {
	m.Lock()
	FatalfPrefix = prefix
	m.Unlock()
}

// Printf formats according to a format specifier and writes to OutDest.
func Printf(format string, a ...any) {
	fmt.Fprint(GetOutDest(), GetPrintfPrefix(), fmt.Sprintf(format, a...))
}

// Errorf formats according to a format specifier and writes to ErrorDest.
func Errorf(format string, a ...any) {
	fmt.Fprint(GetErrorDest(), GetErrorfPrefix(), fmt.Sprintf(format, a...))
}

// Debugf formats according to a format specifier and writes to OutDest if Verbose is true.
func Debugf(format string, a ...any) {
	if Debug {
		fmt.Fprint(GetOutDest(), GetDebugfPrefix(), fmt.Sprintf(format, a...))
	}
}

// Fatalf formats according to a format specifier, writes to ErrorDest and call os.Exit(code).
func Fatalf(code int, format string, a ...any) {
	fmt.Fprint(GetErrorDest(), GetFatalfPrefix(), fmt.Sprintf(format, a...))
	os.Exit(code)
}
