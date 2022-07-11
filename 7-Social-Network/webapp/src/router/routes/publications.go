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
	{
		URI:                    "/publications/{publicationID}/dislike",
		Method:                 http.MethodPost,
		Function:               controllers.DislikePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationID}/edit",
		Method:                 http.MethodGet,
		Function:               controllers.LoadEditPublicationPage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationID}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePublication,
		RequiresAuthentication: true,
	},
}
