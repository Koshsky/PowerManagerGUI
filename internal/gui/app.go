package gui

import (
	"log"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type App struct {
	App     fyne.App
	Window  fyne.Window
	Hub     container.TabItem
}

func NewApp() *App {
	a := &App{
		App: app.New(),
	}

	logo, err := fyne.LoadResourceFromPath("mvs-260.png")
	if err != nil {
		log.Println(err.Error())
	}
	a.App.SetIcon(logo)
	a.Window = a.App.NewWindow("Power Manager Control")
	a.InitWindow()

	return a
}

func (a *App) InitWindow() {
	tab := a.NewHub("1")
	tabsItems := container.NewAppTabs()
	tabsItems.Append(tab)
	a.Window.SetContent(tabsItems)
	a.Window.Resize(fyne.NewSize(700, 600))
}

func (a *App) Run() {
	a.Window.ShowAndRun()
}
