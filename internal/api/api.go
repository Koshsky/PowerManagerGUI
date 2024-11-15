package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (pm *PowerManager) GetInfo() (PowerManagerInfo, error) {
	url := fmt.Sprintf("http://%s/get_info.json", pm.IP)
	response, err := http.Get(url)
	if err != nil {
		return PowerManagerInfo{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return PowerManagerInfo{}, fmt.Errorf("failed to get info: %s", response.Status)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return PowerManagerInfo{}, fmt.Errorf("unexpected Content-Type: %s", contentType)
	}

	// Use io.TeeReader to log the response body to the terminal
	tee := io.TeeReader(response.Body, os.Stdout)

	var info PowerManagerInfo
	if err := json.NewDecoder(tee).Decode(&info); err != nil {
		return PowerManagerInfo{}, err
	}

	return info, nil
}

func (pm *PowerManager) GetAnalog() (SensorData, error) {
	url := fmt.Sprintf("http://%s/get_analog.json", pm.IP)
	response, err := http.Get(url)
	if err != nil {
		return SensorData{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return SensorData{}, fmt.Errorf("failed to get sensor data: %s", response.Status)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return SensorData{}, fmt.Errorf("unexpected Content-Type: %s", contentType)
	}

	// Используем io.TeeReader для логирования содержимого в терминал
	tee := io.TeeReader(response.Body, os.Stdout)

	var data SensorData
	if err := json.NewDecoder(tee).Decode(&data); err != nil {
		return SensorData{}, err
	}

	return data, nil
}

func (pm *PowerManager) GetStatus() (JSONStringer, error) {
	url := fmt.Sprintf("http://%s/get_status.json", pm.IP)
	response, err := http.Get(url)
	if err != nil {
		return MonitorStatus{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return MonitorStatus{}, fmt.Errorf("failed to get status: %s", response.Status)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return MonitorStatus{}, fmt.Errorf("unexpected Content-Type: %s", contentType)
	}

	// Используем io.TeeReader для логирования содержимого в терминал
	tee := io.TeeReader(response.Body, os.Stdout)

	if pm.ManagerType == "GERSManager" {
		var status GERSStatus
		if err := json.NewDecoder(tee).Decode(&status); err != nil {
			return GERSStatus{}, err
		}
		return status, nil
	} else if pm.ManagerType == "MonitorManager" {
		var status MonitorStatus
		if err := json.NewDecoder(tee).Decode(&status); err != nil {
			return MonitorStatus{}, err
		}
		return status, nil
	} else {
		return MonitorStatus{}, nil
	}
}

type Command struct {
	Device string `json:"Device"`
	State  string `json:"state"`
}

func (pm *PowerManager) ChangeDeviceState(device string, command string) (string, error) {
	url := fmt.Sprintf("http://%s/changeState.json", pm.IP)

	cmd := Command{
		Device: device,
		State:  command,
	}

	jsonData, err := json.Marshal(cmd)
	if err != nil {
		msg := fmt.Sprintf("JSON encoding error: %v", err)
		return msg, fmt.Errorf(msg)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		msg := fmt.Sprintf("error creating request: %v", err)
		return msg, fmt.Errorf(msg)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		msg := fmt.Sprintf("error while executing request: %v", err)
		return msg, fmt.Errorf(msg)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("failed to change device state, status: %d", resp.StatusCode)
		return msg, fmt.Errorf(msg)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := fmt.Sprintf("error reading response: %v", err)
		return msg, fmt.Errorf(msg)
	}

	return string(body), nil
}
