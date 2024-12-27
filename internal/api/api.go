package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
)

func (pm *PowerManager) GetInfo() (JSONStringer, error) {
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

	var info PowerManagerInfo
	if err := json.NewDecoder(response.Body).Decode(&info); err != nil {
		return PowerManagerInfo{}, err
	}
	pm.Type = info.Type // update powerManager with type info
	return info, nil
}

func (pm *PowerManager) GetAnalog() (JSONStringer, error) {
	url := fmt.Sprintf("http://%s/get_analog.json", pm.IP)
	response, err := http.Get(url)
	if err != nil {
		return SensorDataMonitor{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return SensorDataMonitor{}, fmt.Errorf("failed to get sensor data: %s", response.Status)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return SensorDataMonitor{}, fmt.Errorf("unexpected Content-Type: %s", contentType)
	}

	if pm.Type == GERSControl {
		var data SensorDataGERS
		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			return StatusGERS{}, err
		}
		return data, nil
	} else if pm.Type == MonitorAssembly {
		var data SensorDataMonitor
		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			return StatusMonitor{}, err
		}
		return data, nil
	} else {
		return StatusMonitor{}, fmt.Errorf("unknown type of power manager: %s", pm.Type)
	}
}

func (pm *PowerManager) GetStatus() (JSONStringer, error) {
	url := fmt.Sprintf("http://%s/get_status.json", pm.IP)
	response, err := http.Get(url)
	if err != nil {
		return StatusMonitor{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return StatusMonitor{}, fmt.Errorf("failed to get status: %s", response.Status)
	}
	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		return StatusMonitor{}, fmt.Errorf("unexpected Content-Type: %s", contentType)
	}

	if pm.Type == GERSControl {
		var status StatusGERS
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			return StatusGERS{}, err
		}
		return status, nil
	} else if pm.Type == MonitorAssembly {
		var status StatusMonitor
		if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
			return StatusMonitor{}, err
		}
		return status, nil
	} else {
		return StatusMonitor{}, fmt.Errorf("unknown type of power manager: %s", pm.Type)
	}
}

func (pm *PowerManager) ChangeState(device, state string) error {
	url := fmt.Sprintf("http://%s/changeState.json", pm.IP)

	data, err := pm.prepareChangeStateBody(device, state)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("JSON encoding error: %v", err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error while executing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to change device state, status: %d", resp.StatusCode)
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	return nil
}

func (pm *PowerManager) prepareChangeStateBody(device, cmd string) (map[string]string, error) {
	if !slices.Contains(pm.Devices, device) {
		return nil, fmt.Errorf("prepareChangeStateBody: unknown device: %s", device)
	} else if !slices.Contains(pm.States, cmd) {
		return nil, fmt.Errorf("prepareChangeStateBody: unknown command: %s", cmd)
	}
	if pm.Type == GERSControl {
		for i := 0; i < len(pm.Devices); i++ {
			if pm.Devices[i] == device {
				return map[string]string{"GERS": strconv.Itoa(i), "state": cmd}, nil
			}
		}
	}
	return map[string]string{"Device": device, "state": cmd}, nil
}
