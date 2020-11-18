package controllers

import (
	"fmt"
	"strconv"
	"../../database"
	"github.com/labstack/echo"
)


// == View
func MyProfileController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)

	user := GetDataUserById(ConvertToMD5(strconv.Itoa(data_users.Id_user)))

	data := response_json{
		"data_user"	 :   user,
	}
	return c.Render(200, "my_profile", data)
}


// == Manipulate
func UpdateProfileController(c echo.Context) error{
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)

	fmt.Println(data_users)

	return c.JSON(200, "true")
}

