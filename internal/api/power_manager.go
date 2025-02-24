package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

const (
	GERSControl     = "GERS Control"
	MonitorAssembly = "Monitor assembly (3.0V)"
	MiniPCAssembly  = "Mini-PC assembly"
)

type Handler interface {
	IsActionAllowedForDevice(device, state string) bool
	GetDevices() []string
	GetActions() []string
}

type PowerManager struct {
	IP       string `json:"ip"`
	Mask     string `json:"mask"`
	Gateway  string `json:"gateway"`
	Login    string `json:"login"`    // for the future
	Password string `json:"password"` // for the future
	Type     string `json:"type"`
	handler  Handler
}

func (pm *PowerManager) IsActionAllowedForDevice(device, cmd string) bool {
	return pm.handler.IsActionAllowedForDevice(device, cmd)
}

func (pm *PowerManager) GetDevices() []string {
	return pm.handler.GetDevices()
}

func (pm *PowerManager) GetActions() []string {
	return pm.handler.GetActions()
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
	deviceType, err := getDeviceTypeMock(ip)
	if err != nil {
		return nil, err
	}
	pm.Type = deviceType

	switch pm.Type {
	case GERSControl:
		pm.handler = &GERSControlManager{}
	case MonitorAssembly:
		pm.handler = &MonitorAssemblyManager{}
	case MiniPCAssembly:
		pm.handler = &MiniPCAssemblyManager{}
	default:
		return nil, fmt.Errorf("cannot create power manager: unknown type of manager: %s", pm.Type)
	}
	return pm, nil
}

func getDeviceType(ip string) (string, error) {
	url := fmt.Sprintf("http://%s/get_info.json", ip)
	response, err := http.Get(url)
	if err != nil {
		return "UNKNOWN", fmt.Errorf("failed to get device info from %s: %w", ip, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "UNKNOWN", fmt.Errorf("received non-200 response from %s: %d", ip, response.StatusCode)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return "UNKNOWN", fmt.Errorf("unexpected content type from %s: %s", ip, contentType)
	}

	var info PowerManagerInfo
	if err := json.NewDecoder(response.Body).Decode(&info); err != nil {
		return "UNKNOWN", fmt.Errorf("failed to decode JSON from %s: %w", ip, err)
	}
	return info.Type, nil
}
