package lib

import (
	"os"
	"reflect"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// GetEnv func for get .env value
func GetEnv(val string) string {
	e := godotenv.Load(".env")
	if e != nil {
		panic(e)
	}
	return os.Getenv(val)
}

// IsNilInterface func validate param required
func IsNilInterface(request ...interface{}) bool {
	for i := 0; i < len(request); i++ {
		if reflect.ValueOf(request[i]).IsNil() {
			return true
		}
	}
	return false
}

//server header
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "ReceiptGo/0.2")
		return next(c)
	}
}

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))
}
