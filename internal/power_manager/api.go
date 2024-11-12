package power_manager

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (pm *PowerManager) GetInfo() error {
	url := fmt.Sprintf("http://%s/get_info.json", pm.IP)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get info: %s", response.Status)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return fmt.Errorf("unexpected Content-Type: %s", contentType)
	}

	var info PowerManagerInfo
	if err := json.NewDecoder(response.Body).Decode(&info); err != nil {
		return err
	}

	info.Print() // TODO: изменить это поведение в соответствии с ТЗ
	return nil
}

func (pm *PowerManager) GetAnalog() error {
	url := fmt.Sprintf("http://%s/get_analog.json", pm.IP)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get sensor data: %s", response.Status)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return fmt.Errorf("unexpected Content-Type: %s", contentType)
	}

	var data SensorData
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return err
	}

	data.Print() // TODO: изменить это поведение в соответствии с ТЗ
	return nil
}

func (pm *PowerManager) GetStatus() error {
	url := fmt.Sprintf("http://%s/get_status.json", pm.IP)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get status: %s", response.Status)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return fmt.Errorf("unexpected Content-Type: %s", contentType)
	}

	var status DeviceStatus
	if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
		return err
	}

	status.Print() // TODO: изменить это поведение в соответствии с ТЗ
	return nil
}

type Command struct {
	Device string `json:"Device"`
	State  string `json:"state"`
}

// func (pm *PowerManager) ChangeDeviceState(device string, command string) error {
// 	url := fmt.Sprintf("http://%s/changeState.json", pm.IP)

// 	cmd := Command{
// 		Device: device,
// 		State:  command,
// 	}

// 	jsonData, err := json.Marshal(cmd)
// 	if err != nil {
// 		return fmt.Errorf("ошибка при кодировании в JSON: %v", err)
// 	}

// 	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return fmt.Errorf("ошибка при создании запроса: %v", err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("ошибка при выполнении запроса: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	// Проверяем статус ответа
// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("не удалось изменить состояние устройства, статус: %d", resp.StatusCode)
// 	}

// 	fmt.Println("Команда получена - тело:", resp.Status)
// 	return nil
// }
