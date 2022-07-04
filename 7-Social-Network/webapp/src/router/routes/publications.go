package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var publicationRoutes = []Route{
	{
		URI:                    "/publications",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationID}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePublication,
		RequiresAuthentication: true,
	},
}
