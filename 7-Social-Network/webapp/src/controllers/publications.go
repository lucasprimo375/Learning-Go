package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
	}

	url := fmt.Sprintf("%s/publications", config.APIURL)

	response, err := requests.MakeRequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(publication))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.ProcessErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
