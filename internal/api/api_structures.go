package api

import "fmt"

type JSONStringer interface {
	Str() string
}

type PowerManagerInfo struct {
	Type              string `json:"Type"`
	Name              string `json:"Name"`
	IP                string `json:"IP"`
	Mask              string `json:"Mask"`
	Gateway           string `json:"Gateway"`
	MAC               string `json:"MAC"`
	PHYAutonegotation string `json:"PHY. Autonegotation"`
	PHYLinkMode       string `json:"PHY. Link mode"`
	PHYLinkSpeed      string `json:"PHY. Link speed"`
	Version           string `json:"Version"`
}

type SensorDataMonitor struct {
	Temperature string `json:"Temperature"`
	Voltage12V  string `json:"12 VDC"`
	Voltage5V   string `json:"5 VDC"`
}
type SensorDataGERS struct {
	Temperature_1 string `json:"Temperature 1"`
	Temperature_2 string `json:"Temperature 2"`
}

type StatusMonitor struct {
	MiniPCGroup1Status   string `json:"[status] Mini PC group 1"`
	MiniPCGroup2Status   string `json:"[status] Mini PC group 2"`
	MonitorStatus        string `json:"[status] Monitor 220 VAC"`
	MiniPCGroup1Relay    string `json:"[relay] Mini PC group 1"`
	MiniPCGroup2Relay    string `json:"[relay] Mini PC group 2"`
	ConverterGroup1Relay string `json:"[relay] Converter group 1"`
	ConverterGroup2Relay string `json:"[relay] Converter group 2"`
	MonitorRelay         string `json:"[relay] Monitor"`
	CommonPowerRelay     string `json:"[relay] Common power"`
	Reserved1Relay       string `json:"[relay] Reserved 1"`
	Reserved2Relay       string `json:"[relay] Reserved 2"`
}

type StatusGERS struct {
	GERS1Status string `json:"GERS 1"`
	GERS2Status string `json:"GERS 2"`
	GERS3Status string `json:"GERS 3"`
	GERS4Status string `json:"GERS 4"`
	GERS5Status string `json:"GERS 5"`
}

func (g StatusGERS) Str() string {
	var info string

	info += "GERS Status Information:\n"
	info += fmt.Sprintf("GERS 1 Status: %s\n", g.GERS1Status)
	info += fmt.Sprintf("GERS 2 Status: %s\n", g.GERS2Status)
	info += fmt.Sprintf("GERS 3 Status: %s\n", g.GERS3Status)
	info += fmt.Sprintf("GERS 4 Status: %s\n", g.GERS4Status)
	info += fmt.Sprintf("GERS 5 Status: %s\n", g.GERS5Status)

	return info
}

func (pm PowerManagerInfo) Str() string {
	var info string

	info += "Power Manager Information:\n"
	info += fmt.Sprintf("Type: %s\n", pm.Type)
	info += fmt.Sprintf("Name: %s\n", pm.Name)
	info += fmt.Sprintf("IP Address: %s\n", pm.IP)
	info += fmt.Sprintf("Subnet Mask: %s\n", pm.Mask)
	info += fmt.Sprintf("Gateway: %s\n", pm.Gateway)
	info += fmt.Sprintf("MAC Address: %s\n", pm.MAC)
	info += fmt.Sprintf("PHY Autonegotation: %s\n", pm.PHYAutonegotation)
	info += fmt.Sprintf("PHY Link Mode: %s\n", pm.PHYLinkMode)
	info += fmt.Sprintf("PHY Link Speed: %s\n", pm.PHYLinkSpeed)
	info += fmt.Sprintf("Version: %s\n", pm.Version)

	return info
}

func (sd SensorDataMonitor) Str() string {
	var info string

	info += "Sensor Data Information:\n"
	info += fmt.Sprintf("Temperature: %s\n", sd.Temperature)
	info += fmt.Sprintf("Voltage 12V: %s\n", sd.Voltage12V)
	info += fmt.Sprintf("Voltage 5V: %s\n", sd.Voltage5V)

	return info
}

func (sd SensorDataGERS) Str() string {
	var info string

	info += "Sensor Data Information:\n"
	info += fmt.Sprintf("Temperature 1: %s\n", sd.Temperature_1)
	info += fmt.Sprintf("Temperature 2: %s\n", sd.Temperature_2)

	return info
}

func (ds StatusMonitor) Str() string {
	var info string

	info += "Device Status Information:\n"
	info += "Feedback:\n"
	info += fmt.Sprintf("  Mini PC Group 1: %s\n", ds.MiniPCGroup1Status)
	info += fmt.Sprintf("  Mini PC Group 2: %s\n", ds.MiniPCGroup2Status)
	info += fmt.Sprintf("  Monitor: %s\n", ds.MonitorStatus)

	info += "Relay:\n"
	info += fmt.Sprintf("  Common Power: %s\n", ds.CommonPowerRelay)
	info += fmt.Sprintf("  Monitor: %s\n", ds.MonitorRelay)
	info += fmt.Sprintf("  Mini PC Group 1: %s\n", ds.MiniPCGroup1Relay)
	info += fmt.Sprintf("  Converter Group 1: %s\n", ds.ConverterGroup1Relay)
	info += fmt.Sprintf("  Mini PC Group 2: %s\n", ds.MiniPCGroup2Relay)
	info += fmt.Sprintf("  Converter Group 2: %s\n", ds.ConverterGroup2Relay)
	info += fmt.Sprintf("  Reserved 1: %s\n", ds.Reserved1Relay)
	info += fmt.Sprintf("  Reserved 2: %s\n", ds.Reserved2Relay)

	return info
}
