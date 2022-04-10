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
	"weatherDesktop/pkg/formatRu"
)

func renderPanelBottom(forecast *api.WeatherForecast) *fyne.Container {
	bottomBox := container.NewHBox()
	for i := 1; i < len(forecast.Daily); i++ {
		_, month, day := time.Unix(int64(forecast.Daily[i].Dt), 0).Date()
		weekday := time.Unix(int64(forecast.Daily[i].Dt), 0).Weekday()

		weatherBottomIcon := container.NewGridWrap(
			fyne.NewSize(150, 150),
			canvas.NewImageFromFile(
				fmt.Sprintf("./assets/weather/%s.png", forecast.Daily[i].Weather[0].Icon),
			),
		)

		monthRu := formatRu.Format(month.String(), true)
		weekdayRu := formatRu.Format(weekday.String(), true)

		weatherConditions := widget.NewLabel(fmt.Sprintf("Погодные условия: %s", forecast.Daily[i].Weather[0].Description))
		weatherConditions.Wrapping = fyne.TextWrapWord

		weatherGroup := widget.NewCard(
			fmt.Sprintf("Дата: %s %v %s", weekdayRu, day, monthRu),
			"",
			container.NewHBox(
				container.NewVBox(
					weatherConditions,
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
			fyne.NewSize(445, 235),
			container.NewGridWithColumns(1, weatherGroup),
		)
		bottomBox.Add(weatherWrapper)
	}

	panelBottom := container.NewGridWrap(
		fyne.NewSize(750, 235),
		container.NewHScroll(bottomBox),
	)

	return panelBottom
}
