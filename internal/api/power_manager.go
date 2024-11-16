package api

const (
	MASK    = "255.255.255.0"
	GATEWAY = "192.168.66.1"
)

type PowerManager struct {
	// TODO: узнать какие параметры НЕ нужны для апи общения
	IP       string `json:"ip"`
	Mask     string `json:"mask"`
	Gateway  string `json:"gateway"`
	Login    string `json:"login"`    // admin
	Password string `json:"password"` // usermvs
}

func NewPowerManager(ip string) *PowerManager {
	return &PowerManager{IP: ip}
}

func CreatePowerManagers(IPs []string) []*PowerManager {
	var powerManagers []*PowerManager
	for _, ip := range IPs {
		p := NewPowerManager(ip)
		powerManagers = append(powerManagers, p)
	}
	return powerManagers
}
