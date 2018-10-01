package actions

import (
	"gopkg.in/mgo.v2/bson"

	"net/http"

	"github.com/labstack/echo"

	"exons/models"
	"exons/utils"
	"log"
)

// exonsunityGetOne ...
func exonsunityGetOne(ctx echo.Context) error {
	exonsunity := models.exonsunity{}

	Connect()

	collectionexonsunities.FindId(bson.ObjectIdHex(ctx.Param("id"))).One(&exonsunity)

	Disconect()
	return ctx.JSON(http.StatusOK, exonsunity)
}

// exonsunityGetAll ...
func exonsunityGetAll(ctx echo.Context) error {
	var exonsunities []models.exonsunity

	Connect()

	err := collectionexonsunities.Find(nil).Sort("-_id").All(&exonsunities)
	utils.CheckErr(err)

	Disconect()
	return ctx.JSON(http.StatusOK, exonsunities)
}

// exonsunityAdd ...
func exonsunityAdd(ctx echo.Context) error {
	exonsunity := models.exonsunity{}

	exonsunity.BindWithContext(ctx)

	Connect()

	err = collectionexonsunities.Insert(exonsunity)
	utils.CheckErr(err)

	Disconect()

	return ctx.JSON(http.StatusOK, exonsunity)
}

//exonsunityUpdate ...
func exonsunityUpdate(ctx echo.Context) error {
	exonsunity := models.exonsunity{}

	exonsunity.BindWithContext(ctx)

	log.Println(exonsunity.ID)

	Connect()

	err := collectionexonsunities.UpdateId(exonsunity.ID, exonsunity)
	utils.CheckErr(err)

	Disconect()

	return ctx.JSON(http.StatusOK, exonsunity)
}

// exonsunityDelete ...
func exonsunityDelete(ctx echo.Context) error {
	exonsunity := models.exonsunity{}

	exonsunity.BindWithContext(ctx)

	Connect()

	collectionexonsunities.RemoveId(exonsunity.ID)

	Disconect()

	return ctx.NoContent(http.StatusOK)
}
