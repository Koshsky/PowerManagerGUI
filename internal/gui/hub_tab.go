package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Koshsky/PowerManagerGUI/internal/api"
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
	OREntry.SetText(operatingRoom) // by default

	scanBtn := hub.createScanButton(OREntry)

	instructionLabel := widget.NewLabel("Please enter the operating room number (1-255):")
	instructionContainer := container.NewHBox(instructionLabel, OREntry)

	content := container.NewVBox(instructionContainer, scanBtn, hub.messageLabel)
	return container.NewTabItem(tabTitle, content)
}

func (h *Hub) createScanButton(OREntry *widget.Entry) *widget.Button {
	return widget.NewButton("SCAN NETWORK AND REFRESH", func() {
		h.ScanAndRefresh(OREntry.Text)
	})
}

func (h *Hub) ScanAndRefresh(operatingRoom string) {
	loadingDialog := dialog.NewProgressInfinite("Network scanning", "Please, wait...", h.app.Window)
	loadingDialog.Show()
	defer loadingDialog.Hide()

	IPs := []string{"10.4.1.5", "10.4.1.30"}

	powerManagers := api.BuildPowerManagers(IPs)

	hub := h.app.NewHub(operatingRoom)
	newTabsItems := container.NewAppTabs()
	newTabsItems.Append(hub)
	for _, pm := range powerManagers {
		if newTab, err := NewManagerTab(pm); err == nil {
			newTabsItems.Append(newTab)
		}
	}
	h.app.Window.SetContent(newTabsItems)
}

func createMessageLabel() *widget.Label {
	return widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
}
