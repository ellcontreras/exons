package actions

import (
	"gopkg.in/mgo.v2/bson"

	"net/http"
	"github.com/labstack/echo"

	"comm/utils"
	"comm/models"
	"log"
)

//CommunityGet ...
func CommunityGetOne(ctx echo.Context) error {
	community := models.Community{}

	Connect()

	collection.FindId(bson.ObjectIdHex(ctx.Param("id"))).One(&community)

	Disconect()
	return ctx.JSON(http.StatusOK, community)
}

func CommunityGetAll(ctx echo.Context) error {
	var communities []models.Community
	
	Connect()

	err := collection.Find(nil).Sort("-_id").All(&communities)
	utils.CheckErr(err)

	Disconect()
	return ctx.JSON(http.StatusOK, communities)
}

//CommunityPost ...
func CommunityAdd(ctx echo.Context) error {
	community := models.Community{}

	community.BindWithContext(ctx)

	Connect()

	err = collection.Insert(community)
	utils.CheckErr(err)
	
	Disconect()

	return ctx.JSON(http.StatusOK, community)
}

//CommunityUpdate ...
func CommunityUpdate(ctx echo.Context) error {
	community := models.Community{}

	community.BindWithContext(ctx)


	log.Println(community.ID)
	
	Connect()

	err := collection.UpdateId(community.ID, community)
	utils.CheckErr(err)

	Disconect()

	return ctx.JSON(http.StatusOK, community)
}

// CommunityDelete ...
func CommunityDelete(ctx echo.Context) error {
	community := models.Community{}

	community.BindWithContext(ctx)

	Connect()

	collection.RemoveId(community.ID)

	Disconect()

	return ctx.NoContent(http.StatusOK)
}
