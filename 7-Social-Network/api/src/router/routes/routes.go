package routes

import "net/http"

type Route struct {
	Uri                    string
	Method                 string
	function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}
