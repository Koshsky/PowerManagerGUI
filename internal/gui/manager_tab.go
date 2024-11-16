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

	var changeContainer *fyne.Container
	if info, err := pm.GetInfo(); err != nil {
		return nil, err
	} else if info.Type == "GERS control" {
		changeContainer = createGERSChangeBox(MessageLabel)
	} else if info.Type == "Monitor assembly (3.0V)" {
		changeContainer = createMonitorChangeBox(MessageLabel)
	}

	content := container.NewVBox(
		createInfoButton(pm, MessageLabel),
		createAnalogButton(pm, MessageLabel),
		createStatusButton(pm, MessageLabel),
		changeContainer,
	)

	return container.NewTabItem(tabTitle, content), nil
}

func createMonitorChangeBox(textDisplay *widget.Label) *fyne.Container {
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

	changeButtons := createPatchButtons(textDisplay, radioGroup, "ON", "OFF", "Reset", "Turn ON", "Turn OFF")
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, textDisplay)

	return changeContainer
}

func createGERSChangeBox(textDisplay *widget.Label) *fyne.Container {
	radioGroup := createPatchRadio(
		"ALL",
		"GERS 1",
		"GERS 2",
		"GERS 3",
		"GERS 4",
		"GERS 5",
	)
	changeButtons := createPatchButtons(textDisplay, radioGroup, "ON", "OFF", "Reset", "HardReset")
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

func createPatchButtons(textDisplay *widget.Label, radioGroup *widget.RadioGroup, texts ...string) *fyne.Container {
	buttons := make([]fyne.CanvasObject, len(texts))

	for i, text := range texts {
		btn := widget.NewButton(text, func(text string, rg *widget.RadioGroup) func() {
			return func() { // TODO: изменить поведение этой функции
				selected := rg.Selected
				println(text + " clicked, RadioGroup selected: " + selected)
				textDisplay.SetText(text + " clicked, RadioGroup selected: " + selected)
			}
		}(text, radioGroup))
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
