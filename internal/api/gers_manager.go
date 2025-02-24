package api

import "slices"

type GERSControlManager struct {
}

func (gcm *GERSControlManager) IsActionAllowedForDevice(device, state string) bool {
	return slices.Contains([]string{"ON", "OFF", "Reset", "HardReset"}, state)
}

func (gcm *GERSControlManager) GetDevices() []string {
	return []string{
		"ALL",
		"GERS 1",
		"GERS 2",
		"GERS 3",
		"GERS 4",
		"GERS 5",
	}
}

func (gcm *GERSControlManager) GetActions() []string {
	return []string{"ON", "OFF", "Reset", "HardReset"}
}
