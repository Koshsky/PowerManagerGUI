package api

import (
	"slices"
	"strings"
)

type MonitorAssemblyManager struct {
}

func (mam *MonitorAssemblyManager) IsActionAllowedForDevice(device, state string) bool {
	if slices.Contains([]string{"ON", "OFF", "Reset"}, state) {
		return strings.HasPrefix(device, "Mini PC")
	}
	return true
}

func (mam *MonitorAssemblyManager) GetDevices() []string {
	return []string{
		"Mini PC 1",
		"Mini PC 2",
		"Monitor",
		"Common power",
		"Converter 1",
		"Converter 2",
		"Reserved 1",
		"Reserved 2",
	}
}

func (mam *MonitorAssemblyManager) GetActions() []string {
	return []string{"ON", "OFF", "Reset", "Turn ON", "Turn OFF"}
}
