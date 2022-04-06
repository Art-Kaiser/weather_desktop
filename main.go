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

	if apiKey, exists := os.LookupEnv("API_WEATHER_KEY"); exists {
		forecast := new(api.WeatherForecast)
		city := new(api.CoordinatesCity)
		//test api

		api.GetWeathersResult(apiKey, forecast)
		api.GetCoordinatesCity(apiKey, "Moskow", city)

		inputCity := widget.NewEntry()
		res := api.GetWeatherResult(apiKey, inputCity)
		UI.Init(res, forecast, inputCity)

		/*go func() {
			for range time.Tick(time.Minute) {
				resTest := api.GetWeatherResult(apiKey, inputCity)
				scripts.InitUI(resTest, forecast, inputCity)
			}
		}()*/
	}
}
