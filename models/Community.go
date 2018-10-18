package models

import (
	"github.com/labstack/echo"

	"exons/utils"

	"gopkg.in/mgo.v2/bson"
)

type Community struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
}

func (c *Community) BindWithContext(ctx echo.Context) {
	err := ctx.Bind(c)

	utils.CheckErr(err)
}
