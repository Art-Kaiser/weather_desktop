package scripts

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	owm "github.com/briandowns/openweathermap"
	"time"
)

func InitUI(res *owm.CurrentWeatherData) {
	app := app.New()
	icon, _ := fyne.LoadResourceFromPath("./assets/sun.png")
	timeDt := time.Unix(int64(res.Dt), 0).String()

	window := app.NewWindow("Watch the weather")
	window.Resize(fyne.NewSize(700, 450))
	window.SetIcon(icon)

	//colorTitle := color.NRGBA{R: 10, G: 52, B: 64, A: 100}
	//title := canvas.NewText("Прогноз на день", colorTitle)

	//title := widget.NewLabel("Прогноз на день")
	footer := widget.NewLabel("Ссылки")

	/*buttonRefresh := widget.NewButton("Обновить", func() {
		fmt.Println(res)
	})*/

	/*header := container.NewHBox(
		title,
		buttonRefresh,
	)*/

	miniIcon, _ := fyne.LoadResourceFromURLString(fmt.Sprintf("http://openweathermap.org/img/wn/%s.png", res.Weather[0].Icon))

	leftBox := container.NewVBox(
		container.NewHBox(widget.NewLabel(fmt.Sprintf("Местоположение: %s", res.Name))),

		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Погодные условия: %s", res.Weather[0].Description)),
			widget.NewIcon(miniIcon),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Температура: %f C°", res.Main.Temp)),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Ветер: %f м/сек", res.Wind.Speed)),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Влажность: %d", res.Main.Humidity)),
		),
	)

	rightBox := container.NewVBox(
		widget.NewIcon(icon),
	)

	rightBox.Resize(fyne.NewSize(250, 250))

	wrapper := container.NewHBox(leftBox, rightBox)

	window.SetContent(container.NewVBox(
		//header,
		widget.NewLabel(timeDt),
		wrapper,
		footer,
	))

	window.ShowAndRun()
}
