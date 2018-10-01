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

	// exonsunity routes for api
	api.POST("/api/exonsunity/add", actions.exonsunityAdd)
	api.PUT("/api/exonsunity/update", actions.exonsunityUpdate)
	api.DELETE("/api/exonsunity/delete", actions.exonsunityDelete)

	// exonsunity routes public
	server.GET("/api/exonsunity/get/:id", actions.exonsunityGetOne)
	server.GET("/api/exonsunity/get/all", actions.exonsunityGetAll)

	// User routes public
	server.POST("/api/user/add", actions.UserAdd)
	server.POST("/api/user/login", actions.UserLogin)
}
