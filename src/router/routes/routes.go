// Package routes contains all the possible routes of the api
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route defines the api route structure
type Route struct {
	URI       string
	Method    string
	Function  func(http.ResponseWriter, *http.Request)
	NeedsAuth bool
}

// Configure insert all routes inside the received router
func Configure(r *mux.Router) *mux.Router {
	routes := UsersRoute

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
