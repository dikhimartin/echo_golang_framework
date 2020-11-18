package main

import (
	"./customlogger"
	"./logincache"
	"./routes"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// connect to cache
	logincache.InitCache()

	//route index
	e := routes.Index()
	e.Validator = &CustomValidator{validator: validator.New()}

	logger := customlogger.GetInstance("SYSTEM")
	logger.Println("Starting Application")

	e.Logger.Fatal(e.Start(":2222"))
}
