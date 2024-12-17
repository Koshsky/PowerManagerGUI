package api

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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

func isGersControlIP(ip string) bool {
	octets := strings.Split(ip, ".")
	if len(octets) != 4 || octets[0] != "10" || octets[1] != "4" {
		return false
	}
	thirdOctet, err := strconv.Atoi(octets[2])
	if err != nil || thirdOctet < 1 || thirdOctet > 255 {
		return false
	}

	return octets[3] == "5"
}

func isMonitorAssemblyIP(ip string) bool {
	octets := strings.Split(ip, ".")
	if len(octets) != 4 || octets[0] != "10" || octets[1] != "4" {
		return false
	}
	thirdOctet, err := strconv.Atoi(octets[2])
	if err != nil || thirdOctet < 1 || thirdOctet > 255 {
		return false
	}

	fourthOctet, err := strconv.Atoi(octets[3])
	return err == nil && 30 <= fourthOctet && fourthOctet <= 38
}

func NewPowerManager(ip string) (*PowerManager, error) {
	pm := &PowerManager{IP: ip}
	if isGersControlIP(ip) {
		pm.Type = "GERS control"
	} else if isMonitorAssemblyIP(ip) {
		pm.Type = "Monitor assembly (3.0V)"
	} else {
		return nil, fmt.Errorf("%s device is not GERS control or Monitor Assembly", ip)
	}
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
