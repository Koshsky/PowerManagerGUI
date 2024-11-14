package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/Koshsky/PowerManagerGUI/internal/api"
	"github.com/Koshsky/PowerManagerGUI/internal/netutils"
)

func ScanAndRefresh(myWindow fyne.Window) {
	IPs, _ := netutils.ScanNetwork() // TODO: протестировать ScanNetwork, ЗАМЕНИТЬ ЗАГЛУШКУ ИМ!!!
	powerManagers := api.CreatePowerManagers(IPs)

	HUB := myWindow.Content().(*container.AppTabs).Items[0]
	newTabsItems := container.NewAppTabs(HUB)
	for _, pm := range powerManagers {
		newTab := NewManagerTab(pm)
		newTabsItems.Append(newTab)
	}
	myWindow.SetContent(newTabsItems)
}
