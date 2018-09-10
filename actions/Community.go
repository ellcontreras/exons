package actions

import (
	"net/http"
	"github.com/labstack/echo"
)

//CommunityGet ...
func CommunityGetOne(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Get Method")
}

func CommunityGetAll(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Get all communities")
}

//CommunityPost ...
func CommunityAdd(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Post Method")
}

//CommunityPut ...
func CommunityUpdate(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Put Method")
}

