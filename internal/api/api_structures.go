package api

import "fmt"

type PowerManagerInfo struct {
	IP      string `json:"ip"`
	Mask    string `json:"mask"`
	Gateway string `json:"gateway"`
	Status  string `json:"status"`  // Например, статус устройства (включено/выключено)
	Load    int    `json:"load"`    // Нагрузка устройства в процентах
	Uptime  int    `json:"uptime"`  // Время работы устройства в секундах
	Version string `json:"version"` // Версия прошивки или ПО устройства
}

type SensorData struct {
	Temperature string `json:"Temperature"` // Температура в формате "25.0 C"
	Voltage12V  string `json:"12 VDC"`      // Напряжение на линии 12 VDC
	Voltage5V   string `json:"5 VDC"`       // Напряжение на линии 5 VDC
}

type MonitorStatus struct {
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

type GERSStatus struct {
	GERS1Status string `json:"GERS 1"`
	GERS2Status string `json:"GERS 2"`
	GERS3Status string `json:"GERS 3"`
	GERS4Status string `json:"GERS 4"`
	GERS5Status string `json:"GERS 5"`
}

type JSONStringer interface {
	Str() string
}

func (g GERSStatus) Str() string {
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
	info += fmt.Sprintf("IP Address: %s\n", pm.IP)
	info += fmt.Sprintf("Subnet Mask: %s\n", pm.Mask)
	info += fmt.Sprintf("Gateway: %s\n", pm.Gateway)
	info += fmt.Sprintf("Status: %s\n", pm.Status)
	info += fmt.Sprintf("Load: %d%%\n", pm.Load)
	info += fmt.Sprintf("Uptime: %d seconds\n", pm.Uptime)
	info += fmt.Sprintf("Version: %s\n", pm.Version)

	return info
}

func (sd SensorData) Str() string {
	var info string

	info += "Sensor Data Information:\n"
	info += fmt.Sprintf("Temperature: %s\n", sd.Temperature)
	info += fmt.Sprintf("Voltage 12V: %s\n", sd.Voltage12V)
	info += fmt.Sprintf("Voltage 5V: %s\n", sd.Voltage5V)

	return info
}

func (ds MonitorStatus) Str() string {
	var info string

	info += "Device Status Information:\n"
	info += fmt.Sprintf("Mini PC Group 1 Status: %s\n", ds.MiniPCGroup1Status)
	info += fmt.Sprintf("Mini PC Group 2 Status: %s\n", ds.MiniPCGroup2Status)
	info += fmt.Sprintf("Monitor Status: %s\n", ds.MonitorStatus)

	info += fmt.Sprintf("Mini PC Group 1 Relay: %s\n", ds.MiniPCGroup1Relay)
	info += fmt.Sprintf("Mini PC Group 2 Relay: %s\n", ds.MiniPCGroup2Relay)
	info += fmt.Sprintf("Converter Group 1 Relay: %s\n", ds.ConverterGroup1Relay)
	info += fmt.Sprintf("Converter Group 2 Relay: %s\n", ds.ConverterGroup2Relay)
	info += fmt.Sprintf("Monitor Relay: %s\n", ds.MonitorRelay)
	info += fmt.Sprintf("Common Power Relay: %s\n", ds.CommonPowerRelay)
	info += fmt.Sprintf("Reserved 1 Relay: %s\n", ds.Reserved1Relay)
	info += fmt.Sprintf("Reserved 2 Relay: %s\n", ds.Reserved2Relay)

	return info
}
