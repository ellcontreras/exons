package actions

import (
	"gopkg.in/mgo.v2"

	"comm/utils"
)

var (
	err                   error
	session               *mgo.Session
	collectionCommunities *mgo.Collection
	collectionUsers       *mgo.Collection
)

func Connect() {
	session, err = mgo.Dial("mongodb://localhost:27017/comm")
	utils.CheckErr(err)

	collectionCommunities = session.DB("comm").C("communities")
	collectionUsers = session.DB("comm").C("users")
}

func Disconect() {
	session.Close()
}
