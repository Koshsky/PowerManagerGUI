package power_manager

import "fmt"

// PowerManagerInfo представляет информацию о Power Manager
type PowerManagerInfo struct {
	IP      string `json:"ip"`
	Mask    string `json:"mask"`
	Gateway string `json:"gateway"`
	Status  string `json:"status"`  // Например, статус устройства (включено/выключено)
	Load    int    `json:"load"`    // Нагрузка устройства в процентах
	Uptime  int    `json:"uptime"`  // Время работы устройства в секундах
	Version string `json:"version"` // Версия прошивки или ПО устройства
}

// SensorData представляет данные с термопары и напряжение на линиях 12 и 5 VDC
type SensorData struct {
	Temperature string `json:"Temperature"` // Температура в формате "25.0 C"
	Voltage12V  string `json:"12 VDC"`      // Напряжение на линии 12 VDC
	Voltage5V   string `json:"5 VDC"`       // Напряжение на линии 5 VDC
}

// DeviceStatus представляет текущее состояние устройств и реле
type DeviceStatus struct {
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

func (pm *PowerManagerInfo) Print() {
	fmt.Println("Power Manager Information:")

	fmt.Printf("IP Address: %s\n", pm.IP)
	fmt.Printf("Subnet Mask: %s\n", pm.Mask)
	fmt.Printf("Gateway: %s\n", pm.Gateway)
	fmt.Printf("Status: %s\n", pm.Status)
	fmt.Printf("Load: %d%%\n", pm.Load)
	fmt.Printf("Uptime: %d seconds\n", pm.Uptime)
	fmt.Printf("Version: %s\n", pm.Version)
}

func (sd *SensorData) Print() {
	fmt.Println("Sensor Data Information:")

	fmt.Printf("Temperature: %s\n", sd.Temperature)
	fmt.Printf("Voltage 12V: %s\n", sd.Voltage12V)
	fmt.Printf("Voltage 5V: %s\n", sd.Voltage5V)
}

func (ds *DeviceStatus) Print() {
	fmt.Println("Device Status Information:")

	fmt.Printf("Mini PC Group 1 Status: %s\n", ds.MiniPCGroup1Status)
	fmt.Printf("Mini PC Group 2 Status: %s\n", ds.MiniPCGroup2Status)
	fmt.Printf("Monitor Status: %s\n", ds.MonitorStatus)

	fmt.Printf("Mini PC Group 1 Relay: %s\n", ds.MiniPCGroup1Relay)
	fmt.Printf("Mini PC Group 2 Relay: %s\n", ds.MiniPCGroup2Relay)
	fmt.Printf("Converter Group 1 Relay: %s\n", ds.ConverterGroup1Relay)
	fmt.Printf("Converter Group 2 Relay: %s\n", ds.ConverterGroup2Relay)
	fmt.Printf("Monitor Relay: %s\n", ds.MonitorRelay)
	fmt.Printf("Common Power Relay: %s\n", ds.CommonPowerRelay)
	fmt.Printf("Reserved 1 Relay: %s\n", ds.Reserved1Relay)
	fmt.Printf("Reserved 2 Relay: %s\n", ds.Reserved2Relay)
}
