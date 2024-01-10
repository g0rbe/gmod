package capability

import (
	"os"
	"testing"
)

func TestCheck(t *testing.T) {

	ok, err := CapabilityCheck(CAP_NET_ADMIN, -1)
	if err != nil {
		t.Errorf("Failed to check: %s\n", err)
	}

	switch {
	case os.Geteuid() != 0 && ok:
		t.Errorf("CAP_CHOWN must NOT be set")
	case os.Geteuid() == 0 && !ok:
		t.Errorf("CAP_CHOWN must be set")
	}
}
