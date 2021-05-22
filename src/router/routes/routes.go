package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/PacktPublishing/Echo-Essentials/chapter2/middlewares"
	"github.com/gorilla/mux"
)

// Route representation
type Route struct {
	URI                    string
	Method                 string
	Handler                func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, authRoute)

	for _, route := range routes {
		if route.RequiredAuthentication {
			r.HandleFunc(route.URI, middlewares.Authenticate(route.Handler)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
		}
	}

	return r
}
