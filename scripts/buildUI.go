package scripts

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	owm "github.com/briandowns/openweathermap"
	"time"
	"weatherDesktop/api"
)

func InitUI(res *owm.CurrentWeatherData, forecast *api.WeatherForecast) {
	app := app.New()
	window := app.NewWindow("Watch the weather")
	window.Resize(fyne.NewSize(550, 450))
	iconApp, _ := fyne.LoadResourceFromPath("./assets/icon.png")
	//weatherIcon := container.NewGridWrap(fyne.NewSize(150, 150), canvas.NewImageFromFile(fmt.Sprintf("./assets/weather/%s.png", res.Weather[0].Icon)))
	resources, _ := fyne.LoadResourceFromURLString(fmt.Sprintf("http://openweathermap.org/img/wn/%s.png", res.Weather[0].Icon))
	weatherIcon := container.NewGridWrap(fyne.NewSize(150, 150), canvas.NewImageFromResource(resources))

	window.SetIcon(iconApp)
	timeDt := time.Unix(int64(res.Dt), 0).String()

	leftBox := container.NewVBox(
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Погодные условия: %s", res.Weather[0].Description)),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Температура: %f C°", res.Main.Temp)),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Ветер: %f м/сек", res.Wind.Speed)),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Влажность: %d%%", res.Main.Humidity)),
		),
	)

	top := widget.NewCard(
		fmt.Sprintf("Местоположение: %s", res.Name),
		timeDt,
		container.NewHBox(
			leftBox,
			weatherIcon,
		))

	bottomBox := container.NewHBox()

	for i := 0; i < len(forecast.Daily); i++ {
		fmt.Println(forecast.Daily[i].Weather[0].Icon)
		resource, _ := fyne.LoadResourceFromURLString(fmt.Sprintf("http://openweathermap.org/img/wn/%s.png", forecast.Daily[i].Weather[0].Icon))
		weatherGroup := widget.NewCard(
			fmt.Sprintf("Дата: %s", time.Unix(int64(forecast.Daily[i].Dt), 0)),
			"",
			container.NewVBox(
				widget.NewLabel(fmt.Sprintf("Погодные условия: %s", forecast.Daily[i].Weather[0].Description)),
				widget.NewLabel(fmt.Sprintf("Температура: %.2f °C", forecast.Daily[i].Temp.Day)),
				widget.NewLabel(fmt.Sprintf("Ветер: %f м/сек", forecast.Daily[i].WindSpeed)),
				widget.NewLabel(fmt.Sprintf("Влажность: %d%%", forecast.Daily[i].Humidity)),
			),
		)

		bottomBox.Add(weatherGroup)
		bottomBox.Add(widget.NewIcon(resource))
	}

	bottom := container.NewGridWrap(
		fyne.NewSize(600, 225),
		container.NewHScroll(bottomBox),
	)

	window.SetContent(container.NewVBox(
		top,
		bottom,
	))

	window.ShowAndRun()
}
