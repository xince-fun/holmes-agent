//go:build linux

package machine

import "testing"

func TestMachineId(t *testing.T) {
	if !hasNamespaceAccess() {
		t.Skip("Skipping MachineId test due to lack of permissions or non-Linux system")
	}

	id := MachineId()

	if id == "" {
		t.Error("Expected non-empty machine-id")
	}
	t.Logf("Machine-Id: %s", id)
}
