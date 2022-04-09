package main

import (
	"fyne.io/fyne/v2/widget"
	"os"
	"weatherDesktop/UI"
	"weatherDesktop/configs"
)

func main() {
	configs.InitConfig()

	if _, exists := os.LookupEnv("API_WEATHER_KEY"); exists {
		inputCity := widget.NewEntry()
		UI.Init(inputCity)
	}
}
