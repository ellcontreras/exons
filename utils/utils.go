package utils

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTCustomClaims
type JWTCustomClaims struct {
	Name  string `json:"name"`
	Username string   `json:"username"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// CheckErr verify if exists an error in the var err
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
