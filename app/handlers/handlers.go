package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CatalinCosma/weatherapp/app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jmoiron/sqlx"
)

const apiKey = "YOUR_OPENWEATHER_APIKEY"

func SetupRoutes(r *gin.Engine, db *sqlx.DB) {
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
	r.GET("/weather", func(c *gin.Context) {
		houstonTemp, err := getTemperature("Houston")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data for Houston"})
			return
		}

		newYorkTemp, err := getTemperature("New York")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data for New York"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"houston_temperature": houstonTemp,
			"nyc_temperature":     newYorkTemp,
		})
	})

	r.POST("/store-weather", func(c *gin.Context) {
		var weatherData struct {
			HoustonTemperature float64 `json:"houston_temperature"`
			NYCTemperature     float64 `json:"nyc_temperature"`
		}
		if err := c.ShouldBindJSON(&weatherData); err != nil {
			c.JSON(400, gin.H{"error": "Bad Request"})
			return
		}

		_, err := db.Exec("INSERT INTO weather_data (houston_temperature, nyc_temperature) VALUES ($1, $2)",
			weatherData.HoustonTemperature, weatherData.NYCTemperature)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		c.JSON(201, gin.H{"message": "Data stored successfully"})
	})

	r.GET("/stored-weather", func(c *gin.Context) {
		var storedWeather models.WeatherData

		rows, err := db.Queryx("SELECT * FROM weather_data")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch stored weather data"})
			return
		}
		defer rows.Close()

		weatherData := make([]interface{}, 0)
		for rows.Next() {
			if err := rows.StructScan(&storedWeather); err != nil {
				c.JSON(500, gin.H{"error": "Failed to scan stored weather data"})
				return
			}
			weatherData = append(weatherData, storedWeather)
		}

		c.JSON(200, weatherData)
	})

}

func getTemperature(city string) (float64, error) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"appid": apiKey,
			"units": "metric",
		}).
		Get("https://api.openweathermap.org/data/2.5/weather?")

	if err != nil {
		return 0, err
	}

	if resp.StatusCode() != http.StatusOK {
		return 0, fmt.Errorf("API request failed with status code: %d", resp.StatusCode())
	}

	var weatherData models.WeatherResponse
	err = json.Unmarshal(resp.Body(), &weatherData)
	if err != nil {
		return 0, err
	}

	return weatherData.Main.Temp, nil
}
