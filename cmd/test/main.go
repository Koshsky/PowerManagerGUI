package main

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Equal Width Elements")

	// Создаем элементы
	textDisplay := widget.NewLabel("Text Display")
	radioGroup := widget.NewRadioGroup([]string{"Option 1", "Option 2", "Option 3"}, func(selected string) {
		println("Selected:", selected)
	})
	buttonContainer := container.NewVBox(
		widget.NewButton("Button 1", func() { println("Button 1 clicked") }),
		widget.NewButton("Button 2", func() { println("Button 2 clicked") }),
	)

	// Создаем контейнер с равными ширинами
	SomeContainer := container.New(layout.NewGridLayout(3), textDisplay, radioGroup, buttonContainer)

	// Устанавливаем контейнер как содержимое окна
	myWindow.SetContent(SomeContainer)
	myWindow.Resize(fyne.NewSize(600, 200)) // Устанавливаем размер окна
	myWindow.ShowAndRun()                    // Запускаем приложение
}