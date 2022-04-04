package main

import (
	"os"
	"weatherDesktop/api"
	"weatherDesktop/configs"
	"weatherDesktop/scripts"
)

func main() {
	configs.InitConfig()

	if apiKey, exists := os.LookupEnv("API_WEATHER_KEY"); exists {
		forecast := new(api.WeatherForecast)
		city := new(api.CoordinatesCity)
		//test api
		res := api.GetWeatherResult(apiKey)
		api.GetWeathersResult(apiKey, forecast)
		api.GetCoordinatesCity(apiKey, "Moskow", city)
		scripts.InitUI(res, forecast)
	}
}
