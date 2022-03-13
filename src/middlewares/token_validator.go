package middlewares

import (
	"devbook-api/src/authentication"
	"devbook-api/src/responses"
	"net/http"
)

func TokenValidator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.TokenValidator(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
