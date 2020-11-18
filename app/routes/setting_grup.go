package routes

import (
	"receipt/application/controllers"
	"github.com/labstack/echo"
)

func SettingGrup(g *echo.Group) {

	DEFINE_URL := "/setting/grup"

	g.GET(DEFINE_URL+"/", controllers.ListSettingGrup)
	g.GET(DEFINE_URL+"/addform/", controllers.AddSettingGrup)
	g.POST(DEFINE_URL+"/addform/", controllers.StoreSettingGrup)
	g.GET(DEFINE_URL+"/editform/:id/", controllers.EditSettingGrup)
	g.POST(DEFINE_URL+"/editform/:id/", controllers.UpdateSettingGrup)
}
