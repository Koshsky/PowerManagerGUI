package gui

import (
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

func NewManagerTab(p *api.PowerManager) *container.TabItem {
	tabTitle := p.IP

	textDisplay := widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
	button1 := createInfoButton(p, textDisplay)
	button2 := createAnalogButton(p, textDisplay)
	button3 := createStatusButton(p, textDisplay)
	deviceEntry, stateEntry := createDeviceStateEntries()
	button4 := createChangeStateButton(p, deviceEntry, stateEntry, textDisplay)

	content := container.NewVBox(button1, button2, button3, deviceEntry, stateEntry, button4, textDisplay)
	return container.NewTabItem(tabTitle, content)
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

func createDeviceStateEntries() (*widget.Entry, *widget.Entry) {
	deviceEntry := widget.NewEntry()
	deviceEntry.SetPlaceHolder("Device")
	stateEntry := widget.NewEntry()
	stateEntry.SetPlaceHolder("State")
	return deviceEntry, stateEntry
}

func createChangeStateButton(p *api.PowerManager, deviceEntry *widget.Entry, stateEntry *widget.Entry, textDisplay *widget.Label) *widget.Button {
	return widget.NewButton("changeState", func() {
		msg, _ := p.ChangeDeviceState(deviceEntry.Text, stateEntry.Text)
		textDisplay.SetText(msg)
	})
}
