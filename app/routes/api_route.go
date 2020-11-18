package routes

import (
	"receipt/application/controllers/api"
	"github.com/labstack/echo"
)

func ApiRoute(g *echo.Group) {
	DEFINE_URL := "/api"

	g.POST(DEFINE_URL +"/getinfologin/", api.Getinfologin)
	g.GET(DEFINE_URL  +"/getsidebarprivilege/", api.GetSidebarPrivilege)

}
