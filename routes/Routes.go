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
	server.GET("/api/community/get/:id", actions.CommunityGetOne)
	server.GET("/api/community/get/all", actions.CommunityGetAll)
	server.GET("api/community/get/all/user/:id", actions.CommunityGetAllUser)

	// User routes public
	server.GET("/api/user/:id", actions.UserGetOne)
	server.POST("/api/user/add", actions.UserAdd)
	server.POST("/api/user/login", actions.UserLogin)
}
