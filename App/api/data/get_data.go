package data

import (
	"fmt"
	"net/http"
	"strings"
	api_middleware "../mymiddleware"
	red "../redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)


//Get GetDataJWT
func GetDataJWT(c echo.Context) (map[string]string, error) {

	cookie, err := c.Cookie(api_middleware.COOKIE_NAME)

	var errs error

	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			fmt.Println("named named cookie not present")
		}

		return nil, err
	}

	if cookie.Value != "" {

		//start connection to redis
		client := red.Connection()
		//end connection

		//start get token with cookie value
		token, err := client.Get(cookie.Value).Result()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		//end get token with cookie value

		if token != "" {

			// get data with jwt
			tokenString := token
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(tokenString, claims, func(get_token *jwt.Token) (interface{}, error) {
				return []byte(api_middleware.JWT_KEY), nil
			})

			// jika session habis
			if err != nil {	
				RedirectSessionExpired(c)
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

func RedirectSessionExpired(c echo.Context) error{
	return c.Redirect(http.StatusTemporaryRedirect, "/session_expire")
}

func ShowDataJWT(c echo.Context) error {

	JSON, err := GetDataJWT(c)

	if err != nil {
		fmt.Println("kosong")
		return err
	}

	return c.JSON(http.StatusOK, JSON)
}
