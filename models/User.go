package models

import (
	"github.com/labstack/echo"

	"gopkg.in/mgo.v2/bson"
	"exons/utils"
)

type User struct {
	ID bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name string	`json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (u *User) BindWithContext(ctx echo.Context) {
	err := ctx.Bind(u)

	utils.CheckErr(err)
}
