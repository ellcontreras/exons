package models

import (
	"github.com/labstack/echo"
)

type Community struct {
	uid string `json:'uid'; bson:object_id`
	title string `json:'title'`
	description string `json:'description'`
}

//Getters

func (c *Community) GetUID() string {
	return c.uid
}

func (c *Community) GetTitle() string {
	return c.title
}

func (c *Community) GetDescription() string {
	return c.description
}

// Setters

func (c *Community) SetUID(uid string) {
	c.uid = uid
}

func (c *Community) SetTitle(title string) {
	c.title = title
}

func (c *Community) SetDescription(description string) {
	c.description = description
}

func (c *Community) BindWithContext(ctx echo.Context) {
	ctx.Bind(c)
}
