package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)

func UserRoute(g *echo.Group) {
	g.POST("/", controllers.AddUser)
	g.GET("/", controllers.GetUsers)
}
