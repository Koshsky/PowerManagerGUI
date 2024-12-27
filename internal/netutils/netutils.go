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
		return []string{}, fmt.Errorf("ScanNetwork: %w", err)
	} else if roomNum < 1 || roomNum > 255 {
		return []string{}, fmt.Errorf("ScanNetwork: the operating room number is outside the range 1-255")
	}

	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("10.4.%d.%d", roomNum, i)
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()
			success, _, err := Ping(ip, 80)
			if err == nil && success {
				mu.Lock()
				allReachableIPs = append(allReachableIPs, ip)
				mu.Unlock()
			}
		}(ip)
	}

	wg.Wait()
	return allReachableIPs, nil
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
