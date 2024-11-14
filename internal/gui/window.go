package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InitWindow(myWindow fyne.Window) {
	tabTitle := "HUB"

	messageLabel := createMessageLabel()
	button := createScanButton(myWindow)

	content := container.NewVBox(button, messageLabel)
	tab := container.NewTabItem(tabTitle, content)
	tabsItems := container.NewAppTabs()
	tabsItems.Append(tab)
	myWindow.SetContent(tabsItems)
	myWindow.Resize(fyne.NewSize(800, 600))
}

func createMessageLabel() *widget.Label {
	return widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
}

func createScanButton(myWindow fyne.Window) *widget.Button {
	return widget.NewButton("SCAN NETWORK AND REFRESH", func() {
		ScanAndRefresh(myWindow)
	})
}
