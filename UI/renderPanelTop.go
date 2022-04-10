package UI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	owm "github.com/briandowns/openweathermap"
	"math"
	"time"
	"weatherDesktop/pkg/formatRu"
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
		fyne.NewSize(356, 150),
		contentCenter,
	)

	return weatherIconWrapper
}

func renderPanelTop(res *owm.CurrentWeatherData) *widget.Card {
	_, month, day := time.Unix(int64(res.Dt), 0).Date()
	updateTime := time.Unix(int64(res.Dt), 0).Format("15:04")
	monthRu := formatRu.Format(month.String(), false)

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

	panelTop := widget.NewCard(
		fmt.Sprintf("%s. Обновлено в: %s по МСК", res.Name, updateTime),
		fmt.Sprintf("Дата: %s %v", monthRu, day),
		container.NewHBox(
			leftBox,
			renderIconWeather(res),
		))
	return panelTop
}
