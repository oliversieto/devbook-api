package routes

import (
	"devbook-api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI               string
	Method            string
	Function          func(http.ResponseWriter, *http.Request)
	HasAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, authRoute)

	for _, route := range routes {
		if route.HasAuthentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.TokenValidator(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
