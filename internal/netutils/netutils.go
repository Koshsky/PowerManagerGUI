package netutils

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func IsValidOctet(input string) bool {
	octet, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return octet >= 1 && octet <= 255
}

func ScanNetwork(operatingRoom string) ([]string, error) {
	if !IsValidOctet(operatingRoom) {
		return []string{}, fmt.Errorf("ScanNetwork: incorrect operating room number [1-255]: %s", operatingRoom)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var allReachableIPs []string

	roomNum, err := strconv.Atoi(operatingRoom)
	if err != nil {
		return allReachableIPs, fmt.Errorf("ScanNetwork: %w", err)
	} else if roomNum < 1 || roomNum > 255 {
		return allReachableIPs, fmt.Errorf("ScanNetwork: the operating room number is outside the range 1-255")
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		gersIPs, _ := scanGERSManagers(operatingRoom)
		mu.Lock()
		allReachableIPs = append(allReachableIPs, gersIPs...)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		kuufsIPs, _ := scanMonitorManagers(operatingRoom)
		mu.Lock()
		allReachableIPs = append(allReachableIPs, kuufsIPs...)
		mu.Unlock()
	}()

	wg.Wait()
	return allReachableIPs, nil
}

func scanGERSManagers(operatingRoom string) ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var reachableIPs []string

	for _, ip := range obtainPossibleIPsForGERSManager(operatingRoom) {
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()
			success, _, err := Ping(ip, 80)
			if err == nil && success {
				mu.Lock()
				reachableIPs = append(reachableIPs, ip)
				mu.Unlock()
			}
		}(ip)
	}

	wg.Wait()
	return reachableIPs, nil
}

func scanMonitorManagers(operatingRoom string) ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var reachableIPs []string

	for _, ip := range obtainPossibleIPsForMonitorManager(operatingRoom) {
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()
			success, _, err := Ping(ip, 80)
			if err == nil && success {
				mu.Lock()
				reachableIPs = append(reachableIPs, ip)
				mu.Unlock()
			}
		}(ip)
	}

	wg.Wait()
	return reachableIPs, nil
}

func Ping(address string, port int) (bool, time.Duration, error) {
	startTime := time.Now()

	addr := fmt.Sprintf("%s:%d", address, port)

	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return false, 0, fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	duration := time.Since(startTime)

	return true, duration, nil
}
