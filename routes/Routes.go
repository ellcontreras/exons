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

<<<<<<< HEAD
	// Community routes for api
	api.POST("/community/add", actions.CommunityAdd)
	api.PUT("/community/update", actions.CommunityUpdate)
	api.DELETE("/community/delete", actions.CommunityDelete)
=======
	// exonsunity routes for api
	api.POST("/api/exonsunity/add", actions.exonsunityAdd)
	api.PUT("/api/exonsunity/update", actions.exonsunityUpdate)
	api.DELETE("/api/exonsunity/delete", actions.exonsunityDelete)
>>>>>>> 7ffc66964085de8bd601e8428984ada54100d2de

	// exonsunity routes public
	server.GET("/api/exonsunity/get/:id", actions.exonsunityGetOne)
	server.GET("/api/exonsunity/get/all", actions.exonsunityGetAll)

	// User routes public
	server.POST("/api/user/add", actions.UserAdd)
	server.POST("/api/user/login", actions.UserLogin)
}
