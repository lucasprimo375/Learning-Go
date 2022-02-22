package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.GetUsers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{followedUserID}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{followedUserID}/unFollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnFollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/followers",
		Method:                 http.MethodGet,
		Function:               controllers.GetFollowers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/following",
		Method:                 http.MethodGet,
		Function:               controllers.GetFollowing,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/updatePassword",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		RequiresAuthentication: true,
	},
}
