package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // connection driver for MySQL
)

// open connection to the database
func Connect() (*sql.DB, error) {
	connectionString := ""

	db, error := sql.Open("mysql", connectionString)
	if error != nil {
		return nil, error
	}

	error = db.Ping()
	if error != nil {
		return nil, error
	}

	return db, nil
}
