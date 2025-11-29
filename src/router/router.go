// Package router handles the route of the application
package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// NewRouter returns a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	configuredRouter := routes.Configure(router)

	return configuredRouter
}
