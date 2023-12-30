// A simple package for logging.
//
// The Severity variable is used to determine what message is printed.
// For example if Severity is LOG_NOTICE, than LOG_NOTICE is printed, but LOG_INFO is not.
// Use GetSeverity() and SetSeverity() to use Severity thread-safe.
package log

import (
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	LOG_EMERG   int = iota // Emergency. The message says the system is unusable.
	LOG_ALERT              // Alert. Action on the message must be taken immediately.
	LOG_CRIT               // Critical. The message states a critical condition.
	LOG_ERR                // Error. The message describes an error.
	LOG_WARNING            // Warning. The message is a warning.
	LOG_NOTICE             // Notice. The message describes a normal but important event.
	LOG_INFO               // Informational. The message is purely informational.
	LOG_DEBUG              // Debug. The message is only for debugging purposes.
)

var (

	// Severity is the package wide severity level used to determine what to write.
	Severity int = LOG_NOTICE

	m *sync.RWMutex
)

func init() {

	m = new(sync.RWMutex)
}

func GetSeverity() int {

	m.RLock()
	defer m.RUnlock()

	return Severity
}

func SetSeverity(v int) {

	m.Lock()

	Severity = v

	m.Unlock()
}

// Log formats according to a format specifier and writes to w if Severity is at least s.
// It returns the number of bytes written and any write error encountered.
// Log is a low-level function.
func Log(s int, w io.Writer, format string, a ...any) (int, error) {

	if GetSeverity() < s {
		return 0, nil
	}

	return fmt.Fprintf(w, format, a...)
}

// Fatalf formats according to a format specifier and writes to STDERR if Severity is at least LOG_ERR.
// It returns the number of bytes written and any write error encountered.
func Fatalf(format string, a ...any) (int, error) {
	return Log(LOG_EMERG, os.Stderr, format, a...)

}

// Errorf formats according to a format specifier and writes to STDERR if Severity is at least LOG_ERR.
// It returns the number of bytes written and any write error encountered.
func Errorf(format string, a ...any) (int, error) {
	return Log(LOG_ERR, os.Stderr, format, a...)
}

// Printf formats according to a format specifier and writes to STDOUT if Severity is at least LOG_NOTICE.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...any) (int, error) {
	return Log(LOG_NOTICE, os.Stdout, format, a...)
}

// Printf formats according to a format specifier and writes to STDOUT if Severity is at least LOG_DEBUG.
// It returns the number of bytes written and any write error encountered.
func Debugf(format string, a ...any) (int, error) {
	return Log(LOG_DEBUG, os.Stdout, format, a...)
}
