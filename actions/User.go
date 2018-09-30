package actions

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"comm/models"
	"comm/utils"
)

// UserAdd ...
func UserAdd(ctx echo.Context) error {
	user := models.User{}
	userEmail := models.User{}
	userUsername := models.User{}

	user.BindWithContext(ctx)

	Connect()

	collectionUsers.Find(bson.M{"email": user.Email}).One(&userEmail)

	collectionUsers.Find(bson.M{"username": user.Username}).One(&userUsername)

	if len(userEmail.Email) > 0 || len(userUsername.Username) > 0 {
		return ctx.NoContent(http.StatusConflict)
	}

	err = collectionUsers.Insert(user)
	utils.CheckErr(err)

	Disconect()

	return ctx.JSON(http.StatusCreated, user)
}
