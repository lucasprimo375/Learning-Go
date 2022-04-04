package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	err = json.Unmarshal(requestBody, &publication)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = userID

	err = publication.Prepare()
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewPublicationsRepository(db)

	publication.ID, err = repository.Create(publication)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

func GetPublications(w http.ResponseWriter, r *http.Request) {}

func GetPublication(w http.ResponseWriter, r *http.Request) {}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {}

func DeletePublication(w http.ResponseWriter, r *http.Request) {}
