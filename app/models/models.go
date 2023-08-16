package models

type WeatherData struct {
	ID                 int     `db:"id"`
	HoustonTemperature float64 `db:"houston_temperature"`
	NYCTemperature     float64 `db:"nyc_temperature"`
	CreatedAt          string  `db:"created_at"`
}

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}
