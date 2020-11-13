package routes

import (
	"../application/controllers"
	lib "../lib"
	"github.com/labstack/echo"
)

func RouteHandlerRedisWithCookie(g *echo.Group) {
	g.Use(lib.CheckRedisWithCookie)
}

func RouteAuthorzation(e *echo.Echo) {
	e.GET("/", controllers.FormLogin)
	e.POST("/login/", controllers.AuthorizationSignIn)
	// e.GET("/logout/", lib.AuthorizationSignOut)
	// e.GET("/logout/redirect/", lib.RedirectLogout)
	// e.POST("/", lib.POSTRedirectLogout)

	// session_expire
	e.GET("/session_expire/", lib.AuthorizationSession)
	e.GET("/session_expire/redirect/", lib.RedirectSession)
}


func AuthRoute(g *echo.Group) {
	g.GET("/", controllers.FormLogin)
}
