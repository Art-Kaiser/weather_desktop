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
		city := new(api.CoordinatesCity)
		//test api

		api.GetWeathersResult(forecast)
		api.GetCoordinatesCity("Moskow", city)

		inputCity := widget.NewEntry()
		res := api.GetWeatherResult(inputCity)

		UI.Init(res, forecast, inputCity)
	}
}
