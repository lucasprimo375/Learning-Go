package router

import "github.com/gorilla/mux"

// Generate returns router with all routes configured
func Generate() *mux.Router {
	return mux.NewRouter()
}
