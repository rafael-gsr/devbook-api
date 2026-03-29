package routes

import (
	"net/http"

	"api/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:       "/posts",
		Method:    http.MethodGet,
		Function:  controllers.GetPosts,
		NeedsAuth: true,
	},

	{
		URI:       "/posts/like/{postID}",
		Method:    http.MethodPut,
		Function:  controllers.LikePost,
		NeedsAuth: true,
	},

	{
		URI:       "/posts/dislike/{postID}",
		Method:    http.MethodPut,
		Function:  controllers.DislikePost,
		NeedsAuth: true,
	},

	{
		URI:       "/posts/user/{userID}",
		Method:    http.MethodGet,
		Function:  controllers.GetPostsByUserID,
		NeedsAuth: true,
	},

	{
		URI:       "/posts/{postID}",
		Method:    http.MethodGet,
		Function:  controllers.GetPostByID,
		NeedsAuth: true,
	},

	{
		URI:       "/posts",
		Method:    http.MethodPost,
		Function:  controllers.CreatePost,
		NeedsAuth: true,
	},

	{
		URI:       "/posts/{postID}",
		Method:    http.MethodPatch,
		Function:  controllers.UpdatePost,
		NeedsAuth: true,
	},

	{
		URI:       "/posts/{postID}",
		Method:    http.MethodDelete,
		Function:  controllers.DeletePost,
		NeedsAuth: true,
	},
}
