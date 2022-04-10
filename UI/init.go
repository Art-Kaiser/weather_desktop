package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
	"time"
	"weatherDesktop/api"
)

func Init(input *widget.Entry) {
	weathersResult := new(api.WeatherForecast)
	res := api.GetWeatherResult(input.Text)
	api.GetWeathersResult(weathersResult, res.GeoPos)

	app := app.New()
	window := app.NewWindow("Следите за погодой")

	if _, exists := os.LookupEnv("API_WEATHER_KEY"); !exists {
		renderPopupInfo("Добавте в config файл .env с вашим API ключом.\n API_WEATHER_KEY=Ваш ключ", window)
	}

	/*renderPopupInfo("Произошла ошибка при получении данных о погоде. "+
	"\n Вероятно это может быть связано с некорректным заполнением поля"+
	"\n  или сторонней ошибкой", window)*/

	window.Resize(fyne.NewSize(550, 470))
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
			res = api.GetWeatherResult(input.Text)

			if len(res.Weather) == 0 {
				renderPopupInfo("Произошла ошибка при получении данных о погоде. "+
					"\n Вероятно это может быть связано с некорректным заполнением поля"+
					"\n  или сторонней ошибкой",
					window)
				return
			}

			api.GetWeathersResult(weathersResult, res.GeoPos)

			dataUpdate := Data{res: res, forecast: weathersResult}

			window.SetContent(
				renderBaseWindow(dataUpdate, component),
			)
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
