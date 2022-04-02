package api

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
)

func GetWeatherResult(apiKey string) *owm.CurrentWeatherData {
	//каждую минуту запрос к апи 60 запросов в час/1440 в сутки/44640 в месяц 54.52167036084512,36.27440741060984
	weather, err := owm.NewCurrent("C", "ru", apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	weather.CurrentByName("Krasnodar, RU")
	fmt.Println(weather.Weather[0].Description, weather.Name, weather.Main)
	return weather
}
