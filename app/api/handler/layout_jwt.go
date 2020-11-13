package handler

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//main jwt
func mainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Println("User name : ", claims["name"], "User ID :", claims["jti"])

	return c.String(http.StatusOK, "You are on the top secret jwt page!")
}
