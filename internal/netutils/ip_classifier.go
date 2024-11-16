package netutils

import (
	"fmt"
	"strconv"
)

func obtainPossibleIPsForMonitorManager(operatingRoom string) []string {
	possibleIPs := []string{}

	roomNum, err := strconv.Atoi(operatingRoom)
	if err != nil || roomNum < 1 || roomNum > 255 {
		return possibleIPs
	}

	for i := 30; i <= 39; i++ {
		possibleIPs = append(possibleIPs, fmt.Sprintf("10.4.%d.%d", roomNum, i))
	}
	return possibleIPs
}

func obtainPossibleIPsForGERSManager(operatingRoom string) []string {
	possibleIPs := []string{}

	roomNum, err := strconv.Atoi(operatingRoom)
	if err != nil || roomNum < 1 || roomNum > 255 {
		return possibleIPs
	}
	
	possibleIPs = append(possibleIPs, fmt.Sprintf("10.4.%d.5", roomNum))
	return possibleIPs
}
