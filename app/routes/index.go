package routes

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo"
	"time"
	"../api"
	api_middleware "../api/mymiddleware"
	red "../api/redis"
	"../routes/dashboard"
	"github.com/dikhimartin/redis"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"	
)

type M map[string]interface{}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func GetRedis(c echo.Context) error {

	param  := c.QueryParam("key")
	client := red.Connection()

	val, err := client.Get(param).Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key", val)
	}

	//start get count key
	var cursor uint64
	var n int
	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(cursor, "key*", 10).Result()
		if err != nil {
			panic(err)
		}
		n += len(keys)
		if cursor == 0 {
			break
		}
	}
	fmt.Printf("found %d keys\n", n)
	//end get count key

	dataredis := M{
		"val": val,
	}

	return c.JSON(http.StatusOK, dataredis)
}

func DeleteRedis(c echo.Context) error {
	// start delete key
	key := c.QueryParam("key")
	client := red.Connection()

	index, err := client.Del(key).Result()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, err)
	}
	fmt.Println(" Index ke-", index)
	// end delete key

	//set expire redis
	expire := client.Expire(key, 0*time.Second)
	fmt.Println("status:", expire)

	return c.JSON(http.StatusOK, index)
}

func Index() *echo.Echo {
	e := echo.New()
	// Adds a trailing slash to the request URI
	e.Pre(middleware.AddTrailingSlash())
	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// this logs the server interaction
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required",
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}

				break
			}
		}

		c.JSON(report.Code, report)
	}

	//Server header
	e.Use(api_middleware.ServerHeader)
	//end server header

	// start log middleware
	api_middleware.LogMiddleware(e)
	//end log middleware

	//set group name cookie
	cookieGroup := e.Group("/lib")
	//end set group  name cookie

	//handler check redis with cookie
	api.RouteHandlerRedisWithCookie(cookieGroup)
	//end handler cookie

	//start redirect sigin
	api.RedirectSignInFunc(cookieGroup)
	//end redirect signin

	//get data jwt
	api.GetDataJWT(cookieGroup)
	//end data jwt

	//start authorization
	api.RouteAuthorzation(e)
	//end authorization

	Map := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
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

	//routeuser
	UserRoute(e.Group("/users"))

	//Dashbord
	dashboard.MyDashboardRoute(cookieGroup)

	// Modul Sample Crud
	SampleCrud(cookieGroup)

	// Modul Setting
	SettingGrup(cookieGroup)
	SettingPrivilege(cookieGroup)
	SettingUser(cookieGroup)
	SettingGrupPrivilege(cookieGroup)

	//Profile 
	MyProfile(cookieGroup)
	
	//Auth 
	AppAuth(e.Group("login"))

	return e
}
