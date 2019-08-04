package main

import (
	"database/sql"
	"os"

	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	// Create Data base connection
	db = ConnectDB()

	// Set info for CLI
	Info()

	// Set Commands & run Commands
	Commands()

	// run CLI
	err := app.Run(os.Args)
	logFatal(err)

}
