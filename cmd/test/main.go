package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Radio Group Example")

	// Создаем переменную для хранения выбранного элемента
	selected := "Mini PC 1"

	radioGroup := widget.NewRadioGroup([]string{
		"Mini PC 1",
		"Mini PC 2",
		"Converter 1",
		"Converter 2",
		"Monitor",
		"Common Power",
		"Reserved 1",
		"Reserved 2",
	}, func(selectedValue string) {
		log.Println("Selected:", selectedValue)
		if selectedValue != "" {
			selected = selectedValue
		}
	})
	radioGroup.Required = true
	// Установка первой кнопки как выбранной по умолчанию
	radioGroup.SetSelected(selected)

	// Создаем CheckBox, который будет использоваться для фиксации выбора
	fixCheckBox := widget.NewCheck("Fix selection", func(checked bool) {
		if !checked {
			// Если CheckBox не отмечен, возвращаем предыдущий выбор
			radioGroup.SetSelected(selected)
		}
	})

	// Добавляем обработчик для RadioGroup
	radioGroup.OnChanged = func(value string) {
		if value != "" {
			selected = value
		}
	}

	// Создаем контейнер и добавляем элементы
	content := container.NewVBox(radioGroup, fixCheckBox)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
