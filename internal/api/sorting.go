package api

import (
	"strconv"
	"strings"
)

type ByType []*PowerManager

func (a ByType) Len() int {
	return len(a)
}

func (a ByType) Less(i, j int) bool {
	if a[i].Type == GERSControl && a[j].Type == GERSControl {
		lastOctetI := getLastOctet(a[i].IP)
		lastOctetJ := getLastOctet(a[j].IP)
		return lastOctetI < lastOctetJ
	}
	return a[i].Type == GERSControl
}

func (a ByType) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func getLastOctet(ip string) int {
	parts := strings.Split(ip, ".")
	lastOctet, _ := strconv.Atoi(parts[3])
	return lastOctet
}
