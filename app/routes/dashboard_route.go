package routes

import (
	"receipt/application/controllers"
	"github.com/labstack/echo"
)

func DashboardRoute(g *echo.Group) {
	g.GET("/", controllers.Dashboard)
}
