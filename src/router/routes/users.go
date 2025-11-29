package routes

import (
	"net/http"

	"api/src/controllers"
)

var UsersRoute = []Route{
	{
		URI:       "/users",
		Method:    http.MethodPost,
		Function:  controllers.CreateUser,
		NeedsAuth: false,
	},

	{
		URI:       "/users",
		Method:    http.MethodGet,
		Function:  controllers.GetUser,
		NeedsAuth: false,
	},

	{
		URI:       "/users/{userId}",
		Method:    http.MethodGet,
		Function:  controllers.GetUserByID,
		NeedsAuth: false,
	},

	{
		URI:       "/users",
		Method:    http.MethodDelete,
		Function:  controllers.DeleteUser,
		NeedsAuth: false,
	},

	{
		URI:       "/users",
		Method:    http.MethodPut,
		Function:  controllers.PutUser,
		NeedsAuth: false,
	},
}
