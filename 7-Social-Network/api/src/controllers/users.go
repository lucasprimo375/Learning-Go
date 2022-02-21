package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User

	error = json.Unmarshal(requestBody, &user)
	if error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repository := repository.NewUsersRepository(db)

	userID, error := repository.Create(user)
	if error != nil {
		log.Fatal(error)
	}

	w.Write([]byte(fmt.Sprintf("Inserted ID: %d", userID)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
