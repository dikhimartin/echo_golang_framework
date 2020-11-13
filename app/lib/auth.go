package lib

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"net/http"
	"github.com/labstack/echo"
)

// formlogin
func FormSignIn(c echo.Context) error {
	cookie, err := c.Cookie(COOKIE_NAME)
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			SetUser("LOGIN")
			Debug("Login Page")
			return c.Render(200, "form_login", nil)
		}
		logs.Println("err cookie signin")
		logs.Println(err)
		return err
	}

	if cookie.Value == "" {
		SetUser("LOGIN")
		Debug("Login Page")
		return c.Render(200, "form_login", nil)
	}
	return c.Redirect(301, "/lib/")
}

func RedirectSignIn(c echo.Context) error {
	return c.Redirect(301, "/lib?after_login=true")
}


// session_expire
func AuthorizationSession(c echo.Context) error {
	current_time 	   := strconv.FormatInt(time.Now().UnixNano(), 10)
	cookie, err := c.Cookie(COOKIE_NAME)
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			return c.Redirect(http.StatusTemporaryRedirect, "/?session_expire=true=?"+current_time)
		}
		logs.Println("err cookie singout")
		logs.Println(err)
		return err
	}

	if cookie.Value != "" {
		//start connection redis
		client := RedisConnection()

		//start remove redis where key
		index, err := client.Del(cookie.Value).Result()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(" Index ke-", index)
		//end remove redis where  key

		//star set expire redis
		boolean := client.Expire(cookie.Value, 0*time.Second)
		fmt.Println("status:", boolean)
		//end set expire redis

		//start set expire cookie
		expCookie := &http.Cookie{
			Name:    COOKIE_NAME,
			Value:   "",
			Path:    "/",
			Expires: time.Unix(0, 0),

			HttpOnly: true,
		}
		c.SetCookie(expCookie)
		//end set expire cookie
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/?session_expire=true=?"+current_time)
}
func RedirectSession(c echo.Context) error {
	current_time 	   := strconv.FormatInt(time.Now().UnixNano(), 10)	
	return c.Redirect(http.StatusTemporaryRedirect, "/?session_expire=true=?"+current_time)
}