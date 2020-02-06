package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"
	"strconv"
	"../models"
	data_user "../../api/data"
	"../../database"
	"database/sql"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)


func MyProfileController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	requested_id := data_users.Id_users

	var id_setting_grup, full_name, gender, email, telephone, address, username, name_grup, status []byte
	err := db.QueryRow("SELECT id_setting_grup, full_name, gender, email, telephone, address,  username, name_grup, status FROM v_get_user WHERE id = ?", requested_id).Scan(&id_setting_grup, &full_name, &gender, &email, &telephone, &address, &username, &name_grup, &status)
	if err != nil {
		logs.Println(err)
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	response := models.SettingUser{
		Id 			:  requested_id,
		Id_group 	:  string(id_setting_grup),
		Full_name 	:  string(full_name),
		Gender 		:  string(gender),
		Email 		:  string(email),
		Telephone 	:  string(telephone),
		Address		:  string(address),
		Username 	:  string(username),
		Name_grup 	:  string(name_grup),
		Status 		:  string(status),
	}


	data = pongo2.Context{
		"setting_user"				:   response,
	}

	return c.Render(http.StatusOK, "my_profile", data)
}

func GetDataProfileController(c echo.Context) error{
	db := database.CreateCon()
	defer db.Close()

	// get data login
	dt_user, err   := data_user.GetDataJWT(c)
	if err != nil{
		logs.Println(err)
	}
	id_users 	   := dt_user["id_users"]
	name_group 	   := dt_user["name_group"]

	// get_data_image
	var full_name, email, telephone, address, image, additional []byte
	data_images, err := db.Prepare("SELECT full_name, email, telephone, address, image, additional FROM tb_setting_user WHERE id = ?")
	if err != nil {
		logs.Println(err)
	}
	err = data_images.QueryRow(id_users).Scan(&full_name, &email, &telephone, &address, &image, &additional)	
	defer data_images.Close()

	data_users := models.GetDataProfile{
	    Id_users 			:    id_users,
	    Name_users 			:    string(full_name),
	    Name_group 			:    name_group,
	    Email 				:    string(email),
	    Telephone 			:    string(telephone),
	    Address 			:    string(address),
	    Image				:    string(image),
	    Extension			:    string(additional),
	}

	data = pongo2.Context{
		"data_users"				:   data_users,
		"time"						:   time.Now().UnixNano(),
	}
	return c.JSON(200, data)
}

func UpdateInlineProfile(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()
	id_users 			:=	data_users.Id_users


	type_post 	:= c.Param("type_post")
	value 	    := c.FormValue("value")

	if type_post == "name"{

		Update, err := db.Prepare("UPDATE tb_setting_user SET full_name=? WHERE id=?")
		if err != nil {
			logs.Println(err)
		}
		defer Update.Close()
		Update.Exec(value, id_users)

	}else if type_post == "email"{

		Update, err := db.Prepare("UPDATE tb_setting_user SET email=? WHERE id=?")
		if err != nil {
			logs.Println(err)
		}
		defer Update.Close()
		Update.Exec(value, id_users)

	}else if type_post == "telephone"{

		Update, err := db.Prepare("UPDATE tb_setting_user SET telephone=? WHERE id=?")
		if err != nil {
			logs.Println(err)
		}
		defer Update.Close()
		Update.Exec(value, id_users)

	}else if type_post == "address"{

		Update, err := db.Prepare("UPDATE tb_setting_user SET address=? WHERE id=?")
		if err != nil {
			logs.Println(err)
		}
		defer Update.Close()
		Update.Exec(value, id_users)

	}


	return c.JSON(http.StatusOK, value)
}

