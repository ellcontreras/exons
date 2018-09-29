package actions

import (
	"net/http"
	"github.com/labstack/echo"

	"comm/utils"
	"comm/models"
)

func AddUser(ctx echo.Context) error {
	user := models.User{}

	user.BindWithContext(ctx)

	Connect()

	err = collection.Insert(user)
	utils.CheckErr(err)
	
	Disconect()

	return ctx.JSON(http.StatusCreated, user)
}
