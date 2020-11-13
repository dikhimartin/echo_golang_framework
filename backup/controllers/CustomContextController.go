package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

type M map[string]interface{}

type MyCustomContext struct {
	echo.Context
}

func (c *MyCustomContext) errorMen() (err error) {
	return c.Render(http.StatusInternalServerError, "error_500", nil)
}
