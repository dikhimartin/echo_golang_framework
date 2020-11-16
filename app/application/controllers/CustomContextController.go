package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)


type CustomContext struct {
	echo.Context
}

func (c *CustomContext) error() (err error) {
	return c.Render(http.StatusInternalServerError, "error_500", nil)
}
