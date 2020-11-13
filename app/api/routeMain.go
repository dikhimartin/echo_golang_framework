package api

import (
	"../api/data"
	"../api/handler"
	"github.com/labstack/echo"
)

func RouteHandlerRedisWithCookie(g *echo.Group) {
	g.Use(handler.CheckRedisWithCookie)
}

func RouteHandlerSession(g *echo.Group) {
	g.Use(handler.CheckSession)
}

func RouteAuthorzation(e *echo.Echo) {
	e.GET("/", handler.FormSignIn)
	e.POST("/masuk/", handler.AuthorizationSignIn)
	e.GET("/logout/", handler.AuthorizationSignOut)
	e.GET("/logout/redirect/", handler.RedirectLogout)
	e.POST("/", handler.POSTRedirectLogout)
	// session_expire
	e.GET("/session_expire/", handler.AuthorizationSession)
	e.GET("/session_expire/redirect/", handler.RedirectSession)
}

func GetDataJWT(g *echo.Group) {
	g.GET("/data/jwt/", data.ShowDataJWT)
}

func RedirectSignInFunc(g *echo.Group) {
	g.POST("/sign/redirect/", handler.RedirectSignIn)
}
