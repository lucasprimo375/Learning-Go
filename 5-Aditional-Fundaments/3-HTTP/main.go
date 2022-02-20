package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Loading users page"))
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
