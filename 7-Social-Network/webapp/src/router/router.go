package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns router with all routes configured
func Generate() *mux.Router {
	router := mux.NewRouter()

	return routes.Configure(router)
}
