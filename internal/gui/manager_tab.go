package gui

import (
	"log"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

func NewManagerTab(p *api.PowerManager) (*container.TabItem, error) {
	tabTitle := p.IP

	textDisplay := widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
	textDisplay.Wrapping = fyne.TextWrapWord // TODO: выяснить имеет ли это вообще смысл

	var changeContainer *fyne.Container
	if info, err := p.GetInfo(); err != nil {
		return nil, err
	} else if info.Type == "GERS control" {
		changeContainer = createGERSChangeBox(textDisplay)
	} else if info.Type == "Monitor assembly (3.0V)" {
		changeContainer = createMonitorChangeBox(textDisplay)
	}

	content := container.NewVBox(
		createInfoButton(p, textDisplay),
		createAnalogButton(p, textDisplay),
		createStatusButton(p, textDisplay),
		changeContainer,
	)

	return container.NewTabItem(tabTitle, content), nil
}

func createMonitorChangeBox(textDisplay *widget.Label) *fyne.Container {
	radioGroup := widget.NewRadioGroup([]string{
		"Mini PC 1",
		"Mini PC 2",
		"Converter 1",
		"Converter 2",
		"Monitor",
		"Common Power",
		"Reserved 1",
		"Reserved 2",
	}, func(selected string) {
		log.Println("Selected:", selected)
		textDisplay.SetText("Selected: " + selected)
	})
	radioGroup.Horizontal = false

	changeButtons := createPatchButtons(textDisplay, radioGroup, "ON", "OFF", "Reset", "Turn ON", "Turn OFF")
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, textDisplay)

	return changeContainer
}

func createPatchButtons(textDisplay *widget.Label, radioGroup *widget.RadioGroup, texts ...string) *fyne.Container {
	buttons := make([]fyne.CanvasObject, len(texts))

	for i, text := range texts {
		btn := widget.NewButton(text, func(text string, rg *widget.RadioGroup) func() {
			return func() {
				selected := rg.Selected
				println(text + " clicked, RadioGroup selected: " + selected)
				textDisplay.SetText(text + " clicked, RadioGroup selected: " + selected)
			}
		}(text, radioGroup))
		buttons[i] = btn
	}

	return container.NewVBox(buttons...)
}

func createGERSChangeBox(textDisplay *widget.Label) *fyne.Container {
	radioGroup := widget.NewRadioGroup([]string{
		"ALL",
		"GERS 1",
		"GERS 2",
		"GERS 3",
		"GERS 4",
		"GERS 5",
	}, func(selected string) {
		println("Selected:", selected)
		textDisplay.SetText("Selected: " + selected)
	})
	radioGroup.Horizontal = false

	changeButtons := createPatchButtons(textDisplay, radioGroup, "ON", "OFF", "Reset", "HardReset")

	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, textDisplay)

	return changeContainer
}

func createInfoButton(p *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_info", func() {
		if info, err := p.GetInfo(); err == nil {
			textDisplay.SetText(info.Str())
		} else {
			log.Println(err)
		}
	})
}

func createAnalogButton(p *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_analog", func() {
		if data, err := p.GetAnalog(); err == nil {
			textDisplay.SetText(data.Str())
		} else {
			log.Println(err)
		}
	})
}

func createStatusButton(p *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_status", func() {
		if status, err := p.GetStatus(); err == nil {
			textDisplay.SetText(status.Str())
		} else {
			log.Println(err)
		}
	})
}
