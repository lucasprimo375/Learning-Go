package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql v1.6.0" // Driver
	"api/src/config"
)

func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.DatabaseConnectionString)
	if error != nil {
		return nil, error
	}

	error = db.Ping()
	if error != nil {
		db.Close()
		return nil, error
	}
	
	return db, nil
}
