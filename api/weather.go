package api

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2/widget"
	owm "github.com/briandowns/openweathermap"
	"log"
	"net/http"
	"os"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

//test

func GetWeatherResult(inputCity *widget.Entry) *owm.CurrentWeatherData {
	apiKey, _ := os.LookupEnv("API_WEATHER_KEY")
	weather, err := owm.NewCurrent("C", "ru", apiKey)
	if err != nil {
		log.Fatalln(err)
	}
	//test location inputCity

	weather.CurrentByName(fmt.Sprintf("%s, RU", inputCity.Text))
	fmt.Println("weather: ", weather)
	return weather
}

func GetWeathersResult(result *WeatherForecast) error {
	apiKey, _ := os.LookupEnv("API_WEATHER_KEY")
	//test coordinates
	res, err := myClient.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?lat=55.744458375950536&lon=37.62184820096254&lang=ru&units=metric&exclude=minutely,hourly,current,alerts&appid=%s", apiKey))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(result)
}

func GetCoordinatesCity(location string, result *CoordinatesCity) error {
	apiKey, _ := os.LookupEnv("API_WEATHER_KEY")
	res, err := myClient.Get(fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?&appid=%s&q=%s", apiKey, location))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(result)
}
