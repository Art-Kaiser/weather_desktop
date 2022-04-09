package UI

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math"
	"time"
	"weatherDesktop/api"
)

func renderPanelBottom(forecast *api.WeatherForecast) *fyne.Container {
	bottomBox := container.NewHBox()

	for i := 1; i < len(forecast.Daily); i++ {
		weatherBottomIcon := container.NewGridWrap(
			fyne.NewSize(150, 150),
			canvas.NewImageFromFile(
				fmt.Sprintf("./assets/weather/%s.png", forecast.Daily[i].Weather[0].Icon),
			),
		)
		_, month, day := time.Unix(int64(forecast.Daily[i].Dt), 0).Date()

		weatherGroup := widget.NewCard(
			fmt.Sprintf("Дата: %s %v", month, day),
			"",
			container.NewHBox(
				container.NewVBox(
					widget.NewLabel(fmt.Sprintf("Погодные условия: %s", forecast.Daily[i].Weather[0].Description)),
					widget.NewLabel(fmt.Sprintf(
						"Температура днём: %v°. Ночью: %v°",
						math.Round(forecast.Daily[i].Temp.Day),
						math.Round(forecast.Daily[i].Temp.Night),
					)),
					widget.NewLabel(fmt.Sprintf("Ветер: %v м/сек", math.Round(forecast.Daily[i].WindSpeed))),
					widget.NewLabel(fmt.Sprintf("Влажность: %d%%", forecast.Daily[i].Humidity)),
				),
				weatherBottomIcon,
			),
		)
		weatherWrapper := container.NewGridWrap(
			fyne.NewSize(550, 225),
			container.NewCenter(weatherGroup),
		)
		bottomBox.Add(weatherWrapper)
	}

	panelBottom := container.NewGridWrap(
		fyne.NewSize(750, 225),
		container.NewHScroll(bottomBox),
	)

	return panelBottom
}