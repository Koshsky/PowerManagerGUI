package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InitWindow(myWindow fyne.Window) {
	tab := NewHubTab(myWindow)
	tabsItems := container.NewAppTabs()
	tabsItems.Append(tab)
	myWindow.SetContent(tabsItems)
	myWindow.Resize(fyne.NewSize(700, 500))
}

func NewHubTab(myWindow fyne.Window) *container.TabItem {
	tabTitle := "HUB"

	messageLabel := createMessageLabel()
	button := createScanButton(myWindow)

	content := container.NewVBox(button, messageLabel)
	return container.NewTabItem(tabTitle, content)
}

func createMessageLabel() *widget.Label {
	return widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
}

func createScanButton(myWindow fyne.Window) *widget.Button {
	return widget.NewButton("SCAN NETWORK AND REFRESH", func() {
		ScanAndRefresh(myWindow)
	})
}
