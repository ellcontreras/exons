package main

import (
	"github.com/labstack/echo/middleware"
	"comm/routes"

	"github.com/labstack/echo"

)

var (
	//Server is the server of the app
	Server *echo.Echo
)

func main() {
	Server = echo.New()

	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	Server.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: `{"time":"${time_rfc3339_nano}","host":"${host}","method":"${method}","uri":"${uri}","status":${status},`+
			`"bytes_in":${bytes_in},` + `"bytes_out":${bytes_out}}` + "\n",
		},
	))

	//Init the routes
	routes.InitRoutes(Server)

	//Make migrations
	//db := DB.Open()
	//db.AutoMigrate(&models.Product{})

	Server.Logger.Fatal(Server.Start(":8080"))
}

