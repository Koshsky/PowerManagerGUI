package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func InitWindow(myWindow fyne.Window) {
	tab := NewHub(myWindow, "1")
	tabsItems := container.NewAppTabs()
	tabsItems.Append(tab)
	myWindow.SetContent(tabsItems)
	myWindow.Resize(fyne.NewSize(700, 500))
}
