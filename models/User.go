package models

import (
	"github.com/labstack/echo"

	"exons/utils"

	"gopkg.in/mgo.v2/bson"
)

// User ...
type User struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Github      string        `json:"github"`
	Facebook    string        `json:"facebook"`
	Twitter     string        `json:"twitter"`
	Description string        `json:"description"`
}

// BindWithContext ...
func (u *User) BindWithContext(ctx echo.Context) {
	err := ctx.Bind(u)

	utils.CheckErr(err)
}
