package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)

func AppAuth(g *echo.Group) {
	g.GET("/", controllers.FormLogin)
	g.POST("/", controllers.ProsesLogin)
}
