package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var mainRoute = Route{
	URI:                    "/home",
	Method:                 http.MethodGet,
	Function:               controllers.LoadMainPage,
	RequiresAuthentication: true,
}
