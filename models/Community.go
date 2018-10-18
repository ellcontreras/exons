package models

import (
	"github.com/labstack/echo"

	"gopkg.in/mgo.v2/bson"
	"exons/utils"
)

type exonsunity struct {
	ID bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title string `json:"title"`
	Description string `json:"description"`
}

func (c *exonsunity) BindWithContext(ctx echo.Context) {
	err := ctx.Bind(c)

	utils.CheckErr(err)
}
