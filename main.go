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
		res := api.GetWeatherResult(apiKey)
		scripts.InitUI(res)
	}
}
