package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

// Error function
func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}

func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)
	/*
		-db name
		-host
		-password
		-port
		-user
	*/
	db, err := sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
