package routes

import (
	"devbook-api/src/controllers/users"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:               "/users",
		Method:            http.MethodGet,
		Function:          users.GetAll,
		HasAuthentication: false,
	},
	{
		URI:               "/users/{id}",
		Method:            http.MethodGet,
		Function:          users.GetOne,
		HasAuthentication: false,
	},
	{
		URI:               "/users",
		Method:            http.MethodPost,
		Function:          users.Create,
		HasAuthentication: false,
	},
	{
		URI:               "/users/{id}",
		Method:            http.MethodPut,
		Function:          users.Update,
		HasAuthentication: false,
	},
	{
		URI:               "/users/{id}",
		Method:            http.MethodDelete,
		Function:          users.Delete,
		HasAuthentication: false,
	},
}
