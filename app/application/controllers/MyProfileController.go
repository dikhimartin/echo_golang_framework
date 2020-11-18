package controllers

import (
	"os"
	"io"
	"strconv"
	"receipt/application/models"
	"receipt/database"
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

	formpassword 	:= c.FormValue("password_confirm")
	formusername 	:= data_users.Username
	type_submit     := c.FormValue("type_submit")

	// check_authentification
	check_authentification := CheckPassword(formusername, formpassword)
	if check_authentification == "username_false"{
		return c.JSON(200, "username_false")
	}else if check_authentification == "password_false"{
		return c.JSON(200, "password_false")
	}

	// reupload_image
	image 	      := FormFile(c, "image")
	if image != "nil"{
		path 	   := "upload/profile_user/"
		if data_users.Image != ""{
			remove_file := RemoveFile(c, path + "/" + data_users.Image)
			if remove_file == 0 {
				logs.Println("error_remove_file")
				return c.Render(500, "error_500", nil)
			}
			logs.Println("remove_file = ", remove_file)
		}

		// Path_Destinasi_file
		folderPath := MakeDirectory(path)
		if folderPath == "0" {
			logs.Println("error_create_directory")
			return c.Render(500, "error_500", nil)
		}

		dst, err := os.Create(folderPath + "/" + image)
		if err != nil {
			logs.Println(err)
			return c.Render(500, "error_500", nil)
		}
		defer dst.Close()

		// Eksekusi Simpan FIle
		file, _ 	   := c.FormFile("image")
		file_image, _  := file.Open()
		defer file_image.Close()		
		if _, err = io.Copy(dst, file_image); err != nil {
			logs.Println(err)
			return c.Render(500, "error_500", nil)
		}
	}else{
		image = data_users.Image
	}	


	// update account
	if type_submit == "profile"{

		full_name := c.FormValue("full_name")
		email 	  := c.FormValue("email")
		telephone := c.FormValue("telephone")
		address   := c.FormValue("address")
		gender 	  := c.FormValue("gender")

		var update models.SettingUser
		update_user := db.Model(&update).Where("id = ?", data_users.Id_user).Updates(map[string]interface{}{
			"full_name"    :    full_name,
			"email"    	   :    email,
			"telephone"    :    telephone,
			"address"      :    address,
			"gender"       :    gender,
			"image"        :    image,
			"updated_at"   :    current_time("2006-01-02 15:04:05"),
		})
		if update_user.Error != nil {
			logs.Println(update_user.Error)
			return c.Render(500, "error_500", nil)
		}
	}else if type_submit == "account"{

		username 	 		 := c.FormValue("username")
		password_val 		 := c.FormValue("password")
		confirm_password_val := c.FormValue("confirm_password")

		// Update Data User
		if password_val == "" && confirm_password_val == ""{
			var update models.SettingUser
			update_user := db.Model(&update).Where("id = ?", data_users.Id_user).Updates(map[string]interface{}{
				"username"     :    username,
				"updated_at"   :    current_time("2006-01-02 15:04:05"),
			})
			if update_user.Error != nil {
				logs.Println(update_user.Error)
				return c.Render(500, "error_500", nil)
			}
		}else if password_val != "" && confirm_password_val != "" {
			var update models.SettingUser
			update_user := db.Model(&update).Where("id = ?", data_users.Id_user).Updates(map[string]interface{}{
				"password"     :    HashPassword(password_val),
				"updated_at"   :    current_time("2006-01-02 15:04:05"),
			})
			if update_user.Error != nil {
				logs.Println(update_user.Error)
				return c.Render(500, "error_500", nil)
			}
		}
	}


	return c.JSON(200, "true")
}


func UpdateInlineProfile(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	id_users 	:=	data_users.Id_user
	
	field 		:= c.Param("field")
	value 	    := c.FormValue("value")

	var update models.SettingUser
	update_user := db.Model(&update).Where("id = ?", id_users).Updates(map[string]interface{}{
		""+ field +""    :    value,
	})
	if update_user.Error != nil {
		logs.Println(update_user.Error)
		return c.JSON(500, "error")
	}

	return c.JSON(200, value)
}
