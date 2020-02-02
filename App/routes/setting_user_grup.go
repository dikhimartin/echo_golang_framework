package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)

func SettingUserGrup(g *echo.Group) {

	DEFINE_URL := "/setting/user_grup"
	g.GET(DEFINE_URL+"/", controllers.ListSettingUserGrup)
	g.GET(DEFINE_URL+"/addform/", controllers.AddSettingUserGrup)
	g.POST(DEFINE_URL+"/addform/", controllers.StoreSettingUserGrup)
	g.GET(DEFINE_URL+"/editform/:id/", controllers.EditSettingUserGrup)
	g.POST(DEFINE_URL+"/editform/:id/", controllers.UpdateSettingUserGrup)
}
