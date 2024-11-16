package gui

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Koshsky/PowerManagerGUI/internal/api"
	"github.com/Koshsky/PowerManagerGUI/internal/netutils"
)

func NewHub(myWindow fyne.Window, operatingRoom string) *container.TabItem {
	tabTitle := "HUB"

	messageLabel := createMessageLabel() // for error printing

	OREntry := widget.NewEntry()
	OREntry.SetPlaceHolder("Enter the operating room number...")
	OREntry.SetText(operatingRoom) // by default

	scanBtn := createScanButton(myWindow, OREntry, messageLabel)

	instructionLabel := widget.NewLabel("Please enter the operating room number (1-255):")

	instructionContainer := container.NewHBox(instructionLabel, OREntry)

	content := container.NewVBox(instructionContainer, scanBtn, messageLabel)
	return container.NewTabItem(tabTitle, content)
}

func ScanAndRefresh(myWindow fyne.Window, operatingRoom string, textDisplay *widget.Label) {
	loadingDialog := dialog.NewProgressInfinite("Network scanning", "Please, wait...", myWindow)
	loadingDialog.Show()
	defer loadingDialog.Hide()

	IPs, err := netutils.ScanNetwork(operatingRoom)
	if err != nil {
		textDisplay.SetText(err.Error())
		return
	}

	powerManagers := api.CreatePowerManagers(IPs)

	hub := NewHub(myWindow, operatingRoom)
	newTabsItems := container.NewAppTabs()
	newTabsItems.Append(hub)
	for _, pm := range powerManagers {
		if newTab, err := NewManagerTab(pm); err == nil {
			newTabsItems.Append(newTab)
		}
	}
	myWindow.SetContent(newTabsItems)
}

func createMessageLabel() *widget.Label {
	return widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
}

func createScanButton(myWindow fyne.Window, OREntry *widget.Entry, messageLabel *widget.Label) *widget.Button {
	return widget.NewButton("SCAN NETWORK AND REFRESH", func() {
		ScanAndRefresh(myWindow, OREntry.Text, messageLabel)
	})
}
