package main

import (
	"fyne.io/fyne/v2/widget"
	"os"
	"weatherDesktop/UI"
	"weatherDesktop/api"
	"weatherDesktop/configs"
)

func main() {
	configs.InitConfig()

	if _, exists := os.LookupEnv("API_WEATHER_KEY"); exists {
		forecast := new(api.WeatherForecast)
		api.GetWeathersResult(forecast)

		inputCity := widget.NewEntry()
		//res := api.GetWeatherResult(inputCity)

		UI.Init(forecast, inputCity)
	}
}
