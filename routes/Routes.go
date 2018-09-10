package routes

import (
	"comm/actions"

	"github.com/labstack/echo"
)

//InitRoutes the routes for the server
func InitRoutes(server *echo.Echo) {
	//User routes
	server.GET("/api/communnity/get/all", actions.CommunityGetOne)
	server.POST("/api/community/get/:id", actions.CommunityGetAll)
	server.PUT("/api/community/add", actions.CommunityAdd)
	server.DELETE("/api/community/update", actions.CommunityUpdate)
}

