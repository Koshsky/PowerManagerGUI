package api

import (
	"slices"
	"strings"
)

const (
	GERSControl     = "GERS Control"
	MonitorAssembly = "Monitor assembly (3.0V)"
	MiniPCAssembly  = "Mini-PC assembly"
)

type guiHandler interface {
	isActionAllowed(device, state string) bool
	getDevices() []string
	getActions() []string
}

type gersControlHandler struct {
}

func (gersControlHandler) isActionAllowed(device, state string) bool {
	return slices.Contains([]string{"ON", "OFF", "Reset", "HardReset"}, state)
}

func (gersControlHandler) getDevices() []string {
	return []string{
		"ALL",
		"GERS 1",
		"GERS 2",
		"GERS 3",
		"GERS 4",
		"GERS 5",
	}
}

func (gersControlHandler) getActions() []string {
	return []string{"ON", "OFF", "Reset", "HardReset"}
}

type monitorAssemblyHandler struct {
}

func (monitorAssemblyHandler) isActionAllowed(device, state string) bool {
	if slices.Contains([]string{"ON", "OFF", "Reset"}, state) {
		return strings.HasPrefix(device, "Mini PC")
	}
	return true
}

func (monitorAssemblyHandler) getDevices() []string {
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

func (monitorAssemblyHandler) getActions() []string {
	return []string{"ON", "OFF", "Reset", "Turn ON", "Turn OFF"}
}

type miniPCAssemblyHandler struct {
}

func (miniPCAssemblyHandler) isActionAllowed(device, state string) bool {
	if slices.Contains([]string{"ON", "OFF", "Reset"}, state) {
		return strings.HasPrefix(device, "Mini PC")
	}
	return true
}

func (miniPCAssemblyHandler) getDevices() []string {
	return []string{
		"Mini PC",
		"Converter",
		"Monitor",
	}
}

func (miniPCAssemblyHandler) getActions() []string {
	return []string{"ON", "OFF", "Reset", "Turn ON", "Turn OFF"}
}
