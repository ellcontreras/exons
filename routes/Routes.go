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
	api.POST("/community/add", actions.CommunityAdd)
	api.PUT("/community/update", actions.CommunityUpdate)
	api.DELETE("/community/delete", actions.CommunityDelete)

	// exonsunity routes public
	server.GET("/api/exonsunity/get/:id", actions.CommunityGetOne)
	server.GET("/api/exonsunity/get/all", actions.CommunityGetAll)

	// User routes public
	server.POST("/api/user/add", actions.UserAdd)
	server.POST("/api/user/login", actions.UserLogin)
}
