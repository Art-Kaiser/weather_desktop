package api

import (
	"encoding/json"
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
	"net/http"
	"os"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

//test

func GetWeatherResult(inputCity string) *owm.CurrentWeatherData {
	apiKey, _ := os.LookupEnv("API_WEATHER_KEY")
	weather, err := owm.NewCurrent("C", "ru", apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	weather.CurrentByName(fmt.Sprintf("%s, RU", inputCity))
	return weather
}

func GetWeathersResult(result *WeatherForecast, coordinates owm.Coordinates) error {
	apiKey, _ := os.LookupEnv("API_WEATHER_KEY")

	lat := coordinates.Latitude
	lon := coordinates.Longitude

	res, err := myClient.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=%f&lon=%f&lang=ru&units=metric&exclude=minutely,hourly,current,alerts&appid=%s", lat, lon, apiKey))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(result)
}
