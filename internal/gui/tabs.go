package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/Koshsky/PowerManagerGUI/internal/api"
	"github.com/Koshsky/PowerManagerGUI/internal/netutils"
)

func ScanAndRefresh(myWindow fyne.Window) {
	IPs, _ := netutils.ScanNetwork()
	powerManagers := api.CreatePowerManagers(IPs)

	hub := NewHubTab(myWindow)
	newTabsItems := container.NewAppTabs()
	newTabsItems.Append(hub)
	for _, pm := range powerManagers {
		newTab := NewManagerTab(pm)
		newTabsItems.Append(newTab)
	}
	myWindow.SetContent(newTabsItems)
}
