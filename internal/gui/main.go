package gui

import (
	"fyne.io/fyne/v2/app"
)

func Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Power Manager Control")
	InitWindow(myWindow)
	myWindow.ShowAndRun()
}
