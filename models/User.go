package models

import (
	"github.com/labstack/echo"

	"exons/utils"

	"gopkg.in/mgo.v2/bson"
)

// User ...
type User struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
}

// BindWithContext ...
func (u *User) BindWithContext(ctx echo.Context) {
	err := ctx.Bind(u)

	utils.CheckErr(err)
}
