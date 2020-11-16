package controllers

import (
	"time"
	"net/http"
	"../models"
	"../../database"
	lib      "../../lib"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func FormLogin(c echo.Context) error {
	return c.Render(200, "form_login", nil)
}

func AuthorizationSignIn(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	formusername := c.FormValue("username")
	formpassword := c.FormValue("password")

	// check_authentification
	check_authentification := CheckPassword(formusername, formpassword)
	if check_authentification == "username_false"{
		return c.Redirect(http.StatusTemporaryRedirect, "/?login_verification=username_false")
	}else if check_authentification == "password_false"{
		return c.Redirect(http.StatusTemporaryRedirect, "/?login_verification=password_false")
	}

	// get_data_auth
	var id_user []byte
	row := db.Table("tb_setting_user").Where("username = ?", formusername).Select("id").Row() // (*sql.Row)
	err := row.Scan(&id_user)
	if err != nil{
		logs.Println("users = "+ formusername + " Has Failed Loged in by password false")
		logs.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, "/?login_verification=password_false")
	}

	// set_key_token
	key_token := "receipt_go_" + ConvertToMD5("receipt_go_" + string(id_user))

	//set_jwt
	struct_claims := lib.JwtClaims{
		jwt.StandardClaims{
			Id 			: key_token,
			ExpiresAt	: time.Now().Add(time.Duration(ConvertStringToInt(lib.GetEnv("SESSION_DURATION")))  * time.Hour).Unix(),
		},
	}
	token, err := lib.CreateJwtToken(struct_claims)
	if err != nil {
		logs.Println("error create jwt token", err)
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	//set_cookie
	cookie := &http.Cookie{}
	cookie.Name    = lib.COOKIE_NAME
	cookie.Value   = key_token
	cookie.Path    = "/"
	cookie.Expires = time.Now().Add(time.Duration(ConvertStringToInt(lib.GetEnv("SESSION_DURATION")))  * time.Hour)
	c.SetCookie(cookie)

	// set_redis
	err = redis_connect.Set(key_token, token, 0).Err()
	if err != nil {
		logs.Println(err)
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	redis_connect.Expire(key_token, time.Duration(ConvertStringToInt(lib.GetEnv("SESSION_DURATION"))) *time.Hour)


	// // save_auth_token
	var user models.SettingUser
	update_token := db.Model(&user).Where("id = ?", string(id_user)).Update("auth_token", token)
	if update_token.Error != nil{
		logs.Println(update_token.Error)
		return c.Render(http.StatusInternalServerError, "error_500", nil)
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

func GetDataLogin(c echo.Context) (models.GetDataLogin) {
	db := database.CreateCon()
	defer db.Close()

	var id, id_grup, full_name, username, email, telephone, address, gender, status []byte
	row := db.Table("tb_setting_user user").Joins("LEFT JOIN tb_setting_user_grup user_grup ON user.id = user_grup.id_setting_user").Where("auth_token = ?", lib.GetTokenRedis(lib.GetKeyJwt(c))).Select("user.id, user_grup.id_setting_grup,  user.full_name, user.username, user.email, user.telephone, user.address, user.gender, user.status").Row() 
	err := row.Scan(&id, &id_grup, &full_name, &username, &email, &telephone, &address, &gender, &status)
	if err != nil {
		logs.Println(err)
	}

	data_users 			:= models.GetDataLogin{
		Id_user  	: 	ConvertStringToInt(string(id)),
		Id_group  	: 	ConvertStringToInt(string(id_grup)),
		Full_name  	: 	string(full_name),
		Username  	: 	string(username),
		Email  		: 	string(email),
		Telephone  	: 	string(telephone),
		Address  	: 	string(address),
		Gender  	: 	string(gender),
		Status  	: 	string(status),
	}

	return data_users
}





