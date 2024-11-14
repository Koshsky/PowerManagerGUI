package netutils

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func ScanGERS() ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var reachableIPs []string

	for x := 1; x < 256; x++ {
		ip := "10.3." + strconv.Itoa(x) + ".150"
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()
			success, _, err := Ping(ip, 80) // TODO: вынести порт в перменную, спросить порт
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

func ScanKUUFS() ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var reachableIPs []string

	for x := 1; x < 256; x++ {
		ip := "10.3." + strconv.Itoa(x) + "." + "69"
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()
			success, _, err := Ping(ip, 80) // TODO: вынести порт в перменную, спросить порт
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

func ScanNetwork() ([]string, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var allReachableIPs []string

	wg.Add(2)

	go func() {
		defer wg.Done()
		gersIPs, err := ScanGERS()
		if err == nil {
			mu.Lock()
			allReachableIPs = append(allReachableIPs, gersIPs...)
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		kuufsIPs, err := ScanKUUFS()
		if err == nil {
			mu.Lock()
			allReachableIPs = append(allReachableIPs, kuufsIPs...)
			mu.Unlock()
		}
	}()

	wg.Wait()
	return allReachableIPs, nil
}

func ScanNetworkDraft() ([]string, error) {
	reachableIPs := []string{}
	for i := 30; i < 39; i++ {
		reachableIPs = append(reachableIPs, "10.4.1."+strconv.Itoa(i))
	}
	return reachableIPs, nil
}

// Ping отправляет TCP запрос на указанный адрес и возвращает результат.
func Ping(address string, port int) (bool, time.Duration, error) {
	startTime := time.Now()

	// Формируем адрес для подключения
	addr := fmt.Sprintf("%s:%d", address, port)

	// Пытаемся установить TCP соединение
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return false, 0, fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Вычисляем время, затраченное на соединение
	duration := time.Since(startTime)

	return true, duration, nil
}
