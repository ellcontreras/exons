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

<<<<<<< HEAD
	collectionCommunities = session.DB("exons").C("communities")
=======
	collectionexonsunities = session.DB("exons").C("exonsunities")
>>>>>>> 7ffc66964085de8bd601e8428984ada54100d2de
	collectionUsers = session.DB("exons").C("users")
}

func Disconect() {
	session.Close()
}
