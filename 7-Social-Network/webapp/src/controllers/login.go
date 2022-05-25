package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// LoadLoginPage renders login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}
