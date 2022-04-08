package UI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	owm "github.com/briandowns/openweathermap"
	"math"
	"time"
	"weatherDesktop/api"
)

func renderIconWeather(res *owm.CurrentWeatherData) *fyne.Container {
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

	return weatherIconWrapper
}

func renderPanelTop(res *owm.CurrentWeatherData, input *widget.Entry) *widget.Card {
	fmt.Println("input render top: ", input.Text)
	leftBox := container.NewVBox(
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Погодные условия: %s", res.Weather[0].Description)),
		),
		container.NewHBox(
			widget.NewLabel(fmt.Sprintf("Температура: %v°", math.Round(res.Main.Temp))),
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
			renderIconWeather(res),
		))
	return panelTop
}

func renderPanelBottom(forecast *api.WeatherForecast) *fyne.Container {
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
					widget.NewLabel(fmt.Sprintf(
						"Температура днём: %v°. Ночью: %v°",
						math.Round(forecast.Daily[i].Temp.Day),
						math.Round(forecast.Daily[i].Temp.Night),
					)),
					widget.NewLabel(fmt.Sprintf("Ветер: %v м/сек", math.Round(forecast.Daily[i].WindSpeed))),
					widget.NewLabel(fmt.Sprintf("Влажность: %d%%", forecast.Daily[i].Humidity)),
				),
				weatherBottomIcon,
			),
		)
		bottomBox.Add(weatherGroup)
	}

	panelBottom := container.NewGridWrap(
		fyne.NewSize(750, 225),
		container.NewHScroll(bottomBox),
	)

	return panelBottom
}

func renderBaseWindow(data Data, component Component) *fyne.Container {
	content := container.NewVBox(
		renderPanelTop(data.res, component.input),
		component.button,
		renderPanelBottom(data.forecast),
	)
	return content
}

type Data struct {
	res      *owm.CurrentWeatherData
	forecast *api.WeatherForecast
}

type Component struct {
	input  *widget.Entry
	button *widget.Button
}

func Init(res *owm.CurrentWeatherData, forecast *api.WeatherForecast, input *widget.Entry) {
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

	data := Data{res: res, forecast: forecast}
	component := Component{input: input, button: buttonChoiceCity}

	window.SetContent(
		renderBaseWindow(data, component),
	)

	dialogChoiceCity.SetOnClosed(func() {
		if len(input.Text) != 0 {
			window.SetContent(
				renderBaseWindow(data, component),
			)
		}

		//тест в случае ошибки
		var errTest bool = true

		if errTest {
			dialog.ShowInformation(
				"Внимание!",
				"Произошла ошибка при получении данных о погоде. \n Вероятно это может быть связано с некорректным заполнением поля\n  или сторонней ошибкой",
				window,
			)
		}
	})

	/*go func() {
		for range time.Tick(time.Minute * 10) {
			fmt.Println("test: ", time.Minute*10)
		}
	}()*/

	window.Show()
	app.Run()
}
