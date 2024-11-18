package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

type ManagerTab struct {
	powerManager *api.PowerManager
	messageLabel *widget.Label
}

func NewManagerTab(pm *api.PowerManager) (*container.TabItem, error) {
	managerTab := &ManagerTab{
		powerManager: pm,
		messageLabel: widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!"),
	}
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
	buttons := make([]fyne.CanvasObject, len(states))

	for i, state := range states {
		btn := widget.NewButton(state, func(state string, rg *widget.RadioGroup) func() {
			return func() {
				device := rg.Selected
				mt.messageLabel.SetText(state + " clicked, RadioGroup selected: " + device)
				mt.powerManager.ChangeState(device, state)
			}
		}(state, radioGroup))
		buttons[i] = btn
	}

	return container.NewVBox(buttons...)
}

func (mt *ManagerTab) createInfoButton() *widget.Button {
	return widget.NewButton("get_info", func() {
		if info, err := mt.powerManager.GetInfo(); err == nil {
			mt.messageLabel.SetText(info.Str())
		} else {
			mt.messageLabel.SetText(err.Error())
		}
	})
}

func (mt *ManagerTab) createAnalogButton() *widget.Button {
	return widget.NewButton("get_analog", func() {
		if data, err := mt.powerManager.GetAnalog(); err == nil {
			mt.messageLabel.SetText(data.Str())
		} else {
			mt.messageLabel.SetText(err.Error())
		}
	})
}

func (mt *ManagerTab) createStatusButton() *widget.Button {
	return widget.NewButton("get_status", func() {
		if status, err := mt.powerManager.GetStatus(); err == nil {
			mt.messageLabel.SetText(status.Str())
		} else {
			mt.messageLabel.SetText(err.Error())
		}
	})
}
