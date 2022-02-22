package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
	}

	var user models.User

	error = json.Unmarshal(requestBody, &user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)

	userOnDatabase, error := repository.GetByEmail(user.Email)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	error = security.CheckPassword(user.Password, userOnDatabase.Password)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, _ := authentication.CreateToken(userOnDatabase.ID)

	w.Write([]byte(token))
}
