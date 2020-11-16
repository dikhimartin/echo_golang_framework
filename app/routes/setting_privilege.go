package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)

func SettingPrivilege(g *echo.Group) {
	DEFINE_URL := "/setting/privilege"

	g.GET(DEFINE_URL+"/", controllers.ListSettingPrivilege)
	g.GET(DEFINE_URL+"/addform/", controllers.AddSettingPrivilege)
	// g.POST(DEFINE_URL+"/addform/", controllers.StoreSettingPrivilege)
	g.GET(DEFINE_URL+"/editform/:id/", controllers.EditSettingPrivilege)
	// g.POST(DEFINE_URL+"/editform/:id/", controllers.UpdateSettingPrivilege)
}
