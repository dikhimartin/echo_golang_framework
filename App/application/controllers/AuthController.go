package controllers

import (
	// "fmt"

	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	// "../models"
	// "../database"
	"../../customlogger"
	"../../logincache"
	"github.com/labstack/echo"
)

var authCode = "026556298d4ff7fc742a2daeb1748b0a"

func FormLogin(c echo.Context) error {
	customlogger.SetUser("LOGIN")
	customlogger.Debug("Login Page")
	return c.Render(http.StatusOK, "form_login", nil)
}

func ProsesLogin(c echo.Context) error {

	idUser := "1"

	// Create a new random session token
	sessionToken := createToken(idUser)

	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	_, err := logincache.Cache.Do("SETEX", sessionToken, "120", "admin")
	if err != nil {
		// If there is an error in setting the cache, return an internal server error
		// c.Response().WriteHeader(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"token": sessionToken,
		})
		// return c.Render(http.StatusInternalServerError, "form_login", nil)
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds, the same as the cache
	// http.SetCookie(c.Response(), &http.Cookie{
	// 	Name:    "session_token",
	// 	Value:   sessionToken,
	// 	Expires: time.Now().Add(120 * time.Second),
	// })
	writeCookie(c, sessionToken)

	// return c.Redirect(301, "/")

	return c.JSON(http.StatusOK, echo.Map{
		"token": sessionToken,
	})

	// return c.Render(http.StatusOK, "form_login", nil)
}

func createToken(idUser string) string {
	tm := time.Now().UnixNano()

	// logger.Println(tm)
	// return strconv.FormatInt(tm, 21)

	hasher := md5.New()
	hasher.Write([]byte(strconv.FormatInt(tm, 21) + authCode + idUser))
	return hex.EncodeToString(hasher.Sum(nil))
}

func writeCookie(c echo.Context, sessionToken string) {
	cookie := new(http.Cookie)
	cookie.Name = "_t"
	cookie.Value = sessionToken
	cookie.Expires = time.Now().Add(24 * time.Second)
	c.SetCookie(cookie)
	// return c.String(http.StatusOK, "write a cookie")
}

func readCookie(c echo.Context, cookieName string) string {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return err.Error()
	}
	return cookie.Value
	// return c.String(http.StatusOK, "read a cookie")
}
