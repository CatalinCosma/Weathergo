package main

import (
	"github.com/CatalinCosma/weatherapp/app/handlers"
	"github.com/CatalinCosma/weatherapp/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	// cfg := config.LoadConfig()

	dbConn, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	r := gin.Default()
	handlers.SetupRoutes(r, dbConn)

	r.Run(":8080")
}
