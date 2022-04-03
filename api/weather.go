package api

import (
	"encoding/json"
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetWeatherResult(apiKey string) *owm.CurrentWeatherData {
	weather, err := owm.NewCurrent("C", "ru", apiKey)
	if err != nil {
		log.Fatalln(err)
	}
	//test location
	weather.CurrentByName("Москва, RU")
	return weather
}

func GetWeathersResult(apiKey string, result *WeatherForecast) error {
	//test coordinates
	res, err := myClient.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=55.744458375950536&lon=37.62184820096254&lang=ru&exclude=minutely,hourly,current,alerts&appid=%s", apiKey))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(result)
}

func GetCoordinatesCity(apiKey, location string, result *CoordinatesCity) error {
	res, err := myClient.Get(fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?&appid=%s&q=%s", apiKey, location))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(result)
}
