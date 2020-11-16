package lib

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var COOKIE_NAME string     = "SESSION"
var JWT_KEY     string 	   = "MYSECRET"

//struct jwt
type JwtClaims struct {
	jwt.StandardClaims
}

func GetKeyJwt(c echo.Context) string{
	jwt, err := GetDataJWT(c)
	if err != nil{
		logs.Println(err)
	}
	return jwt["jti"]
}

func GetDataJWT(c echo.Context) (map[string]string, error) {
	cookie, err := c.Cookie(COOKIE_NAME)
	var errs error
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			logs.Println("named named cookie not present")
		}
		return nil, err
	}
	if cookie.Value != "" {
		//start connection to redis
		client := RedisConnection()
		//start get token with cookie value
		token, err := client.Get(cookie.Value).Result()
		if err != nil {
			logs.Println(err)
			return nil, err
		}
		//end get token with cookie value

		if token != "" {
			// get data with jwt
			tokenString := token
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(tokenString, claims, func(get_token *jwt.Token) (interface{}, error) {
				return []byte(JWT_KEY), nil
			})
			// jika session habis
			if err != nil {	
				// RedirectSessionExpired(c)
				return nil, err
			}
			mapJWT := make(map[string]string)
			// add data to mapJWT
			for key, v := range claims {
				mapJWT[fmt.Sprintf("%v", key)] = fmt.Sprintf("%v", v)
			}			
			return mapJWT, nil
		}
		return nil, errs
	}

	return nil, errs
}

//create jwt
func CreateJwtToken(claims JwtClaims) (string, error) {
	rawToken   := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(JWT_KEY))
	if err != nil {
		fmt.Println("createJwtToken")
		return "", err
	}
	return token, nil
}

//check jwt
func MainMiddleware(e *echo.Group) {
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(JWT_KEY),
		TokenLookup:   "header:UMKMKita",
	}))
}

func ShowDataJWT(c echo.Context) error {
	JSON, err := GetDataJWT(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, JSON)
}


