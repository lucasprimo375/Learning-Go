package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                    "/create-user",
		Method:                 http.MethodGet,
		Function:               controllers.LoadCreateUserPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/search-users",
		Method:                 http.MethodGet,
		Function:               controllers.LoadSearchUsersPage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUserProfile,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/unFollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnFollow,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/Follow",
		Method:                 http.MethodPost,
		Function:               controllers.Follow,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/profile",
		Method:                 http.MethodGet,
		Function:               controllers.LoadLoggedUserProfile,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/edit-profile",
		Method:                 http.MethodGet,
		Function:               controllers.LoadEditProfilePage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/edit-profile",
		Method:                 http.MethodPut,
		Function:               controllers.EditProfile,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/update-password",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUpdatePasswordPage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/update-password",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/delete-profile",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteProfile,
		RequiresAuthentication: true,
	},
}
