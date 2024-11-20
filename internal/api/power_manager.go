package api

import (
	"fmt"
	"log"
)

const (
	MASK    = "255.255.255.0"
	GATEWAY = "192.168.66.1"
)

type PowerManager struct {
	// TODO: узнать какие параметры НУЖНЫ для апи общения
	// сейчас мне кажется что достаточно только IP. даже логин и пароль избыточен...
	IP       string   `json:"ip"`
	Mask     string   `json:"mask"`
	Gateway  string   `json:"gateway"`
	Login    string   `json:"login"`    // admin
	Password string   `json:"password"` // usermvs
	Type     string   `json:"type"`     // "GERS control" / "Monitor assembly (3.0V)"
	Devices  []string `json:"devices"`
	States   []string `json:"commands"`
}

func NewPowerManager(ip string) (*PowerManager, error) {
	pm := &PowerManager{IP: ip}
	if ip == "10.4.1.5" {
		pm.Type = "GERS control"
	} else {
		pm.Type = "Monitor assembly (3.0V)"
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
		return nil, fmt.Errorf("cannot create power manager: uknown type of manager: " + pm.Type)
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
