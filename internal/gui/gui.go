package gui

import (
	"log"
	"strconv"

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
	// TODO: установить значения по умолчанию для этих ентрис В СООТВЕТСТВИЕ С ФАКТИЧЕСКИМИМ ЗНАЧЕНИЯМИ
	startLabel := widget.NewLabel("Minimum value of the 4th octet of the IP address")
	startEntry := widget.NewEntry()
	startEntry.SetPlaceHolder("Start value")
	startEntry.SetText("100")

	endLabel := widget.NewLabel("Maximum value of the 4th octet of the IP address")
	endEntry := widget.NewEntry()
	endEntry.SetPlaceHolder("End value")
	endEntry.SetText("120")

	messageLabel := widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
	button := widget.NewButton("SCAN NETWORK AND REFRESH", func() {
		a, err1 := strconv.Atoi(startEntry.Text)
		b, err2 := strconv.Atoi(endEntry.Text)
		if err1 != nil {
			messageLabel.SetText("Start: " + err1.Error())
		} else if err2 != nil {
			messageLabel.SetText("End: " + err2.Error())
		} else {
			IPs, _ := api.ScanNetworkDraft(a, b)
			RefreshTabs(myWindow, IPs)
		}
	})

	content := container.NewVBox(
		startLabel, startEntry,
		endLabel, endEntry,
		button,
		messageLabel,
	)
	tab := container.NewTabItem(tabTitle, content)
	tabsItems := container.NewAppTabs()
	tabsItems.Append(tab)
	myWindow.SetContent(tabsItems)
	myWindow.Resize(fyne.NewSize(800, 600))
}

func RefreshTabs(myWindow fyne.Window, IPs []string) {
	powerManagers := api.CreatePowerManagers(IPs)

	tabsItems := myWindow.Content().(*container.AppTabs)

	for i := len(tabsItems.Items) - 1; i > 0; i-- {
		tabsItems.Remove(tabsItems.Items[i])
	}
	for i := 0; i < len(powerManagers); i++ {
		newTab := NewManagerTab(powerManagers[i])
		tabsItems.Append(newTab)
	}

	myWindow.SetContent(tabsItems)
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
