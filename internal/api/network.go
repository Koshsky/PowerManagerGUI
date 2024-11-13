package api

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

func ScanNetworkDraft(start, end int) ([]string, error) {
	IPs := []string{}
	for i := start; i < end; i++ {
		IPs = append(IPs, "10.2.1."+strconv.Itoa(i))
	}
	return IPs, nil
}

// ScanNetwork принимает диапазон адресов и возвращает список доступных IP-адресов.
func ScanNetwork(start, end int) ([]string, error) {
	// TODO: узнать значение start end для КУУФС
	var wg sync.WaitGroup
	var mu sync.Mutex
	var reachableIPs []string

	for a := 1; a <= 255; a++ { // operation room
		for b := start; b <= end; b++ {
			ip := fmt.Sprintf("10.2.%d.%d", a, b)
			wg.Add(1)

			go func(ip string) {
				defer wg.Done()
				success, _, err := Ping(ip)
				if err == nil && success {
					mu.Lock()
					reachableIPs = append(reachableIPs, ip)
					mu.Unlock()
				}
			}(ip)
		}
	}

	wg.Wait()
	return reachableIPs, nil
}

// Ping отправляет ICMP Echo запрос на указанный адрес и возвращает результат.
func Ping(address string) (bool, time.Duration, error) {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		return false, 0, fmt.Errorf("failed to create pinger: %v", err)
	}

	pinger.Count = 3
	pinger.Timeout = 2 * time.Second

	err = pinger.Run()
	if err != nil {
		return false, 0, fmt.Errorf("failed to run pinger: %v", err)
	}

	stats := pinger.Statistics()
	if stats.PacketsRecv > 0 {
		return true, stats.AvgRtt, nil
	}

	return false, 0, nil
}
