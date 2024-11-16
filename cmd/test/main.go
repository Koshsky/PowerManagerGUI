package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Loading Indicator Example")

	// Создаем прогресс бар
	progressBar := widget.NewProgressBar()
	progressBar.SetValue(0)

	// Кнопка для запуска сетевой операции
	startButton := widget.NewButton("Start Network Operation", func() {
		// Блокируем кнопку во время загрузки
		startButton.Disable()

		// Запуск сетевой операции в отдельной горутине
		go func(startButton *widget.Button) {
			// Симуляция сетевой операции
			for i := 0; i <= 100; i++ {
				// Обновляем прогресс бар
				progressBar.SetValue(float64(i) / 100)
				time.Sleep(50 * time.Millisecond) // Симуляция задержки
			}

			// Завершение операции
			startButton.Enable()    // Разблокируем кнопку
			progressBar.SetValue(0) // Сбрасываем прогресс бар

			// Обновляем интерфейс (например, показываем сообщение)
			myWindow.SetContent(container.NewVBox(
				widget.NewLabel("Network operation completed!"),
				startButton,
			))
		}(startButton)
	})

	myWindow.SetContent(container.NewVBox(
		startButton,
		layout.NewSpacer(),
		progressBar,
	))
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
