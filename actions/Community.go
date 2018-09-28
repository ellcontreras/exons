package actions

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"net/http"
	"github.com/labstack/echo"

	"comm/utils"
	"comm/models"
	"log"
)

var (
	err error
	session *mgo.Session
	collection *mgo.Collection
)

func Connect() {
	session, err = mgo.Dial("mongodb://localhost:27017/comm")
	utils.CheckErr(err)

	collection = session.DB("comm").C("communities")
}

func Disconect() {
	session.Close()
}

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

//CommunityPut ...
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
