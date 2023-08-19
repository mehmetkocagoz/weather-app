package model

import "mehmetkocagz/database"

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

func GetWeatherDataFromPostgres(city string) WeatherData {
	db := database.Connect()
	query := "SELECT * FROM weatherdata WHERE city = $1"
	row, err := db.Query(query, city)
	if err != nil {
		panic(err)
	}
	var weatherData WeatherData
	for row.Next() {
		err := row.Scan(&weatherData.Name, &weatherData.Main.Temp, &weatherData.Main.FeelsLike, &weatherData.Main.TempMin, &weatherData.Main.TempMax, &weatherData.Main.Pressure, &weatherData.Main.Humidity)
		if err != nil {
			panic(err)
		}
	}
	defer db.Close()
	return weatherData
}
