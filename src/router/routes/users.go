package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Handler:                controllers.UserCreate,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Handler:                controllers.UserGetAll,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Handler:                controllers.UserGet,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Handler:                controllers.UserUpdate,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Handler:                controllers.UserDelete,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/users/{id}/follow",
		Method:                 http.MethodPost,
		Handler:                controllers.UserFollow,
		RequiredAuthentication: true,
	},
}
