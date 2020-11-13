package handler

import (
	"net/http"

	"github.com/labstack/echo"
	// "fmt"
)

//check cookie
func CheckSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// result, err := sessionManager.Get(c, GET_ID_SESSION)
		// if err != nil {
		// 	return err
		// }

		// if result == nil {
		// 	return c.String(http.StatusOK, "empty result")
		// } else {
		// 	user := result.(UserModel)
		// 	fmt.Println(result)

		// 	_, err := json.Marshal(user)
		// 	// periksa apakah penyandian berhasil
		// 	if err != nil {
		// 		panic(err)
		// 	}

		// 	return next(c)
		// }

		// log.Println("you dont have the right awaaion")
		return c.String(http.StatusUnauthorized, "you dont have the right cookie")

		// return c.Redirect(200, "/")
	}
}
