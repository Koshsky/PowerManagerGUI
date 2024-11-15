package netutils

import (
	"slices"
	"strconv"
)

func GetManagerTypeByIP(ip string) string {
	if isGERSManager(ip) {
		return "GERSManager"
	} else if isMonitorManager(ip) {
		return "MonitorManager"
	}
	return "Unknown"
}

func obtainPossibleIPsForMonitorManager() []string {
	possibleIPs := []string{}
	for OR := 1; OR < 256; OR++ {
		for i := 30; i <= 39; i++ {
			possibleIPs = append(possibleIPs, "10.4."+strconv.Itoa(OR)+"."+strconv.Itoa(i))
		}
	}
	return possibleIPs
}

func obtainPossibleIPsForGERSManager() []string {
	possibleIPs := []string{}
	for OR := 1; OR < 256; OR++ {
		possibleIPs = append(possibleIPs, "10.4."+strconv.Itoa(OR)+".5")
	}
	return possibleIPs
}

func isMonitorManager(ip string) bool {
	return slices.Contains(obtainPossibleIPsForMonitorManager(), ip)
}

func isGERSManager(ip string) bool {
	return slices.Contains(obtainPossibleIPsForGERSManager(), ip)
}
