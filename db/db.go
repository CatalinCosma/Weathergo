package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {

	db, err := sqlx.Connect("postgres", "user=postgres password=111 dbname=weatherapp sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
