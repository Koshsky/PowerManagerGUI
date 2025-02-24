package api

import (
	"slices"
	"strings"
)

type MiniPCAssemblyManager struct {
}

func (mpam *MiniPCAssemblyManager) IsActionAllowedForDevice(device, state string) bool {
	if slices.Contains([]string{"ON", "OFF", "Reset"}, state) {
		return strings.HasPrefix(device, "Mini PC")
	}
	return true
}

func (mpam *MiniPCAssemblyManager) GetDevices() []string {
	return []string{
		"Mini PC",
		"Converter",
		"Monitor",
	}
}

func (mpam *MiniPCAssemblyManager) GetActions() []string {
	return []string{"ON", "OFF", "Reset", "Turn ON", "Turn OFF"}
}
