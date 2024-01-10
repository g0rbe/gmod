package capability

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// See more: linux/capability.h
func CapToMask(x uint32) uint64 {
	return (1 << ((x) & 31))
}

// CapabilityCheck checks the given capability on pid.
// If pid is -1, checks the current process.
// The capability is read from /proc/<pid>/status.
func CapabilityCheck(cap uint32, pid int) (bool, error) {

	if pid == -1 {
		pid = os.Getpid()
	}

	path := fmt.Sprintf("/proc/%d/status", pid)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, fmt.Errorf("pid not exist: %d", pid)
	}

	out, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	lines := strings.Split(string(out), "\n")

	for i := range lines {

		if lines[i] == "" {
			continue
		}

		fields := strings.Fields(lines[i])

		if len(fields) != 2 {
			continue
		}

		if fields[0] != "CapEff:" {
			continue
		}

		eff, err := strconv.ParseUint(fields[1], 16, 64)
		if err != nil {
			return false, err
		}

		return eff&CapToMask(cap) != 0, nil
	}

	return false, fmt.Errorf("CapEff not exist in %s", path)
}
