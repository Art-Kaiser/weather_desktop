package main

import (
	"os"
	"weatherDesktop/configs"
	"weatherDesktop/internal/UI"
)

func main() {
	configs.InitConfig()

	if _, exists := os.LookupEnv("API_WEATHER_KEY"); exists {
		UI.Init()
	}
}
