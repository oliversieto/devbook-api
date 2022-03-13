package routes

import (
	"net/http"

	"devbook-api/src/controllers/auth"
)

var authRoute = Route{
	URI:               "/auth",
	Method:            http.MethodPost,
	Function:          auth.Authentication,
	HasAuthentication: false,
}
