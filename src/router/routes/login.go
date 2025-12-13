package routes

import (
	"net/http"

	"api/src/controllers"
)

var loginRoute = Route{
	URI:       "/login",
	Method:    http.MethodPost,
	Function:  controllers.Login,
	NeedsAuth: true,
}
