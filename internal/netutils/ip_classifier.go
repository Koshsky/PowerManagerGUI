package netutils

import (
	"strconv"
)

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
