package routes

import (
	"receipt/application/controllers"
	"github.com/labstack/echo"
)

func SettingGrupPrivilege(g *echo.Group) {
	DEFINE_URL := "/setting/grup_privilege"

	g.GET(DEFINE_URL+"/", controllers.ListSettingGrupPrivilege)
	g.GET(DEFINE_URL+"/addform/", controllers.AddSettingGrupPrivilege)
	g.POST(DEFINE_URL+"/addform/", controllers.StoreSettingGrupPrivilege)
	g.GET(DEFINE_URL+"/editform/:id/", controllers.EditSettingGrupPrivilege)
	g.POST(DEFINE_URL+"/editform/:id/", controllers.UpdateSettingGrupPrivilege)
	g.GET(DEFINE_URL+"/show_permissions/:id/", controllers.ShowSettingGrupPrivilege)
}
