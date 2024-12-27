package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	MASK    = "255.255.255.0"
	GATEWAY = "192.168.66.1"
)

type PowerManager struct {
	IP       string   `json:"ip"`
	Mask     string   `json:"mask"`
	Gateway  string   `json:"gateway"`
	Login    string   `json:"login"`    // for the future
	Password string   `json:"password"` // for the future
	Type     string   `json:"type"`     // "GERS control" / "Monitor assembly (3.0V)"
	Devices  []string `json:"devices"`
	States   []string `json:"commands"`
}

func getDeviceType(ip string) (string, error) {
	url := fmt.Sprintf("http://%s/get_info.json", ip)
	response, err := http.Get(url)
	if err != nil {
		return "UNKNOWN", fmt.Errorf("unknown device type: %s", ip)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "UNKNOWN", fmt.Errorf("unknown device type: %s", ip)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return "UNKNOWN", fmt.Errorf("unknown device type: %s", ip)
	}

	var info PowerManagerInfo
	if err := json.NewDecoder(response.Body).Decode(&info); err != nil {
		return "UNKNOWN", fmt.Errorf("unknown device type: %s", ip)
	}
	return info.Type, nil
}

func NewPowerManager(ip string) (*PowerManager, error) {
	pm := &PowerManager{IP: ip}
	deviceType, err := getDeviceType(ip)
	if err != nil {
		return nil, fmt.Errorf("NewPowerManager error: %v", err)
	}
	pm.Type = deviceType

	if pm.Type == "GERS control" {
		pm.Devices = []string{
			"ALL",
			"GERS 1",
			"GERS 2",
			"GERS 3",
			"GERS 4",
			"GERS 5",
		}
		pm.States = []string{"ON", "OFF", "Reset", "HardReset"}
	} else if pm.Type == "Monitor assembly (3.0V)" {
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
		return nil, fmt.Errorf("cannot create power manager: uknown type of manager: %s", pm.Type)
	}
	return pm, nil
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
	return powerManagers
}
