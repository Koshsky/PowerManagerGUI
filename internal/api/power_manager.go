package api

import (
	"fmt"
	"log"
	"sort"
)

type PowerManager struct {
	IP       string `json:"ip"`
	Mask     string `json:"mask"`
	Gateway  string `json:"gateway"`
	Login    string `json:"login"`    // for the future
	Password string `json:"password"` // for the future
	Type     string `json:"type"`
	handler  guiHandler
}

func (pm *PowerManager) IsActionAllowed(device, cmd string) bool {
	return pm.handler.isActionAllowed(device, cmd)
}

func (pm *PowerManager) GetDevices() []string {
	return pm.handler.getDevices()
}

func (pm *PowerManager) GetActions() []string {
	return pm.handler.getActions()
}

func BuildPowerManagers(IPs []string) []*PowerManager {
	var powerManagers []*PowerManager
	for _, ip := range IPs {
		p, err := NewPowerManager(ip)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		powerManagers = append(powerManagers, p)
	}
	sort.Sort(ByType(powerManagers))
	return powerManagers
}

func NewPowerManager(ip string) (*PowerManager, error) {
	pm := &PowerManager{IP: ip}
	deviceType, err := getDeviceType(ip)
	if err != nil {
		return nil, err
	}
	pm.Type = deviceType

	switch pm.Type {
	case GERSControl:
		pm.handler = &gersControlHandler{}
	case MonitorAssembly:
		pm.handler = &monitorAssemblyHandler{}
	case MiniPCAssembly:
		pm.handler = &miniPCAssemblyHandler{}
	default:
		return nil, fmt.Errorf("cannot create power manager: unknown type of manager: %s", pm.Type)
	}
	return pm, nil
}
