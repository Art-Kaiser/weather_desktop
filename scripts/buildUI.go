package scripts

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	owm "github.com/briandowns/openweathermap"
)

func InitUI(res *owm.CurrentWeatherData) {
	app := app.New()
	window := app.NewWindow("Watch the weather")

	hello := widget.NewLabel("Прогноз на день")

	window.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Обновить", func() {
			fmt.Println(res)
			hello.SetText("res")
		}),

		widget.NewLabel("Прогноз на неделю"),
	))

	window.ShowAndRun()
}
