package actions

import (
	"gopkg.in/mgo.v2"

	"comm/utils"
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
