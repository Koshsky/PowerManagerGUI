package gui

import (
	"log"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

func NewManagerTab(p *api.PowerManager) *container.TabItem {
	tabTitle := p.IP

	textDisplay := widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
	textDisplay.Wrapping = fyne.TextWrapWord // TODO: выяснить имеет ли это вообще смысл

	var changeContainer *fyne.Container
	if p.Type == "GERSManager" {
		changeContainer = createGERSChangeBox(textDisplay)
	} else if p.Type == "MonitorManager" {
		changeContainer = createMonitorChangeBox(textDisplay)
	}

	content := container.NewVBox(
		createInfoButton(p, textDisplay),
		createAnalogButton(p, textDisplay),
		createStatusButton(p, textDisplay),
		changeContainer,
	)

	return container.NewTabItem(tabTitle, content)
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
		println("Selected:", selected)
		textDisplay.SetText("Selected: " + selected)
	})
	radioGroup.Horizontal = false

	btnON := widget.NewButton("ON", func() {
		println("ON clicked")
		textDisplay.SetText("ON clicked")
	})
	btnOFF := widget.NewButton("OFF", func() {
		println("OFF clicked")
		textDisplay.SetText("OFF clicked")
	})
	btnReset := widget.NewButton("Reset", func() {
		println("Reset clicked")
		textDisplay.SetText("Reset clicked")
	})
	btnTurnON := widget.NewButton("Turn ON", func() {
		println("Turn ON clicked")
		textDisplay.SetText("Turn ON clicked")
	})
	btnTurnOFF := widget.NewButton("Turn OFF", func() {
		println("Turn OFF clicked")
		textDisplay.SetText("Turn OFF clicked")
	})

	changeButtons := container.NewVBox(btnON, btnOFF, btnReset, btnTurnON, btnTurnOFF)
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, textDisplay)

	return changeContainer
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

	btnON := widget.NewButton("ON", func() {
		println("ON clicked")
		textDisplay.SetText("ON clicked")
	})
	btnOFF := widget.NewButton("OFF", func() {
		println("OFF clicked")
		textDisplay.SetText("OFF clicked")
	})
	btnReset := widget.NewButton("Reset", func() {
		println("Reset clicked")
		textDisplay.SetText("Reset clicked")
	})
	btnHardReset := widget.NewButton("HardReset", func() {
		println("HardReset clicked")
		textDisplay.SetText("HardReset clicked")
	})

	changeButtons := container.NewVBox(btnON, btnOFF, btnReset, btnHardReset)

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
