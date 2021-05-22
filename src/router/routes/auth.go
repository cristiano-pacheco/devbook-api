package routes

import (
	"api/src/controllers"
	"net/http"
)

var authRoute = Route{
	URI:                    "/auth/token",
	Method:                 http.MethodPost,
	Handler:                controllers.IssueToken,
	RequiredAuthentication: false,
}
