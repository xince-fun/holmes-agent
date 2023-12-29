//go:build linux

package machine

import (
	"os"
	"path"
	"strings"

	"github.com/xince-fun/holmes-agent/pkg/logger"
)

func MachineId() string {
	for _, p := range []string{"sys/devices/virtual/dmi/id/product_uuid", "etc/machine-id", "var/lib/dbus/machine-id"} {
		payload, err := os.ReadFile(path.Join("/proc/1/root", p))
		if err != nil {
			logger.DefaultLogger.Warnln("failed to read machine-id:", err)
			continue
		}
		id := strings.TrimSpace(strings.Replace(string(payload), "-", "", -1))
		logger.DefaultLogger.Infof("read machine-id: %s", id)
		return id
	}
	return ""
}
