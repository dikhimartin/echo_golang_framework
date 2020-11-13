package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	api_middleware "../mymiddleware"
	red "../redis"
	"github.com/labstack/echo"
)

//check cookie
func CheckRedisWithCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(api_middleware.COOKIE_NAME)

		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				log.Println("err contains")
				log.Println(err)
				return c.Redirect(http.StatusTemporaryRedirect, "/session_expire")
			}

			log.Println("err cookie")
			log.Println(err)
			return c.Redirect(http.StatusTemporaryRedirect, "/session_expire")
		}

		if cookie.Value != "" {

			client := red.Connection()

			val, err := client.Get(cookie.Value).Result()
			if err != nil {
				fmt.Println(err)
			}

			if val != "" {
				return next(c)
			}

			log.Println("Redis Empty")
			return c.Redirect(http.StatusTemporaryRedirect, "/session_expire")

		}

		log.Println("Cookie Empty")
		return c.Redirect(http.StatusTemporaryRedirect, "/session_expire")

	}
}
