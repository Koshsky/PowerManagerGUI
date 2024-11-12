package gui

import (
	"log"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

func Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Power Manager Control")
	// TODO: РЕАЛИЗОВАТЬ СКАНИРОВАНИЕ, ПОЛУЧИТЬ СПИСОК IPS
	IPs := []string{"10.2.1.121", "10.2.1.122", "10.2.2.121", "10.2.3.121"} // TODO: не забыть убрать это!!

	powerManagers := api.CreatePowerManagers(IPs)

	tabsItems := container.NewAppTabs()
	for _, p := range powerManagers {
		tab := NewTab(p)
		tabsItems.Append(tab)
	}
	myWindow.SetContent(tabsItems)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

func NewTab(p *api.PowerManager) *container.TabItem {
	tabTitle := p.IP

	textDisplay := widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
	button1 := widget.NewButton("get_info", func() {
		if info, err := p.GetInfo(); err == nil {
			textDisplay.SetText(info.Str())
		} else {
			log.Println(err)
		}
	})
	button2 := widget.NewButton("get_analog", func() {
		if data, err := p.GetAnalog(); err == nil {
			textDisplay.SetText(data.Str())
		} else {
			log.Println(err)
		}
	})
	button3 := widget.NewButton("get_status", func() {
		if status, err := p.GetInfo(); err == nil {
			textDisplay.SetText(status.Str())
		} else {
			log.Println(err)
		}
	})
	deviceLabel := widget.NewEntry()
	deviceLabel.SetPlaceHolder("Device")
	stateLabel := widget.NewEntry()
	stateLabel.SetPlaceHolder("State")
	button4 := widget.NewButton("changeState", func() {
		msg, _ := p.ChangeDeviceState(deviceLabel.Text, stateLabel.Text)
		textDisplay.SetText(msg)
	})
	content := container.NewVBox(
		button1, button2, button3,
		widget.NewLabel(""),
		deviceLabel, stateLabel,
		button4,
		textDisplay)
	tab := container.NewTabItem(tabTitle, content)
	return tab
}
