# Weather app
It's a small service that uses the OpenWeather api for offering temperature data for different cities

## Prerequisites
* go 1.20.7
* postgres (create a new database with name weatherapp)
* goose

## Instalation

* clone or download repository
* in the project root run `go mod tidy` and `go mod vendor`

## DB setup and migrations

First start the postgres server and run the DB migrations from db/migrations

`protocol://username:password@host:port/database`
`goose postgres postgres://postgres:111@localhost:5432/weatherapp up`


## How to use

To run the application, in the project root run

`go run build && ./weatherapp`

This will start a web server on port 8080, with 4 routes defined:
* `GET/http://localhost:8080/hello` - an example route 
* `GET/http://localhost:8080/weather` - this will get the temperature for new york city and huston texas 
* `POST/http://localhost:8080/store-weather` - this will be used to store the temperature from previous route in de db
* `GET/http://localhost:8080/stored-weather` - this will be used to get the stored data from de db
