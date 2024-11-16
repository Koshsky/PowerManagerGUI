package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

func ScanAndRefresh(myWindow fyne.Window) {
	// IPs, _ := netutils.ScanNetwork()
	IPs := []string{"10.4.1.5", "10.4.1.30"}
	powerManagers := api.CreatePowerManagers(IPs)

	hub := NewHubTab(myWindow)
	newTabsItems := container.NewAppTabs()
	newTabsItems.Append(hub)
	for _, pm := range powerManagers {
		log.Println(pm)
		newTab, err := NewManagerTab(pm)
		if err != nil {
			log.Println("new tub error: ", err)
			continue
		}
		newTabsItems.Append(newTab)
	}
	myWindow.SetContent(newTabsItems)
}
