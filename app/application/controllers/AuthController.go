package controllers

import (
	"net/http"
	"../../database"
	"github.com/labstack/echo"
)

func FormLogin(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	return c.Render(200, "form_login", nil)
}

func AuthorizationSignIn(c echo.Context) error {
	formusername := c.FormValue("username")
	formpassword := c.FormValue("password")

	// check_authentification
	check_authentification := CheckPassword(formusername, formpassword)
	if check_authentification == "username_false"{
		return c.Redirect(http.StatusTemporaryRedirect, "/?login_verification=username_false")
	}else if check_authentification == "password_false"{
		return c.Redirect(http.StatusTemporaryRedirect, "/?login_verification=password_false")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/lib/sign/redirect/")
}


func CheckPassword(formusername, formpassword string) string{
	db := database.CreateCon()
	defer db.Close()

	//check_username
	var username, password []byte
	row := db.Table("tb_setting_user").Where("username = ?", formusername).Select("username, password").Row() // (*sql.Row)
	err := row.Scan(&username, &password)
	if err != nil{
		logs.Println("users = "+ string(username) + " Has Failed Loged in by username false")
		logs.Println(err)
		return "username_false"
	}

	// check_password
	validate_password := CheckPasswordHash(formpassword, string(password))
	if validate_password == true{
		logs.Println("users = "+ string(username) + " Has Loged in Succesfully")
	}else{
		logs.Println("users = "+ string(username) + " Has Failed Loged in by password false")
		return "password_false"
	}
	return "success"
}