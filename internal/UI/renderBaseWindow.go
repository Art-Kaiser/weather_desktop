package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	owm "github.com/briandowns/openweathermap"
	"weatherDesktop/api"
)

type Data struct {
	res      *owm.CurrentWeatherData
	forecast *api.WeatherForecast
}

type Component struct {
	input  *widget.Entry
	button *widget.Button
}

func renderBaseWindow(data Data, component Component) *fyne.Container {
	content := container.NewVBox(
		renderPanelTop(data.res),
		component.button,
		renderPanelBottom(data.forecast),
	)
	return content
}
