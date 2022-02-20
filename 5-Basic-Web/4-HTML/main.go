package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	u := user{"Lucas", "lucas@email.com"}

	templates.ExecuteTemplate(w, "home.html", u)
}

type user struct {
	Name  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Listening on port 5000")

	log.Fatal(http.ListenAndServe(":5000", nil))
}
