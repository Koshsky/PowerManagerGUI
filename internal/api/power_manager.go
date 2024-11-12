package api

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type PowerManager struct {
	IP       string `json:"ip"`       // 10.2.[0-255].[какой-то ограниченный диапазон]
	Mask     string `json:"mask"`     // 255.255.255.0
	Gateway  string `json:"gateway"`  // 192.168.66.1
	Login    string `json:"login"`    // admin
	Password string `json:"password"` // usermvs
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

// Ping отправляет ICMP Echo запрос на адрес Power Manager и возвращает результат.
func (pm *PowerManager) Ping() (bool, error) {
	address := pm.IP
	conn, err := net.Dial("ip4:icmp", address)
	if err != nil {
		return false, fmt.Errorf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Создание ICMP Echo запроса
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   1,
			Seq:  1,
			Data: []byte("PING"),
		},
	}
	bytes, err := msg.Marshal(nil)
	if err != nil {
		return false, fmt.Errorf("failed to marshal message: %v", err)
	}

	// Отправка запроса
	start := time.Now() // Сохраняем текущее время
	if _, err := conn.Write(bytes); err != nil {
		return false, fmt.Errorf("failed to write: %v", err)
	}

	// Чтение ответа
	reply := make([]byte, 1500)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, err := conn.Read(reply)
	if err != nil {
		return false, fmt.Errorf("failed to read: %v", err)
	}

	// Разбор ответа
	rmsg, err := icmp.ParseMessage(1, reply[:n])
	if err != nil {
		return false, fmt.Errorf("failed to parse message: %v", err)
	}
	if rmsg.Type == ipv4.ICMPTypeEchoReply {
		duration := time.Since(start)                       // Измеряем время
		fmt.Printf("Ping to %s took %v\n", pm.IP, duration) // Выводим время пинга
		return true, nil
	}
	return false, nil
}
