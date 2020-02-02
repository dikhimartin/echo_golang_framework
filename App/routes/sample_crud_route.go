package routes

import (
	"../application/controllers"
	"github.com/labstack/echo"
)

func SampleCrud(g *echo.Group) {

	DEFINE_URL := "/sample_crud"

	g.GET(DEFINE_URL+"/", controllers.ListSampleCrud)
	g.GET(DEFINE_URL+"/addform/", controllers.AddSampleCrud)
	g.POST(DEFINE_URL+"/addform/", controllers.StoreSampleCrud)
	g.GET(DEFINE_URL+"/editform/:id/", controllers.EditSampleCrud)
	g.POST(DEFINE_URL+"/editform/:id/", controllers.UpdateSampleCrud)
	g.POST(DEFINE_URL+"/deleteform/:id/", controllers.DeleteSampleCrud)
	g.POST(DEFINE_URL+"/delete_all/:id/", controllers.DeleteAllSampleCrud)
}
