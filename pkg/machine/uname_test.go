//go:build linux

package machine

import (
	"os"
	"testing"
)

func TestUname(t *testing.T) {
	if !hasNamespaceAccess() {
		t.Skip("Failed Uname test due to lack of permissions or non-Linux system")
	}

	hostname, kernelVersion, err := Uname()
	if err != nil {
		t.Errorf("Uname() error = %v", err)
	}
	if hostname == "" {
		t.Error("Expected non-empty hostname")
	}
	if kernelVersion == "" {
		t.Error("Expected non-empty kernel version")
	}
	t.Logf("Hostname: %s, KernelVersion: %s", hostname, kernelVersion)
}

func hasNamespaceAccess() bool {
	return os.Getuid() == 0
}
