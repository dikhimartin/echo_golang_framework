package mymiddleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var COOKIE_NAME string = "LBSDA5"
var JWT_KEY string 	   = "MYSECRET"

//struct jwt
type JwtClaims struct {
	Id_users       string `json:"id_users"`
	Id_group       string `json:"id_group"`
	Name_users     string `json:"name_users"`
	Name_group     string `json:"name_group"`
	Image     	   string `json:"image"`
	Extension      string `json:"extension"`
	jwt.StandardClaims
}

//create jwt
func CreateJwtToken(claims JwtClaims) (string, error) {

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte(JWT_KEY))
	if err != nil {
		fmt.Println("createJwtToken")
		return "", err
	}

	return token, nil
}

//server header
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "LibyteTech/0.1")

		return next(c)
	}
}

//create log
func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))
}

//check jwt
func MainMiddleware(e *echo.Group) {
	//jwt middleware
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(JWT_KEY),
		TokenLookup:   "header:Libyte",
	}))
}
