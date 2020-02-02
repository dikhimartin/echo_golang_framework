package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)

func MyProfile(g *echo.Group) {
	DEFINE_URL := "/my_profile"
	g.GET(DEFINE_URL+"/", controllers.MyProfileController)
}
