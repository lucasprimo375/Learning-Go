package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, e *http.Request) {
	w.Write([]byte("Creating user"))
}

func GetUsers(w http.ResponseWriter, e *http.Request) {
	w.Write([]byte("Getting all users"))
}

func GetUser(w http.ResponseWriter, e *http.Request) {
	w.Write([]byte("Getting user"))
}

func UpdateUser(w http.ResponseWriter, e *http.Request) {
	w.Write([]byte("Updating user"))
}

func DeleteUser(w http.ResponseWriter, e *http.Request) {
	w.Write([]byte("Deleting user"))
}
