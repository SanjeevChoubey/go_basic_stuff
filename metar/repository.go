package main

import (
	"database/sql"
)

type MetarRepository struct{}

func (m *MetarRepository) AddWeather(db *sql.DB, rawText string, id string) int {

	var airportID int
	err := db.QueryRow("Insert into weather_metar(airport_code,raw_text) values($1,$2) returning airport_id",
		&id, &rawText).Scan(&airportID)
	logFatal(err)

	return airportID
}
