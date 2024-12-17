package gui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

type ManagerTab struct {
	powerManager *api.PowerManager
	messageLabel *widget.Label
	changeBox    ChangeBox
	lastGet      string
}

func NewManagerTab(pm *api.PowerManager) *container.TabItem {
	managerTab := &ManagerTab{
		powerManager: pm,
		messageLabel: widget.NewLabel(""),
		lastGet:      "get_info",
	}
	managerTab.InitChangeBox()
	go managerTab.UpdateMessage()

	managerTab.messageLabel.Wrapping = fyne.TextWrapWord

	tabTitle := pm.IP
	content := managerTab.createContent()

	return container.NewTabItem(tabTitle, content)
}
func (mt *ManagerTab) UpdateMessage() {
	for {
		switch mt.lastGet {
		case "get_info":
			mt.updateLabelFromFunc(mt.powerManager.GetInfo)
		case "get_analog":
			mt.updateLabelFromFunc(mt.powerManager.GetAnalog)
		case "get_status":
			mt.updateLabelFromFunc(mt.powerManager.GetStatus)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (mt *ManagerTab) updateLabelFromFunc(getFunc func() (api.JSONStringer, error)) {
	if info, err := getFunc(); err == nil {
		mt.messageLabel.SetText(info.Str())
	}
}

func (mt *ManagerTab) createContent() *fyne.Container {
	content := container.NewVBox(
		mt.newGetButton("get_info"),
		mt.newGetButton("get_analog"),
		mt.newGetButton("get_status"),
		mt.changeBox.Box,
	)

	return content
}
func (mt *ManagerTab) newGetButton(method string) *widget.Button {
	return widget.NewButton(method, func() {
		mt.lastGet = method
	})
}
