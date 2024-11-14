package netutils

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func ScanNetwork() ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var allReachableIPs []string

	wg.Add(2)

	go func() {
		defer wg.Done()
		gersIPs, err := scanGERSManagers()
		if err == nil {
			mu.Lock()
			allReachableIPs = append(allReachableIPs, gersIPs...)
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		kuufsIPs, err := scanMonitorManagers()
		if err == nil {
			mu.Lock()
			allReachableIPs = append(allReachableIPs, kuufsIPs...)
			mu.Unlock()
		}
	}()

	wg.Wait()
	// return allReachableIPs, nil  // TODO: не забудь исправить это при тестировании у стенда
	return []string{"10.3.1.69", "10.3.1.150"}, nil
}

func scanGERSManagers() ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var reachableIPs []string

	for _, ip := range obtainPossibleIPsForGERSManager() {
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()
			success, _, err := ping(ip, 80)
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

func scanMonitorManagers() ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var reachableIPs []string

	for _, ip := range obtainPossibleIPsForMonitorManager() {
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()
			success, _, err := ping(ip, 80)
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

func ping(address string, port int) (bool, time.Duration, error) {
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
