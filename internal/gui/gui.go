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
	InitWindow(myWindow)
	myWindow.ShowAndRun()
}

func InitWindow(myWindow fyne.Window) {
	tabTitle := "HUB"

	// TODO: такое сообщение используется очень часто. вынести как константу!
	messageLabel := widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
	button := widget.NewButton("SCAN NETWORK AND REFRESH", func() {
		ScanAndRefresh(myWindow)
	})

	content := container.NewVBox(
		button,
		messageLabel,
	)
	tab := container.NewTabItem(tabTitle, content)
	tabsItems := container.NewAppTabs()
	tabsItems.Append(tab)
	myWindow.SetContent(tabsItems)
	myWindow.Resize(fyne.NewSize(800, 600))
}

func ScanAndRefresh(myWindow fyne.Window) {
	IPs, _ := api.ScanNetworkDraft() // TODO: протестировать ScanNetwork, ЗАМЕНИТЬ ЗАГЛУШКУ ИМ!!!
	powerManagers := api.CreatePowerManagers(IPs)

	HUB := myWindow.Content().(*container.AppTabs).Items[0]
	newTabsItems := container.NewAppTabs(HUB)
	for i := 0; i < len(powerManagers); i++ {
		newTab := NewManagerTab(powerManagers[i])
		newTabsItems.Append(newTab)
	}
	myWindow.SetContent(newTabsItems)
}

func NewManagerTab(p *api.PowerManager) *container.TabItem {
	tabTitle := p.IP

	textDisplay := widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
	button1 := widget.NewButton("get_info", func() {
		if info, err := p.GetInfo(); err == nil {
			textDisplay.SetText(info.Str())
		} else {
			textDisplay.SetText(api.Draft())
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
	deviceEntry := widget.NewEntry()
	deviceEntry.SetPlaceHolder("Device")
	stateEntry := widget.NewEntry()
	stateEntry.SetPlaceHolder("State")
	button4 := widget.NewButton("changeState", func() {
		msg, _ := p.ChangeDeviceState(deviceEntry.Text, stateEntry.Text)
		textDisplay.SetText(msg)
	})
	content := container.NewVBox(
		button1, button2, button3,
		deviceEntry, stateEntry,
		button4,
		textDisplay)
	tab := container.NewTabItem(tabTitle, content)
	return tab
}
