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
	LastGet      string
	changeLabel  *widget.Label
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
		time.Sleep(1 * time.Second)
	}
}

func (mt *ManagerTab) updateLabelFromFunc(getFunc func() (api.JSONStringer, error)) {
	if info, err := getFunc(); err == nil {
		mt.messageLabel.SetText(info.Str())
	} else {
		mt.messageLabel.SetText(err.Error())
	}
}

func NewManagerTab(pm *api.PowerManager) (*container.TabItem, error) {
	managerTab := &ManagerTab{
		powerManager: pm,
		messageLabel: widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!"),
		changeLabel:  widget.NewLabel("Device:\nState:"),
	}
	go managerTab.UpdateMessage()

	managerTab.messageLabel.Wrapping = fyne.TextWrapWord // TODO: выяснить имеет ли это вообще смысл

	tabTitle := pm.IP
	content := managerTab.createContent()

	return container.NewTabItem(tabTitle, content), nil
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

func (mt *ManagerTab) createChangeBox() *fyne.Container {
	radioGroup := mt.createPatchRadio(mt.powerManager.Devices...)
	changeButtons := mt.createPatchButtons(radioGroup, mt.powerManager.States...)
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, mt.messageLabel)

	return changeContainer
}

func (mt *ManagerTab) createPatchRadio(texts ...string) *widget.RadioGroup {
	radioGroup := widget.NewRadioGroup(texts, func(selected string) {})
	radioGroup.SetSelected(texts[0])
	radioGroup.Required = true
	radioGroup.Horizontal = false
	return radioGroup
}

func (mt *ManagerTab) createPatchButtons(radioGroup *widget.RadioGroup, states ...string) *fyne.Container {
	buttons := make([]fyne.CanvasObject, len(states)+1)

	for i, state := range states {
		btn := widget.NewButton(state, func(state string, rg *widget.RadioGroup) func() {
			return func() {
				device := rg.Selected
				mt.LastGet = "get_status"
				mt.powerManager.ChangeState(device, state)
				mt.changeLabel.SetText("Device selected: " + device + "\nButton clicked: " + state)
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
