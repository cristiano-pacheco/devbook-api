package routes

import "net/http"

// Route representation
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}
