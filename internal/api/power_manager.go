package api

type PowerManager struct {
	IP       string `json:"ip"`       // 10.2.[0-255].[какой-то ограниченный диапазон]
	Mask     string `json:"mask"`     // 255.255.255.0
	Gateway  string `json:"gateway"`  // 192.168.66.1
	Login    string `json:"login"`    // admin
	Password string `json:"password"` //usermvs
}

func NewPowerManager(IP string) *PowerManager {
	return &PowerManager{IP: IP}
}

func CreatePowerManagers(IPs []string) []*PowerManager {
	var powerManagers []*PowerManager
	for _, ip := range IPs {
		pm := NewPowerManager(ip)
		powerManagers = append(powerManagers, pm)
	}
	return powerManagers
}
