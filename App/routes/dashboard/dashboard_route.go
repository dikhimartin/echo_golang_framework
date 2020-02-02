package dashboard

import (
	"../../application/controllers"
	"github.com/labstack/echo"
)

func MyDashboardRoute(g *echo.Group) {
	g.GET("/dashboard/", controllers.ListDashboard)
	g.POST("/dashboard/getinfologin/", controllers.Getinfologin)
	g.GET("/dashboard/getsidebarprivilege/", controllers.GetSidebarPrivilege)
}
