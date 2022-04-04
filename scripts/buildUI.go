package scripts

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	owm "github.com/briandowns/openweathermap"
	"math"
	"time"
	"weatherDesktop/api"
)

func InitUI(res *owm.CurrentWeatherData, forecast *api.WeatherForecast) {
	app := app.New()
	window := app.NewWindow("Watch the weather")
	window.Resize(fyne.NewSize(550, 450))
	window.SetFixedSize(true)

	iconApp, _ := fyne.LoadResourceFromPath("./assets/icon.png")

	//text := canvas.NewText("Overlay", color.Black)

	weatherIcon := container.NewGridWrap(
		fyne.NewSize(160, 160),
		canvas.NewImageFromFile(
			fmt.Sprintf("./assets/weather/%s.png", res.Weather[0].Icon),
		),
	)

	contentCenter := container.New(layout.NewCenterLayout(), weatherIcon)

	weatherIconWrapper := container.NewGridWrap(
		fyne.NewSize(375, 150),
		contentCenter,
	)

	window.SetIcon(iconApp)

	leftBox := container.NewVBox(
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Погодные условия: %s", res.Weather[0].Description)),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Температура: %v C°", math.Round(res.Main.Temp))),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Ветер: %v м/сек", math.Round(res.Wind.Speed))),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Влажность: %d%%", res.Main.Humidity)),
		),
	)

	_, monthRes, dayRes := time.Unix(int64(res.Dt), 0).Date()

	panelTop := widget.NewCard(
		fmt.Sprintf("Местоположение: %s", res.Name),
		fmt.Sprintf("Дата: %s %v", monthRes, dayRes),
		container.NewHBox(
			leftBox,
			weatherIconWrapper,
		))

	bottomBox := container.NewHBox()

	for i := 1; i < len(forecast.Daily); i++ {
		weatherBottomIcon := container.NewGridWrap(
			fyne.NewSize(150, 150),
			canvas.NewImageFromFile(
				fmt.Sprintf("./assets/weather/%s.png", forecast.Daily[i].Weather[0].Icon),
			),
		)
		_, month, day := time.Unix(int64(forecast.Daily[i].Dt), 0).Date()

		weatherGroup := widget.NewCard(
			fmt.Sprintf("Дата: %s %v", month, day),
			"",
			container.NewHBox(
				container.NewVBox(
					widget.NewLabel(fmt.Sprintf("Погодные условия: %s", forecast.Daily[i].Weather[0].Description)),
					widget.NewLabel(fmt.Sprintf("Температура: %v °C", math.Round(forecast.Daily[i].Temp.Day))),
					widget.NewLabel(fmt.Sprintf("Ветер: %v м/сек", math.Round(forecast.Daily[i].WindSpeed))),
					widget.NewLabel(fmt.Sprintf("Влажность: %d%%", forecast.Daily[i].Humidity)),
				),
				weatherBottomIcon,
			),
		)

		bottomBox.Add(weatherGroup)
	}

	panelBottom := container.NewGridWrap(
		fyne.NewSize(600, 225),
		container.NewHScroll(bottomBox),
	)

	window.SetContent(container.NewVBox(
		panelTop,
		panelBottom,
	))

	window.ShowAndRun()
}
