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
		fyne.NewSize(750, 225),
		container.NewHScroll(bottomBox),
	)

	return panelBottom
}

func Init(res *owm.CurrentWeatherData, forecast *api.WeatherForecast, input *widget.Entry) {
	app := app.New()
	window := app.NewWindow("Watch the weather")

	window.Resize(fyne.NewSize(550, 450))
	window.SetFixedSize(true)
	window.SetMaster()

	iconApp, err := fyne.LoadResourceFromPath("./assets/icon.png")
	if err != nil {
		window.SetIcon(iconApp)
	}

	dialogChoiceCity := dialog.NewCustomConfirm(
		"Выбор города",
		"Сохранить",
		"Отклонить",
		input,
		func(b bool) {
			if b {
				fmt.Println("input:", input.Text)
				window.SetContent(renderPanelTop(res, input))
			}
		},
		window,
	)

	dialogChoiceCity.Resize(fyne.NewSize(265, 115))

	buttonChoiceCity := widget.NewButton("Поменять город", func() {
		dialogChoiceCity.Show()
	})

	content := container.NewVBox(
		renderPanelTop(res, input),
		buttonChoiceCity,
		renderPanelBottom(forecast),
	)

	window.SetContent(content)

	dialogChoiceCity.SetOnClosed(func() {
		//	fmt.Println(input.Text)
		//	window.SetContent(content)
	})

	go func() {
		for range time.Tick(time.Minute * 10) {
			fmt.Println("test: ", time.Minute*10)
		}
	}()

	/*go func() {
		for range time.Tick(time.Minute) {
			//для запросов по текущему дню
			//formatted := time.Now().Format("It`s: 3:04:05")
			fmt.Println("inputCity.Text: ", input.Text)
			resTest := api.GetWeatherResult("", input)

			window.SetContent(container.NewVBox(
				//panelTop,
				renderPanelTop(resTest, app, input),
				panelBottom,
			))
		}
	}()*/
	/*go func() {
		for range time.Tick(time.Second) {
			fmt.Println("test: ", time.Second)
		}
	}()*/

	window.Show()
	app.Run()
}
