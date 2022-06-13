package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/utils"
)

// LoadLoginPage renders login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadCreateUserPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "create.html", nil)
}

func LoadMainPage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)

	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)

	fmt.Println(response, err)

	utils.ExecuteTemplate(w, "home.html", nil)
}
