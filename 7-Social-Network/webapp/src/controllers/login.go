package controllers

import "net/http"

// LoadLoginPage renders login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login page"))
}
