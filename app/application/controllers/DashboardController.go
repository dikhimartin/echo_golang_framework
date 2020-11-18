package controllers

import (
	"github.com/labstack/echo"
)


func Dashboard(c echo.Context) error {
	data_users	:= GetDataLogin(c)

	data := response_json{
		"data_users"		:   data_users,
		"date"				:   current_time("02 January 2006"),
	}

	return c.Render(200, "dashboard", data)
}
