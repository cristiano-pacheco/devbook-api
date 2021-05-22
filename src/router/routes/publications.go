package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationRoutes = []Route{
	{
		URI:                    "/publications",
		Method:                 http.MethodPost,
		Handler:                controllers.PublicationCreate,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/publications",
		Method:                 http.MethodGet,
		Handler:                controllers.PublicationGetAll,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/publications/{id}",
		Method:                 http.MethodGet,
		Handler:                controllers.PublicationGet,
		RequiredAuthentication: true,
	},
	// {
	// 	URI:                    "/publications/{id}",
	// 	Method:                 http.MethodPut,
	// 	Handler:                controllers.PublicationUpdate,
	// 	RequiredAuthentication: true,
	// },
	// {
	// 	URI:                    "/publications/{id}",
	// 	Method:                 http.MethodDelete,
	// 	Handler:                controllers.PublicationDelete,
	// 	RequiredAuthentication: true,
	// },
}
