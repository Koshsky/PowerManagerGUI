package power_manager

type PowerManager struct {
	IP       string `json:"ip"`
	Mask     string `json:"mask"`
	Gateway  string `json:"gateway"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
