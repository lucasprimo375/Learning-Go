package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIError struct {
	Error string `json:error`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ProcessErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var err APIError

	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
