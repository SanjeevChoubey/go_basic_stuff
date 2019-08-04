package main

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

var db *sql.DB

// func logFatal(err error) {
// 	if err != nil {
// 		log.Fatalln(err)

// 	}

// }

// This function will be used to create postgre connection
func CreateConnection() (*sql.DB, error) {
	//pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	pgURL, err := pq.ParseURL("postgres://djtebaig:swjeqkyjJnB_QIKOoIxlI2dYzrMhVi03@satao.db.elephantsql.com:5432/djtebaig")
	//logFatal(err)
	if err != nil {
		log.Fatalln(err)
		return nil, err

	}

	db, err := sql.Open("postgres", pgURL)
	//logFatal(err)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	logFatal(err)
	if err != nil {
		log.Fatalln(err)
		return nil, err

	}

	return db, nil

}
