package gui

import (
	"log"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Run() {
	myApp := app.New()
	logo, err := fyne.LoadResourceFromPath("mvs-260.png")
	if err != nil {
		log.Println(err.Error())
	}
	myApp.SetIcon(logo)
	myWindow := myApp.NewWindow("Power Manager Control")
	InitWindow(myWindow)
	myWindow.ShowAndRun()
}
