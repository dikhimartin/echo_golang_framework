package api

import(
	"time"
	controllers ".."
	"github.com/labstack/echo"
)

func Getinfologin(c echo.Context) error{
	data_users	:= controllers.GetDataLogin(c)
	response := response_json{
		"data_users"				:   data_users,
		"time"						:   time.Now().UnixNano(),
	}
	return c.JSON(200, response)
}

