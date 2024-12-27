package gui

import (
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Koshsky/PowerManagerGUI/internal/api"
	"github.com/Koshsky/PowerManagerGUI/internal/netutils"
)

type Hub struct {
	app           *App
	operatingRoom string
	messageLabel  *widget.Label
}

func (a *App) NewHub(operatingRoom string) *container.TabItem {
	hub := &Hub{
		app:           a,
		operatingRoom: operatingRoom,
		messageLabel:  createMessageLabel(),
	}

	tabTitle := "HUB"

	OREntry := widget.NewEntry()
	OREntry.SetPlaceHolder("Enter the operating room number...")
	OREntry.SetText(operatingRoom)

	instructionLabel := widget.NewLabel("Please enter the operating room number (1-255):")
	instructionContainer := container.NewHBox(instructionLabel, OREntry)
	check := widget.NewCheck("SCAN ALL", func(checked bool) {
		if checked {
			OREntry.Disable()
		} else {
			OREntry.Enable()
		}
	})
	scanButton := hub.createScanButton(OREntry, check, "Scan network and refresh")

	content := container.NewVBox(instructionContainer, check, scanButton, hub.messageLabel)
	return container.NewTabItem(tabTitle, content)
}

func (h *Hub) createScanButton(OREntry *widget.Entry, check *widget.Check, label string) *widget.Button {
	return widget.NewButton(label, func() {
		progressBar := widget.NewProgressBarInfinite()
		loadingDialog := dialog.NewCustomWithoutButtons("Network scanning", container.NewVBox(progressBar), h.app.Window)
		loadingDialog.Show()
		defer loadingDialog.Hide()

		newTabsItems := container.NewAppTabs()
		hub := h.app.NewHub(OREntry.Text)
		newTabsItems.Append(hub)

		powerManagers := []*api.PowerManager{}
		if !check.Checked {
			powerManagers = h.scanSelected(OREntry.Text)
		} else {
			powerManagers = h.scanAll()
		}
		for _, pm := range powerManagers {
			newTab := NewManagerTab(pm)
			newTabsItems.Append(newTab)
		}

		h.app.Window.SetContent(newTabsItems)
	})
}

func (h *Hub) scanSelected(operatingRoom string) []*api.PowerManager {
	IPs, err := netutils.ScanNetwork(operatingRoom)
	if err != nil {
		h.messageLabel.SetText(err.Error())
		return nil
	}
	return api.BuildPowerManagers(IPs)
}

func (h *Hub) scanAll() []*api.PowerManager {
	powerManagers := []*api.PowerManager{}
	for or := 1; or < 256; or++ {  // TODO: add parallelism and test at the bench
		powerManagers = append(powerManagers, h.scanSelected(strconv.Itoa(or))...)
	}
	return powerManagers
}

func createMessageLabel() *widget.Label {
	return widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
}
