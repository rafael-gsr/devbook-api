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
		NeedsAuth: true,
	},

	{
		URI:       "/users/followers",
		Method:    http.MethodGet,
		Function:  controllers.FindFollowers,
		NeedsAuth: true,
	},

	{
		URI:       "/users/{userID}",
		Method:    http.MethodGet,
		Function:  controllers.GetUserByID,
		NeedsAuth: true,
	},

	{
		URI:       "/users/{userID}",
		Method:    http.MethodDelete,
		Function:  controllers.DeleteUser,
		NeedsAuth: true,
	},

	{
		URI:       "/users/{userID}",
		Method:    http.MethodPut,
		Function:  controllers.PutUser,
		NeedsAuth: true,
	},

	{
		URI:       "/users/{userID}/follow",
		Method:    http.MethodPost,
		Function:  controllers.FollowUser,
		NeedsAuth: true,
	},

	{
		URI:       "/users/{userID}/unfollow",
		Method:    http.MethodPost,
		Function:  controllers.UnfollowUser,
		NeedsAuth: true,
	},
}
