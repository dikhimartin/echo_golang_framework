package routes

import (
	"receipt/application/controllers"
	"github.com/labstack/echo"
)

func SampleCrud(g *echo.Group) {

	DEFINE_URL := "/sample_crud"

	g.GET(DEFINE_URL+"/", controllers.ListSampleCrudController)
	g.GET(DEFINE_URL+"/addform/", controllers.AddSampleCrudController)
	g.POST(DEFINE_URL+"/addform/", controllers.StoreSampleCrudController)
	g.GET(DEFINE_URL+"/editform/:id/", controllers.EditSampleCrudController)
	g.POST(DEFINE_URL+"/editform/:id/", controllers.UpdateSampleCrudController)
	g.POST(DEFINE_URL+"/deleteform/:id/", controllers.DeleteSampleCrudController)
	g.POST(DEFINE_URL+"/delete_all/:id/", controllers.DeleteAllSampleCrudController)
}
