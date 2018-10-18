package actions

import (
	"gopkg.in/mgo.v2/bson"

	"net/http"

	"github.com/labstack/echo"

	"exons/models"
	"exons/utils"
	"log"
)

<<<<<<< HEAD
// CommunityGetOne ...
func CommunityGetOne(ctx echo.Context) error {
	Community := models.Community{}

	Connect()

	collectionCommunities.FindId(bson.ObjectIdHex(ctx.Param("id"))).One(&Community)

	Disconect()
	return ctx.JSON(http.StatusOK, Community)
=======
// exonsunityGetOne ...
func exonsunityGetOne(ctx echo.Context) error {
	exonsunity := models.exonsunity{}

	Connect()

	collectionexonsunities.FindId(bson.ObjectIdHex(ctx.Param("id"))).One(&exonsunity)

	Disconect()
	return ctx.JSON(http.StatusOK, exonsunity)
>>>>>>> 7ffc66964085de8bd601e8428984ada54100d2de
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

<<<<<<< HEAD
// CommunityAdd ...
func CommunityAdd(ctx echo.Context) error {
	Community := models.Community{}

	Community.BindWithContext(ctx)

	Connect()

	err = collectionCommunities.Insert(Community)
=======
// exonsunityAdd ...
func exonsunityAdd(ctx echo.Context) error {
	exonsunity := models.exonsunity{}

	exonsunity.BindWithContext(ctx)

	Connect()

	err = collectionexonsunities.Insert(exonsunity)
>>>>>>> 7ffc66964085de8bd601e8428984ada54100d2de
	utils.CheckErr(err)

	Disconect()

<<<<<<< HEAD
	return ctx.JSON(http.StatusOK, Community)
}

//CommunityUpdate ...
func CommunityUpdate(ctx echo.Context) error {
	Community := models.Community{}

	Community.BindWithContext(ctx)

	log.Println(Community.ID)

	Connect()

	err := collectionCommunities.UpdateId(Community.ID, Community)
=======
	return ctx.JSON(http.StatusOK, exonsunity)
}

//exonsunityUpdate ...
func exonsunityUpdate(ctx echo.Context) error {
	exonsunity := models.exonsunity{}

	exonsunity.BindWithContext(ctx)

	log.Println(exonsunity.ID)

	Connect()

	err := collectionexonsunities.UpdateId(exonsunity.ID, exonsunity)
>>>>>>> 7ffc66964085de8bd601e8428984ada54100d2de
	utils.CheckErr(err)

	Disconect()

<<<<<<< HEAD
	return ctx.JSON(http.StatusOK, Community)
}

// CommunityDelete ...
func CommunityDelete(ctx echo.Context) error {
	Community := models.Community{}

	Community.BindWithContext(ctx)

	Connect()

	collectionCommunities.RemoveId(Community.ID)
=======
	return ctx.JSON(http.StatusOK, exonsunity)
}

// exonsunityDelete ...
func exonsunityDelete(ctx echo.Context) error {
	exonsunity := models.exonsunity{}

	exonsunity.BindWithContext(ctx)

	Connect()

	collectionexonsunities.RemoveId(exonsunity.ID)
>>>>>>> 7ffc66964085de8bd601e8428984ada54100d2de

	Disconect()

	return ctx.NoContent(http.StatusOK)
}
