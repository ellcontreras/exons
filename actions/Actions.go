package actions

import (
	"gopkg.in/mgo.v2"

	"exons/utils"
)

var (
	err                   error
	session               *mgo.Session
	collectionexonsunities *mgo.Collection
	collectionUsers       *mgo.Collection
)

func Connect() {
	session, err = mgo.Dial("mongodb://localhost:27017/exons")
	utils.CheckErr(err)

	collectionexonsunities = session.DB("exons").C("exonsunities")
	collectionUsers = session.DB("exons").C("users")
}

func Disconect() {
	session.Close()
}
