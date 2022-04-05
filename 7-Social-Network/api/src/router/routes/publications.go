package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationsRoutes = []Route{
	{
		URI:                    "/publications",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications",
		Method:                 http.MethodGet,
		Function:               controllers.GetPublications,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationID}",
		Method:                 http.MethodGet,
		Function:               controllers.GetPublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationID}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationID}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/publications",
		Method:                 http.MethodGet,
		Function:               controllers.GetPublicationsByUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationID}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePublication,
		RequiresAuthentication: true,
	},
}
