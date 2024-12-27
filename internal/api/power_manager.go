package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"sort"
	"strings"
)

const (
	GERSControl     = "GERS Control"
	MonitorAssembly = "Monitor assembly (3.0V)"
)

type PowerManager struct {
	IP       string   `json:"ip"`
	Mask     string   `json:"mask"`
	Gateway  string   `json:"gateway"`
	Login    string   `json:"login"`    // for the future
	Password string   `json:"password"` // for the future
	Type     string   `json:"type"`
	Devices  []string `json:"devices"`
	States   []string `json:"commands"`
}

func (pm *PowerManager) IsActionAllowedForDevice(device, state string) bool {
	if slices.Contains([]string{"ON", "OFF", "Reset"}, state) {
		return pm.Type == GERSControl || strings.HasPrefix(device, "Mini PC")
	}
	return true
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
	deviceType, err := getDeviceTypeMock(ip)  // TODO: replace with getDeviceType
	if err != nil {
		return nil, err
	}
	pm.Type = deviceType

	if pm.Type == GERSControl {
		pm.Devices = []string{
			"ALL",
			"GERS 1",
			"GERS 2",
			"GERS 3",
			"GERS 4",
			"GERS 5",
		}
		pm.States = []string{"ON", "OFF", "Reset", "HardReset"}
	} else if pm.Type == MonitorAssembly {
		pm.Devices = []string{
			"Mini PC 1",
			"Mini PC 2",
			"Monitor",
			"Common power",
			"Converter 1",
			"Converter 2",
			"Reserved 1",
			"Reserved 2",
		}
		pm.States = []string{"ON", "OFF", "Reset", "Turn ON", "Turn OFF"}
	} else {
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
