package controllers

import (
	"net/http"
	"time"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/astaxie/beego/utils/pagination"
)

var errorEntryData string
var errorDuplikatData string
var errorFeedback error
var now = time.Now()


// paginate session
var (
	paginator = &pagination.Paginator{}
	data      = pongo2.Context{}
)

type dashboard map[string]interface{}

func ListDashboard(c echo.Context) error {

	cc := &MyCustomContext{c}
	data_users	:= cc.getDataLogin()

	currentTime := time.Now()
	get_date 	:= currentTime.Format("2006-01-02")
	format_date := currentTime.Format("02 January 2006")

	data = pongo2.Context{
		"data_users"				:   data_users,
		"format_date"				:   format_date,
		"get_date"					:   get_date}


	return c.Render(http.StatusOK, "dashboard", data)
}



