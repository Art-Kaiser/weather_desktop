package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"time"
	"weatherDesktop/api"
)

func Init(input *widget.Entry) {
	weathersResult := new(api.WeatherForecast)
	res := api.GetWeatherResult(input.Text)
	api.GetWeathersResult(weathersResult, res.GeoPos)

	app := app.New()
	window := app.NewWindow("Следите за погодой")

	window.Resize(fyne.NewSize(550, 450))
	window.SetFixedSize(true)
	window.SetMaster()

	iconApp, err := fyne.LoadResourceFromPath("./assets/icon.png")
	if err != nil {
		window.SetIcon(iconApp)
	}

	dialogChoiceCity := dialog.NewCustom(
		"Выбор города",
		"Сохранить",
		input,
		window,
	)
	dialogChoiceCity.Resize(fyne.NewSize(265, 115))

	buttonChoiceCity := widget.NewButton("Поменять город", func() {
		dialogChoiceCity.Show()
	})

	data := Data{res: res, forecast: weathersResult}
	component := Component{input: input, button: buttonChoiceCity}

	window.SetContent(
		renderBaseWindow(data, component),
	)

	dialogChoiceCity.SetOnClosed(func() {
		if len(input.Text) != 0 {
			resUpdate := api.GetWeatherResult(input.Text)
			api.GetWeathersResult(weathersResult, resUpdate.GeoPos)

			dataUpdate := Data{res: resUpdate, forecast: weathersResult}

			window.SetContent(
				renderBaseWindow(dataUpdate, component),
			)

			/*if err != nil {
				dialog.ShowInformation(
					"Внимание!",
					"Произошла ошибка при получении данных о погоде. "+
						"\n Вероятно это может быть связано с некорректным заполнением поля"+
						"\n  или сторонней ошибкой",
					window,
				)
			}*/
		}
	})

	go func() {
		count := 0
		for range time.Tick(time.Minute) {
			count++
			res = api.GetWeatherResult(input.Text)
			data = Data{res: res, forecast: weathersResult}

			window.SetContent(
				renderBaseWindow(data, component),
			)

			if count%10 == 0 {
				api.GetWeathersResult(weathersResult, res.GeoPos)
				data = Data{res: res, forecast: weathersResult}

				window.SetContent(
					renderBaseWindow(data, component),
				)
				count = 0
			}
		}
	}()

	window.Show()
	app.Run()
}
