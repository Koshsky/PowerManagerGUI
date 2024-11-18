package main

import (
	"log"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	// Загрузка иконки из файла
	logo, err := fyne.LoadResourceFromPath("mvs-260.png")
	if err != nil {
		log.Println(err.Error())
	}
	myApp.SetIcon(logo) // Укажите путь к вашему файлу иконки

	w := myApp.NewWindow("Мое приложение")
	w.SetContent(container.NewVBox(
		widget.NewLabel("Привет, Fyne!"),
	))

	w.ShowAndRun()
}
