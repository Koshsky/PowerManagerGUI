package gui

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

func NewManagerTab(pm *api.PowerManager) (*container.TabItem, error) {
	tabTitle := pm.IP

	MessageLabel := widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
	MessageLabel.Wrapping = fyne.TextWrapWord // TODO: выяснить имеет ли это вообще смысл

	changeContainer := createChangeBox(pm, MessageLabel)

	content := container.NewVBox(
		createInfoButton(pm, MessageLabel),
		createAnalogButton(pm, MessageLabel),
		createStatusButton(pm, MessageLabel),
		changeContainer,
	)

	return container.NewTabItem(tabTitle, content), nil
}

func createChangeBox(pm *api.PowerManager, textDisplay *widget.Label) *fyne.Container {
	radioGroup := createPatchRadio(pm.Devices...)
	changeButtons := createPatchButtons(pm, textDisplay, radioGroup, pm.States...)
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, textDisplay)

	return changeContainer
}

func createPatchRadio(texts ...string) *widget.RadioGroup {
	radioGroup := widget.NewRadioGroup(texts, func(selected string) {})
	radioGroup.SetSelected(texts[0])
	radioGroup.Required = true
	radioGroup.Horizontal = false
	return radioGroup
}

func createPatchButtons(pm *api.PowerManager, textDisplay *widget.Label, radioGroup *widget.RadioGroup, states ...string) *fyne.Container {
	buttons := make([]fyne.CanvasObject, len(states))

	for i, state := range states {
		btn := widget.NewButton(state, func(state string, rg *widget.RadioGroup) func() {
			return func() {
				device := rg.Selected
				textDisplay.SetText(state + " clicked, RadioGroup selected: " + device)
				pm.ChangeState(device, state)
			}
		}(state, radioGroup))
		buttons[i] = btn
	}

	return container.NewVBox(buttons...)
}

func createInfoButton(pm *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_info", func() {
		if info, err := pm.GetInfo(); err == nil {
			textDisplay.SetText(info.Str())
		} else {
			textDisplay.SetText(err.Error())
		}
	})
}

func createAnalogButton(pm *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_analog", func() {
		if data, err := pm.GetAnalog(); err == nil {
			textDisplay.SetText(data.Str())
		} else {
			textDisplay.SetText(err.Error())
		}
	})
}

func createStatusButton(pm *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_status", func() {
		if status, err := pm.GetStatus(); err == nil {
			textDisplay.SetText(status.Str())
		} else {
			textDisplay.SetText(err.Error())
		}
	})
}
