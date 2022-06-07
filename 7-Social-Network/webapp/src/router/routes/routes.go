package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"webapp/src/middlewares"
)

// Route represents all routes in the web application
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// Configure sets all routes into the router
func Configure(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, mainRoute)

	for _, route := range routes {
		if route.RequiresAuthentication {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