func ConfirmUpdateProfile(c echo.Context) error{

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()
	id_users  			:= data_users.Id_users
	type_submit         := c.FormValue("type_submit")

	// get_current_time
	t 				:= time.Now()
	current_time 	:= t.Format("2006-01-02")

	// get_username
	var usernames []byte
	check_username, err := db.Prepare("SELECT username FROM tb_setting_user WHERE id = ?")
	if err != nil {
		logs.Println(err)
		return c.Render(http.StatusInternalServerError, "error_505", nil)	
	}
	err = check_username.QueryRow(id_users).Scan(&usernames)	
	defer check_username.Close()

	formpassword 	:= c.FormValue("password_confirm");
	formusername 	:= string(usernames);

	// //start cek database
	var salt, password, username string
	sqlStatement := `SELECT salt, password, username FROM v_get_user WHERE username = ?`
	row := db.QueryRow(sqlStatement, formusername)
	errCheck := row.Scan(&salt, &password, &username)
	if errCheck != nil {
		if errCheck == sql.ErrNoRows {
			logs.Println("username_false")
			return c.JSON(200, "username_false")
		} else {
			logs.Println(errCheck)
		}
	}
	// //end check database

	// MD5 PASSWORD
	var str_password string = formpassword
	hasher_password := md5.New()
	hasher_password.Write([]byte(str_password))
	md5password := hex.EncodeToString(hasher_password.Sum(nil))

	// set password dan salt
	var salt_password string = salt + md5password

	hasher_salt_password := md5.New()
	hasher_salt_password.Write([]byte(salt_password))
	get_password := hex.EncodeToString(hasher_salt_password.Sum(nil))

	//validasi password
	if get_password != password{
		logs.Println("password_false")
		return c.JSON(200, "password_false")
	}
	//end validasi password

	// Proses Update Data Users
	if type_submit == "profile"{

		full_name := c.FormValue("full_name")
		gender 	  := c.FormValue("gender")
		email 	  := c.FormValue("email")
		telephone := c.FormValue("telephone")
		address   := c.FormValue("address")

		Update, err := db.Prepare("UPDATE tb_setting_user SET full_name=?, gender=?, email=?, telephone =?, address=? , update_date = ? WHERE id=?")
		if err != nil {
			logs.Println(err)
		}
		defer Update.Close()
		Update.Exec(full_name, gender, email, telephone, address, current_time,  id_users)

	}else if type_submit == "account"{

		username 	 		 := c.FormValue("username")
		password_val 		 := c.FormValue("password")
		confirm_password_val := c.FormValue("confirm_password")


		jam   := strconv.Itoa(now.Hour())
		menit := strconv.Itoa(now.Minute())
		detik := strconv.Itoa(now.Second())

		// SALT
		var str_salt string = jam + menit + detik

		hasher_salt := md5.New()
		hasher_salt.Write([]byte(str_salt))
		salt := hex.EncodeToString(hasher_salt.Sum(nil))

		// MD5 PASSWORD
		var str_password string = password_val
		hasher_password := md5.New()
		hasher_password.Write([]byte(str_password))
		md5password := hex.EncodeToString(hasher_password.Sum(nil))
		// PASSWORD FINAL
		var salt_password string = salt + md5password

		hasher_salt_password := md5.New()
		hasher_salt_password.Write([]byte(salt_password))
		password := hex.EncodeToString(hasher_salt_password.Sum(nil))


		// Update Data User
		if password_val == "" && confirm_password_val == ""{
			logs.Println("Changes Username")

			UpdateUser, err := db.Prepare("UPDATE tb_setting_user SET username = ?, update_date=? WHERE id=?")
			if err != nil{
				logs.Println(err)
				return c.Render(http.StatusInternalServerError, "error_500", nil)			
			}
			UpdateUser.Exec(username, id_users)
			defer UpdateUser.Close()

		}else if password_val != "" && confirm_password_val != "" {
			logs.Println("Changes Password")

			UpdateUser, err := db.Prepare("UPDATE tb_setting_user SET username = ?, password =? , salt =?, update_date=? WHERE id=?")
			if err != nil{
				logs.Println(err)
				return c.Render(http.StatusInternalServerError, "error_500", nil)			
			}
			UpdateUser.Exec(username, password, salt, current_time, id_users)
			defer UpdateUser.Close()

		}

	}

	return c.JSON(200, "true")
}

