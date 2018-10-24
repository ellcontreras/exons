package actions

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"exons/models"
	"exons/utils"
)

// UserAdd ...
func UserAdd(ctx echo.Context) error {
	user := models.User{}
	userEmail := models.User{}
	userUsername := models.User{}

	user.BindWithContext(ctx)

	Connect()

	collectionUsers.Find(bson.M{"email": user.Email}).One(&userEmail)

	collectionUsers.Find(bson.M{"username": user.Username}).One(&userUsername)

	if len(userEmail.Email) > 0 || len(userUsername.Username) > 0 {
		return ctx.NoContent(http.StatusConflict)
	}

	err = collectionUsers.Insert(user)
	utils.CheckErr(err)

	Disconect()

	return ctx.JSON(http.StatusCreated, user)
}

// UserLogin ...
func UserLogin(ctx echo.Context) error {
	user := models.User{}
	userDB := models.User{}

	user.BindWithContext(ctx)

	log.Println(user)

	Connect()

	dataStruct := bson.M{
		"email":    user.Email,
		"password": user.Password,
	}

	collectionUsers.Find(dataStruct).One(&userDB)

	Disconect()

	if len(userDB.Name) > 0 {
		claims := &utils.JWTCustomClaims{
			ID:       userDB.ID,
			Name:     userDB.Name,
			Email:    userDB.Email,
			Username: userDB.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("secret"))
		utils.CheckErr(err)

		return ctx.JSON(http.StatusOK, echo.Map{
			"token": t,
			"user":  userDB,
		})
	}

	return ctx.NoContent(http.StatusUnauthorized)
}

// UserGetOne ...
func UserGetOne(ctx echo.Context) error {
	user := models.User{}

	log.Print(ctx.Param("id"))

	Connect()

	collectionUsers.FindId(bson.ObjectIdHex(ctx.Param("id"))).One(&user)

	Disconect()

	return ctx.JSON(http.StatusOK, user)
}
