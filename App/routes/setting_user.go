package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)

func SettingUser(g *echo.Group) {
	DEFINE_URL := "/setting/user"
	g.GET(DEFINE_URL+"/", controllers.ListSettingUser)
	g.GET(DEFINE_URL+"/addform/", controllers.AddSettingUser)
	g.POST(DEFINE_URL+"/addform/", controllers.StoreSettingUser)
	g.GET(DEFINE_URL+"/editform/:id/", controllers.EditSettingUser)
	g.POST(DEFINE_URL+"/editform/:id/", controllers.UpdateSettingUser)
	g.POST(DEFINE_URL+"/check_username/", controllers.CheckUsername)
}
