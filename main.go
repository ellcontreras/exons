package main

import (
	"comm/routes"

	"github.com/labstack/echo"
)

var (
	//Server is the server of the app
	Server *echo.Echo
)

func main() {
	Server = echo.New()

	//Init the routes
	routes.InitRoutes(Server)

	//Make migrations
	//db := DB.Open()
	//db.AutoMigrate(&models.Product{})

	Server.Logger.Fatal(Server.Start(":8080"))
}

