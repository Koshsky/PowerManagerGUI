package gui

import (
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
	OREntry.SetText(operatingRoom) // по умолчанию

	scanBtn := hub.createScanButton(OREntry)

	instructionLabel := widget.NewLabel("Please enter the operating room number (1-255):")
	instructionContainer := container.NewHBox(instructionLabel, OREntry)

	content := container.NewVBox(instructionContainer, scanBtn, hub.messageLabel)
	return container.NewTabItem(tabTitle, content)
}

// createScanButton создает кнопку для сканирования сети и обновления.
func (h *Hub) createScanButton(OREntry *widget.Entry) *widget.Button {
	return widget.NewButton("SCAN NETWORK AND REFRESH", func() {
		h.ScanAndRefresh(OREntry.Text)
	})
}

// ScanAndRefresh выполняет сканирование сети и обновляет вкладки.
func (h *Hub) ScanAndRefresh(operatingRoom string) {
	loadingDialog := dialog.NewProgressInfinite("Network scanning", "Please, wait...", h.app.Window)
	loadingDialog.Show()
	defer loadingDialog.Hide()

	IPs, err := netutils.ScanNetwork(operatingRoom)
	if err != nil {
		h.messageLabel.SetText(err.Error())
		return
	}

	powerManagers := api.CreatePowerManagers(IPs)

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

// createMessageLabel создает метку для отображения сообщений.
func createMessageLabel() *widget.Label {
	return widget.NewLabel("There will be more information here soon, but for now just enjoy the emptiness!")
}
