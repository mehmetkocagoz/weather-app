package datascraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mehmetkocagz/database"
	"net/http"
)

type WeatherData struct {
	Name string `json:"name"`
	Main Main   `json:"main"`
}
type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
}

func ScrapeDataFromOpenWeatherAPI(city string) {
	apiKey := "4121b54183e3a277b8a37e48302d2ed2"
	URL := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey + "&units=metric"

	resp, err := http.Get(URL)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var weatherData WeatherData
	err = json.Unmarshal(jsonBytes, &weatherData)
	if err != nil {
		fmt.Println(err)
	}
	// Insert Data to Postgres

	db := database.Connect()
	// delete old data with same city name
	deleteSmt := "DELETE FROM weatherdata WHERE city = $1"
	db.Exec(deleteSmt, weatherData.Name)
	insertSmt := "INSERT INTO weatherdata (city, temp, feelslike, tempmin, tempmax, pressure, humidity) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	db.Exec(insertSmt, weatherData.Name, weatherData.Main.Temp, weatherData.Main.FeelsLike, weatherData.Main.TempMin, weatherData.Main.TempMax, weatherData.Main.Pressure, weatherData.Main.Humidity)
}
