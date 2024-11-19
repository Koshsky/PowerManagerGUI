package gui

import (
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
		changeLabel:  widget.NewLabel("Device: " + pm.Devices[0] + "\nButton pressed:"),
		States:       make(map[string]string),
	}
	go managerTab.UpdateMessage()

	managerTab.messageLabel.Wrapping = fyne.TextWrapWord

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
	changeContainer := mt.createChangeBox()

	content := container.NewVBox(
		mt.createInfoButton(),
		mt.createAnalogButton(),
		mt.createStatusButton(),
		changeContainer,
	)

	return content
}

// TODO: создать для этого контейнера отдельный класс
// Это позволит более крепко связать радио и кнопки, что позволит выполнить следующее:
// TODO: добавить блокировку ON OFF RESET для устройств, управляемых Monitor assembl, кроме mini PC
func (mt *ManagerTab) createChangeBox() *fyne.Container {
	radioGroup := mt.createPatchRadio(mt.powerManager.Devices...)
	changeButtons := mt.createPatchButtons(radioGroup, mt.powerManager.States...)
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, mt.messageLabel)

	return changeContainer
}

func (mt *ManagerTab) createPatchRadio(devices ...string) *widget.RadioGroup {
	radioGroup := widget.NewRadioGroup(devices, func(selected string) {})
	radioGroup.SetSelected(devices[0])
	radioGroup.Required = true
	radioGroup.Horizontal = false
	radioGroup.OnChanged = func(selected string) {
		mt.changeLabelText(selected, mt.States[selected])
	}
	return radioGroup
}

func (mt *ManagerTab) changeLabelText(device, state string) {
	mt.changeLabel.SetText("Device selected: " + device + "\nButton clicked: " + state)
}

func (mt *ManagerTab) createPatchButtons(radioGroup *widget.RadioGroup, states ...string) *fyne.Container {
	buttons := make([]fyne.CanvasObject, len(states)+1)

	for i, state := range states {
		btn := widget.NewButton(state, func(state string, rg *widget.RadioGroup) func() {
			return func() {
				selected := rg.Selected
				mt.LastGet = "get_status"
				mt.States[selected] = state
				mt.powerManager.ChangeState(selected, state)
				mt.changeLabelText(selected, state)
				mt.changeLabelText(selected, mt.States[selected])

			}
		}(state, radioGroup))
		buttons[i] = btn
	}
	buttons[len(states)] = mt.changeLabel

	return container.NewVBox(buttons...)
}

func (mt *ManagerTab) createInfoButton() *widget.Button {
	return widget.NewButton("get_info", func() {
		mt.LastGet = "get_info"
	})
}

func (mt *ManagerTab) createAnalogButton() *widget.Button {
	return widget.NewButton("get_analog", func() {
		mt.LastGet = "get_analog"
	})
}

func (mt *ManagerTab) createStatusButton() *widget.Button {
	return widget.NewButton("get_status", func() {
		mt.LastGet = "get_status"
	})
}
