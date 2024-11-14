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
	textDisplay.Wrapping = fyne.TextWrapWord // Установим оборачивание текста

	button1 := createInfoButton(p, textDisplay)
	button2 := createAnalogButton(p, textDisplay)
	button3 := createStatusButton(p, textDisplay)

	var changeContainer *fyne.Container
	if p.ManagerType == "GERSManager" {
		changeContainer = createGERSManagerChangeContainer(textDisplay)
	} else if p.ManagerType == "MonitorManager" {
		changeContainer = createMonitorManagerChangeContainer(textDisplay)
	}

	content := container.NewVBox(
		button1, button2, button3,
		changeContainer,
	)

	// Устанавливаем минимальный и максимальный размер для content
	content.Resize(fyne.NewSize(600, 400))     // Установка начального размера

	return container.NewTabItem(tabTitle, content)
}

func createMonitorManagerChangeContainer(textDisplay *widget.Label) *fyne.Container {
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
	radioGroup.Horizontal = false // Вертикальное расположение

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

	// Создаем и возвращаем changeContainer с ограничением по размерам
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, textDisplay)
	changeContainer.Resize(fyne.NewSize(500, 300))     // Установка начального размера

	return changeContainer
}

func createGERSManagerChangeContainer(textDisplay *widget.Label) *fyne.Container {
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
	radioGroup.Horizontal = false // Вертикальное расположение

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

	// Создаем и возвращаем changeContainer с ограничением по размерам
	changeContainer := container.NewAdaptiveGrid(3, radioGroup, changeButtons, textDisplay)
	changeContainer.Resize(fyne.NewSize(500, 300))   // Установка начального размера

	return changeContainer
}

func createInfoButton(p *api.PowerManager, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("get_info", func() {
		if info, err := p.GetInfo(); err == nil {
			textDisplay.SetText(info.Str())
		} else {
			textDisplay.SetText(api.Draft())
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
		if status, err := p.GetInfo(); err == nil {
			textDisplay.SetText(status.Str())
		} else {
			log.Println(err)
		}
	})
}
