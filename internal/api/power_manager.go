package api

import (
	"github.com/Koshsky/PowerManagerGUI/internal/netutils"
)

const (
	MASK    = "255.255.255.0"
	GATEWAY = "192.168.66.1"
)

type PowerManager struct {
	IP          string `json:"ip"` // 10.2.[0-255].[какой-то ограниченный диапазон]
	Mask        string `json:"mask"`
	Gateway     string `json:"gateway"`
	Login       string `json:"login"`    // admin
	Password    string `json:"password"` // usermvs
	ManagerType string `json:"manager_type"`
}

// TODO: добавить логин и пароль как параметры функции.
// mask and gateway must be constant
// TODO: нужно обсудить выводить ли логин и пароль как параметры в окне или СДЕЛАТЬ ИХ КОНСТАНТАМИ.
func NewPowerManager(ip string) *PowerManager {
	return &PowerManager{IP: ip, ManagerType: netutils.GetManagerTypeByIP(ip)}
}

// TODO: Добавить гибкую настройку логина и пароля с помощью переменных
// mask, gateway - константы
// ip, login, password - параметры функции. логин и пароль для всех айпи одинаковый.
func CreatePowerManagers(IPs []string) []*PowerManager {
	var powerManagers []*PowerManager
	for _, ip := range IPs {
		p := NewPowerManager(ip)
		powerManagers = append(powerManagers, p)
	}
	return powerManagers
}
