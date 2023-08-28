package middleware

import (
	"errors"
	"net/http"
)

var UnauthorizedError = errors.New("Invalid Username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {

		}
	})
}
