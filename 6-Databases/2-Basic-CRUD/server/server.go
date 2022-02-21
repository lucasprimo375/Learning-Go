package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:name`
	Email string `json:email`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requisitionBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Failed to read requsition body"))
		return
	}

	var user user

	error = json.Unmarshal(requisitionBody, &user)
	if error != nil {
		w.Write([]byte("Failed to convert user to struct"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("insert into users (name, email) values (?, ?)")
	if error != nil {
		fmt.Println(error)
		w.Write([]byte("Failed to create statement"))
		return
	}
	defer statement.Close()

	insertion, error := statement.Exec(user.Name, user.Email)
	if error != nil {
		w.Write([]byte("Failed to execute statement"))
		return
	}

	insertedId, error := insertion.LastInsertId()
	if error != nil {
		w.Write([]byte("Failed to get last inserted ID"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted ID %d successfully", insertedId)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	rows, error := db.Query("select * from users")
	if error != nil {
		w.Write([]byte("Failed to get users"))
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user

		error := rows.Scan(&user.ID, &user.Name, &user.Email)
		if error != nil {
			w.Write([]byte("Failed to scan user"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)

	error = json.NewEncoder(w).Encode(users)
	if error != nil {
		w.Write([]byte("Failed to convert users to json"))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {

}
