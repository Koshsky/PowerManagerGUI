package netutils

import (
	"fmt"
	"strconv"
	"time"

	nmap "github.com/Ullaakut/nmap/v2"
)

func ScanNetwork(operatingRoom string) ([]string, error) {
	if !isValidOctet(operatingRoom) {
		return []string{}, fmt.Errorf("ScanNetwork: incorrect operating room number [1-255]: %s", operatingRoom)
	}

	roomNum, err := strconv.Atoi(operatingRoom)
	if err != nil {
		return []string{}, fmt.Errorf("ScanNetwork: %w", err)
	} else if roomNum < 1 || roomNum > 255 {
		return []string{}, fmt.Errorf("ScanNetwork: the operating room number is outside the range 1-255")
	}

	subnet := fmt.Sprintf("10.4.%d.0/24", roomNum)
	return scanSubnet(subnet)
}

func scanSubnet(subnet string) ([]string, error) {
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(subnet),
		nmap.WithPingScan(),
		nmap.WithHostTimeout(1*time.Second), // Устанавливаем таймаут в 1 секунду
	)
	if err != nil {
		return []string{}, fmt.Errorf("error creating scanner: %v", err)
	}

	result, _, err := scanner.Run()
	if err != nil {
		return []string{}, fmt.Errorf("scanning error: %v", err)
	}
	print(result)

	var allReachableIPs []string
	for _, host := range result.Hosts {
		if host.Status.State == "up" {
			allReachableIPs = append(allReachableIPs, host.Addresses[0].Addr)
		}
	}
	return allReachableIPs, nil
}

func isValidOctet(input string) bool {
	octet, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return octet >= 1 && octet <= 255
}
