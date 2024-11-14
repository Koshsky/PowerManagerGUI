package netutils

import (
	"slices"
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
	// possibleIPs := []string{}
	// for OR := 1; OR < 256; OR++ {
	// 	for i := 30; i <= 39; i++ {
	// 		possibleIPs = append(possibleIPs, "10.4."+strconv.Itoa(OR)+"."+strconv.Itoa(i))
	// 	}
	// }
	// return possibleIPs  // TODO: когда согласует карту сети изменить это поведение
	return []string{"10.3.1.69"}
}

func obtainPossibleIPsForGERSManager() []string {
	// possibleIPs := []string{}
	// for OR := 1; OR < 256; OR++ {
	// 	possibleIPs = append(possibleIPs, "10.4."+strconv.Itoa(OR)+".5")
	// }
	// return possibleIPs  // TODO: когда согласует карту сети изменить это поведение
	return []string{"10.3.1.150"}
}

func isMonitorManager(ip string) bool {
	return slices.Contains(obtainPossibleIPsForMonitorManager(), ip)
}

func isGERSManager(ip string) bool {
	return slices.Contains(obtainPossibleIPsForGERSManager(), ip)
}
