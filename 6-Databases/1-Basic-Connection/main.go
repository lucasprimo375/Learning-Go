package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectionString := ""

	db, error := sql.Open("mysql", connectionString)

	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()

	error = db.Ping()
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Connection open")

	rows, error := db.Query("select * from users")
	if error != nil {
		log.Fatal(error)
	}
	defer rows.Close()

	fmt.Println(rows)
}
