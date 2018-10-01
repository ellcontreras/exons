package actions

import (
	"gopkg.in/mgo.v2"

	"exons/utils"
)

var (
	err                   error
	session               *mgo.Session
	collectionCommunities *mgo.Collection
	collectionUsers       *mgo.Collection
)

func Connect() {
	session, err = mgo.Dial("mongodb://localhost:27017/exons")
	utils.CheckErr(err)

	collectionCommunities = session.DB("exons").C("comminities")
	collectionUsers = session.DB("exons").C("users")
}

func Disconect() {
	session.Close()
}
