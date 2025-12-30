// Package routes contains all the possible routes of the api
package routes

import (
	"net/http"

	"api/src/middleware"

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
	routes = append(routes, loginRoute)
	routes = append(routes, publicationRoutes...)

	for _, route := range routes {
		loggedHandlerFunc := middleware.Logger(route.Function)

		if route.NeedsAuth {
			r.HandleFunc(route.URI, middleware.Authenticate(loggedHandlerFunc)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, loggedHandlerFunc).Methods(route.Method)
		}
	}

	return r
}
