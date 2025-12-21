// Package middleware adds all the requests interceptors
package middleware

import (
	"fmt"
	"net/http"

	"api/src/authorization"
	"api/src/responses"
)

// Logger is the middleware responsible for logging the requests
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("method:%s, uri:%s, host:%s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate authenticates the user requests
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticating...")

		if error := authorization.ValidateToken(r); error != nil {
			responses.Error(w, http.StatusUnauthorized, error)
			return
		}

		next(w, r)
	}
}
