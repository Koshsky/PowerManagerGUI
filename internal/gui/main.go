package gui

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Run() {
	myApp := app.New()
	logo, _ := fyne.LoadResourceFromPath("mvs-260.png")
	myApp.SetIcon(logo)
	myWindow := myApp.NewWindow("Power Manager Control")
	InitWindow(myWindow)
	myWindow.ShowAndRun()
}
