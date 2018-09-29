package routes

import (
	"comm/actions"
	"comm/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//InitRoutes the routes for the server
func InitRoutes(server *echo.Echo) {
	api := server.Group("/api", middleware.JWTWithConfig(middleware.JWTConfig{
		Claims: &utils.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}))

	// Community routes for api
	api.POST("/api/community/add", actions.CommunityAdd)
	api.PUT("/api/community/update", actions.CommunityUpdate)
	api.DELETE("/api/community/delete", actions.CommunityDelete)

	// Community routes public
	server.GET("/api/community/get/:id", actions.CommunityGetOne)
	server.GET("/api/community/get/all", actions.CommunityGetAll)
}
