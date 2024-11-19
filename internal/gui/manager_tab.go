package gui

import (
	"slices"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

type ManagerTab struct {
	powerManager *api.PowerManager
	messageLabel *widget.Label
	changeLabel  *widget.Label
	LastGet      string
	States       map[string]string
}

func NewManagerTab(pm *api.PowerManager) (*container.TabItem, error) {
	managerTab := &ManagerTab{
		powerManager: pm,
		messageLabel: widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!"),
		changeLabel:  widget.NewLabel("Device: " + pm.Devices[0] + "\nButton pressed: no one"),
		States:       make(map[string]string),
	}
	go managerTab.UpdateMessage()

	managerTab.messageLabel.Wrapping = fyne.TextWrapWord
	managerTab.changeLabel.Wrapping = fyne.TextWrapWord

	tabTitle := pm.IP
	content := managerTab.createContent()

	for _, device := range pm.Devices {
		managerTab.States[device] = ""
	}

	return container.NewTabItem(tabTitle, content), nil
}
func (mt *ManagerTab) UpdateMessage() {
	for {
		switch mt.LastGet {
		case "get_info":
			mt.updateLabelFromFunc(mt.powerManager.GetInfo)
		case "get_analog":
			mt.updateLabelFromFunc(mt.powerManager.GetAnalog)
		case "get_status":
			mt.updateLabelFromFunc(mt.powerManager.GetStatus)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (mt *ManagerTab) updateLabelFromFunc(getFunc func() (api.JSONStringer, error)) {
	if info, err := getFunc(); err == nil { // TODO: есть ли необходимость логировать ошибки?
		mt.messageLabel.SetText(info.Str())
	}
}

func (mt *ManagerTab) createContent() *fyne.Container {
	changeContainer := mt.NewChangeBox()

	content := container.NewVBox(
		mt.newGetButton("get_info"),
		mt.newGetButton("get_analog"),
		mt.newGetButton("get_status"),
		changeContainer,
	)

	return content
}

func (mt *ManagerTab) NewChangeBox() *fyne.Container {
	radioGroup := mt.createRadio(mt.powerManager.Devices...)
	changeButtons := mt.createPatchButtons(radioGroup, mt.powerManager.States...)
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, mt.messageLabel)

	return changeContainer
}

func (mt *ManagerTab) createRadio(devices ...string) *widget.RadioGroup {
	radioGroup := widget.NewRadioGroup(devices, func(selected string) {})
	radioGroup.SetSelected(devices[0])
	radioGroup.Required = true
	radioGroup.Horizontal = false
	radioGroup.OnChanged = func(device string) {
		mt.changeLabel.SetText("Device selected: " + device + "\nButton clicked: " + mt.States[device])
	}
	return radioGroup
}

func (mt *ManagerTab) createPatchButtons(radioGroup *widget.RadioGroup, states ...string) *fyne.Container {
	buttons := make([]fyne.CanvasObject, len(states)+1)

	for i, state := range states {
		btn := widget.NewButton(state, func() {
			mt.handleButtonClick(state, radioGroup)
		})

		buttons[i] = btn
	}
	buttons[len(states)] = mt.changeLabel

	return container.NewVBox(buttons...)
}

func (mt *ManagerTab) handleButtonClick(cmd string, rg *widget.RadioGroup) {
	device := rg.Selected
	mt.LastGet = "get_status"
	if !mt.isDeviceActionAllowed(device, cmd) {
		mt.changeLabel.SetText(cmd + " isn't allowed for " + device)
	} else {
		if device == "ALL" { // special case for ALL GERS (not only GERS btw)
			for device := range mt.States {
				mt.States[device] = cmd
			}
		}
		mt.States[device] = cmd
		// mt.powerManager.ChangeState(device, cmd)
		mt.changeLabel.SetText("Device selected: " + device + "\nButton clicked: " + cmd)
	}
}

func (mt *ManagerTab) isDeviceActionAllowed(device, state string) bool {
	if slices.Contains([]string{"ON", "OFF", "Reset"}, state) {
		return mt.powerManager.Type == "GERS control" || strings.HasPrefix(device, "Mini PC")
	}
	return true
}

func (mt *ManagerTab) newGetButton(method string) *widget.Button {
	return widget.NewButton(method, func() {
		mt.LastGet = method
	})
}
