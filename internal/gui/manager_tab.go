package gui

import (
	"fmt"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

func NewManagerTab(pm *api.PowerManager) (*container.TabItem, error) {
	tabTitle := pm.IP

	MessageLabel := widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
	MessageLabel.Wrapping = fyne.TextWrapWord // TODO: выяснить имеет ли это вообще смысл

	var changeContainer *fyne.Container
	if pm.Type == "GERS control" {
		changeContainer = createGERSChangeBox(pm, MessageLabel)
	} else if pm.Type == "Monitor assembly (3.0V)" {
		changeContainer = createMonitorChangeBox(pm, MessageLabel)
	} else {
		return nil, fmt.Errorf("NewManagerTab: cannot create new tab: uknown type of PowerManager: %v", pm.Type)
	}

	content := container.NewVBox(
		createInfoButton(pm, MessageLabel),
		createAnalogButton(pm, MessageLabel),
		createStatusButton(pm, MessageLabel),
		changeContainer,
	)

	return container.NewTabItem(tabTitle, content), nil
}

func createMonitorChangeBox(p *api.PowerManager, textDisplay *widget.Label) *fyne.Container {
	radioGroup := createPatchRadio(
		"Mini PC 1",
		"Mini PC 2",
		"Converter 1",
		"Converter 2",
		"Monitor",
		"Common Power",
		"Reserved 1",
		"Reserved 2",
	)

	changeButtons := createPatchButtons(p, textDisplay, radioGroup, "ON", "OFF", "Reset", "Turn ON", "Turn OFF")
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, textDisplay)

	return changeContainer
}

func createGERSChangeBox(p *api.PowerManager, textDisplay *widget.Label) *fyne.Container {
	radioGroup := createPatchRadio(
		"ALL",
		"GERS 1",
		"GERS 2",
		"GERS 3",
		"GERS 4",
		"GERS 5",
	)
	changeButtons := createPatchButtons(p, textDisplay, radioGroup, "ON", "OFF", "Reset", "HardReset")
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

func createPatchButtons(p *api.PowerManager, textDisplay *widget.Label, radioGroup *widget.RadioGroup, states ...string) *fyne.Container {
	buttons := make([]fyne.CanvasObject, len(states))

	for i, state := range states {
		btn := widget.NewButton(state, func(state string, rg *widget.RadioGroup) func() {
			return func() {
				device := rg.Selected
				textDisplay.SetText(state + " clicked, RadioGroup selected: " + device)
				p.ChangeState(device, state)
			}
		}(state, radioGroup))
		buttons[i] = btn
	}

	return container.NewVBox(buttons...)
}

func createInfoButton(p *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_info", func() {
		if info, err := p.GetInfo(); err == nil {
			textDisplay.SetText(info.Str())
		} else {
			textDisplay.SetText(err.Error())
		}
	})
}

func createAnalogButton(p *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_analog", func() {
		if data, err := p.GetAnalog(); err == nil {
			textDisplay.SetText(data.Str())
		} else {
			textDisplay.SetText(err.Error())
		}
	})
}

func createStatusButton(p *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_status", func() {
		if status, err := p.GetStatus(); err == nil {
			textDisplay.SetText(status.Str())
		} else {
			textDisplay.SetText(err.Error())
		}
	})
}
