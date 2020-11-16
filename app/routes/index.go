package routes

import (
	"net/http"
	"html/template"
	"io"
	lib "../lib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Index() *echo.Echo {
	e := echo.New()

	// handling error_not_found
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.Render(http.StatusInternalServerError, "error_404", nil)
	}

	// Adds a trailing slash to the request URI
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(lib.ServerHeader)

	// start log middleware
	lib.LogMiddleware(e)

	//set group_auth
	AuthGroup := e.Group("/lib")

	//check cookie & Cookie
	RouteHandlerRedisWithCookie(AuthGroup)

	//authorization
	RouteAuthorzation(e)

	//redirect sigin
	RedirectSignIn(AuthGroup)


	Map := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseFiles(
			"template/dashboard.html",
			"template/header.html",
			"template/layout.html",
			"template/sidebar_layout.html",
			"template/sidebar_privilege.html",
			"template/top-nav.html",
			"template/footer.html",
			"template/lastfooter.html",

			// View Error
			"view/error/500.html",
			"view/error/403.html",
			"view/error/404.html",

			/*
			 |--------------------------------------------------------------------------
			 | MODUL SAMPLE CRUD
			 |--------------------------------------------------------------------------
			*/
			"view/sample_crud/list.html",
			"view/sample_crud/add-form.html",
			"view/sample_crud/edit-form.html",


			/*
			 |--------------------------------------------------------------------------
			 | MODUL SETTING
			 |--------------------------------------------------------------------------
			*/

			// GROUP
			"view/setting/grup/list.html",
			"view/setting/grup/add-form.html",
			"view/setting/grup/edit-form.html",

			// PRIVILEGE
			"view/setting/privilege/list.html",
			"view/setting/privilege/add-form.html",
			"view/setting/privilege/edit-form.html",

			// USER
			"view/setting/user/list.html",
			"view/setting/user/add-form.html",
			"view/setting/user/edit-form.html",
			
			// GRUP PRIVILEGE
			"view/setting/grup_privilege/list.html",
			"view/setting/grup_privilege/add-form.html",
			"view/setting/grup_privilege/edit-form.html",
			"view/setting/grup_privilege/show-permissions.html",


			//PROFILE
			"view/profile/my-profile.html",


			//LOGIN
			"view/auth/login.html",
			
		)).Funcs(Map),
	}
	e.Renderer = renderer


	// define_path
	e.Static("static", "assets")
	e.Static("upload", "upload")

	

	//Dashbord
	DashboardRoute(AuthGroup)

	// Modul Sample Crud
	// SampleCrud(cookieGroup)

	// Modul Setting
	SettingGrup(AuthGroup)
	SettingPrivilege(AuthGroup)
	// SettingUser(cookieGroup)
	// SettingGrupPrivilege(cookieGroup)

	//Profile 
	// MyProfile(cookieGroup)
	
	ApiRoute(AuthGroup)

	return e
}
