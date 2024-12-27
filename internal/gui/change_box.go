package gui

import (
	"fmt"
	"slices"
	"strings"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Koshsky/PowerManagerGUI/internal/api"
)

type ChangeBox struct {
	RG          *widget.RadioGroup
	Buttons     *fyne.Container
	ChangeLabel *widget.Label
	Box         *fyne.Container
	MT          *ManagerTab
	States      map[string]string
}

func (mt *ManagerTab) InitChangeBox() *fyne.Container {
	cb := ChangeBox{
		MT:          mt,
		States:      make(map[string]string),
		ChangeLabel: widget.NewLabel("Device: " + mt.powerManager.Devices[0] + "\nButton pressed: no one"),
	}
	for _, device := range mt.powerManager.Devices {
		cb.States[device] = "no one"
	}
	cb.initRadio(mt.powerManager.Devices...)
	cb.initButtons(mt.powerManager.States...)
	cb.Box = container.NewAdaptiveGrid(3, cb.RG, cb.Buttons, mt.messageLabel)
	mt.changeBox = cb
	return cb.Box
}

func (cb *ChangeBox) initRadio(devices ...string) {
	cb.RG = widget.NewRadioGroup(devices, func(selected string) {
	})
	cb.RG.SetSelected(devices[0])
	cb.RG.Required = true
	cb.RG.Horizontal = false
	cb.RG.OnChanged = func(selected string) { cb.handleRadioChange(selected) }
}

func (cb *ChangeBox) initButtons(states ...string) {
	buttons := make([]fyne.CanvasObject, len(states)+1)

	for i, state := range states {
		btn := widget.NewButton(state, func() {
			cb.handleButtonClick(state)
		})

		buttons[i] = btn
	}
	buttons[len(states)] = cb.ChangeLabel

	cb.Buttons = container.NewVBox(buttons...)
}

func (cb *ChangeBox) handleButtonClick(cmd string) {
	device := cb.RG.Selected
	cb.MT.lastGet = "get_status"
	if !cb.isDeviceActionAllowed(device, cmd) {
		cb.ChangeLabel.SetText(cmd + " command is not allowed for " + device)
	} else {
		if device == "ALL" { // special case for ALL GERS
			for device := range cb.States {
				cb.States[device] = cmd
			}
		} else {
			cb.States[device] = cmd
			if _, exists := cb.States["ALL"]; exists {
				cb.States["ALL"] = ""
			}
		}

		if err := cb.MT.powerManager.ChangeState(device, cmd); err != nil {
			fmt.Println("failed to change state: %w", err)
		}
		cb.ChangeLabel.SetText("Device selected: " + device + "\nButton clicked: " + cmd)
	}
}

func (cb *ChangeBox) handleRadioChange(selected string) {
	cb.ChangeLabel.SetText("Device selected: " + selected + "\nButton clicked: " + cb.States[selected])
	for _, obj := range cb.Buttons.Objects {
		if button, ok := obj.(*widget.Button); ok {
			if cb.isDeviceActionAllowed(selected, button.Text) {
				button.Enable()
			} else {
				button.Disable()
			}
		}
	}
}

func (cb *ChangeBox) isDeviceActionAllowed(device, state string) bool {  // TODO: убрать как метод PowerManager.
	if slices.Contains([]string{"ON", "OFF", "Reset"}, state) {
		return cb.MT.powerManager.Type == api.GERSControl || strings.HasPrefix(device, "Mini PC")
	}
	return true
}
