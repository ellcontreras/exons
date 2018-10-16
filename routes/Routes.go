package routes

import (
	"exons/actions"
	"exons/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//InitRoutes the routes for the server
func InitRoutes(server *echo.Echo) {
	api := server.Group("/api", middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &utils.JWTCustomClaims{},
		SigningKey: []byte("secret"),
	}))

	// Community routes for api
	api.POST("/Community/add", actions.CommunityAdd)
	api.PUT("/Community/update", actions.CommunityUpdate)
	api.DELETE("/Community/delete", actions.CommunityDelete)

	// Community routes public
	server.GET("/api/Community/get/:id", actions.CommunityGetOne)
	server.GET("/api/Community/get/all", actions.CommunityGetAll)

	// User routes public
	server.POST("/api/user/add", actions.UserAdd)
	server.POST("/api/user/login", actions.UserLogin)
}
