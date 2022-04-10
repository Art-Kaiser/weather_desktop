package main

import (
	"os"
	"weatherDesktop/UI"
	"weatherDesktop/configs"
)

func main() {
	configs.InitConfig()

	if _, exists := os.LookupEnv("API_WEATHER_KEY"); exists {
		UI.Init()
	}
}
