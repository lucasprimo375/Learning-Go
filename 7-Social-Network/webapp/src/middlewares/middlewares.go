package middlewares

import (
	"log"
	"net/http"

	"webapp/src/cookies"
)

func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		nextFunction(w, r)
	}
}

func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := cookies.Read(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		nextFunction(w, r)
	}
}
