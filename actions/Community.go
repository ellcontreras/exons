package actions

import (
	"gopkg.in/mgo.v2/bson"

	"net/http"

	"github.com/labstack/echo"

	"exons/models"
	"exons/utils"
	"log"
)

// CommunityGetOne ...
func CommunityGetOne(ctx echo.Context) error {
	Community := models.Community{}

	Connect()

	collectionCommunities.FindId(bson.ObjectIdHex(ctx.Param("id"))).One(&Community)

	Disconect()
	return ctx.JSON(http.StatusOK, Community)
}

// CommunityGetAll ...
func CommunityGetAll(ctx echo.Context) error {
	var communities []models.Community

	Connect()

	err := collectionCommunities.Find(nil).Sort("-_id").All(&communities)
	utils.CheckErr(err)

	Disconect()
	return ctx.JSON(http.StatusOK, communities)
}

// CommunityAdd ...
func CommunityAdd(ctx echo.Context) error {
	Community := models.Community{}

	Community.BindWithContext(ctx)

	Connect()

	err = collectionCommunities.Insert(Community)
	utils.CheckErr(err)

	Disconect()

	return ctx.JSON(http.StatusOK, Community)
}

//CommunityUpdate ...
func CommunityUpdate(ctx echo.Context) error {
	Community := models.Community{}

	Community.BindWithContext(ctx)

	log.Println(Community.ID)

	Connect()

	err := collectionCommunities.UpdateId(Community.ID, Community)
	utils.CheckErr(err)

	Disconect()

	return ctx.JSON(http.StatusOK, Community)
}

// CommunityDelete ...
func CommunityDelete(ctx echo.Context) error {
	Community := models.Community{}

	Community.BindWithContext(ctx)

	Connect()

	collectionCommunities.RemoveId(Community.ID)

	Disconect()

	return ctx.NoContent(http.StatusOK)
}
