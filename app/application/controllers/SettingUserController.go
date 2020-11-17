package controllers

import (	
	"os"
	"io"
	"strconv"
	"../models"
	"../../database"
	"github.com/labstack/echo"
	"github.com/dikhimartin/beego-v1.12.0/utils/pagination"
)

// == View
func ListSettingUser(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.user_2") == false{
		return c.Render(403, "error_403", nil)
	}

	var selected     string
	var whrs 		 string
	var search  	 string
	var searchStatus string

	if reqSearch := c.FormValue("search"); reqSearch != "" {
		search = reqSearch
	}
	if reqSearchStatus := c.FormValue("searchStatus"); reqSearchStatus != "" {
		searchStatus = reqSearchStatus
	}


	selected = "SELECT id, id_setting_grup, full_name, email, username, name_grup, status, image, extension"
	if search != "" {
		ors := " FROM v_get_user WHERE concat(full_name, email, username, name_grup) LIKE '%" + search + "%' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "Y" {
		ors := " FROM v_get_user WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "N" {
		ors := " FROM v_get_user WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else {
		whrs = " FROM v_get_user"
	}

	rows, err := db.Raw(selected + whrs + " ORDER BY full_name ASC").Rows()
	if err != nil {
		logs.Println(err)
		return c.Render(500, "error_500", nil)
	}
	defer rows.Close()

	each 	:= models.ModelUser{}
	result  := []models.ModelUser{}

	for rows.Next() {
		var id , id_setting_grup, full_name, email, username, name_grup, status, image, extension []byte

		err = rows.Scan(&id, &id_setting_grup, &full_name, &email, &username, &name_grup, &status, &image, &extension)
		if err != nil {
			logs.Println(err)
			return c.Render(500, "error_500", nil)
		}

		each.ID 				= ConvertToMD5(string(id))
		each.Id_setting_grup 	= string(id_setting_grup)
		each.Full_name 			= string(full_name)
		each.Email 				= string(email)
		each.Username 			= string(username)
		each.Name_Grup 			= string(name_grup)
		each.Status 			= string(status)
		each.Image 				= string(image)
		each.Additional 		= string(extension)

		result = append(result, each)
	}

	postsPerPage := 10
	paginator 	 = pagination.NewPaginator(c.Request(), postsPerPage, len(result))

	// fetch the next posts "postsPerPage"
	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)
	mydatas := []models.ModelUser{}
	for _, num := range idrange {
		if num <= len(result)-1 {
			numdata := result[num]
			mydatas = append(mydatas, numdata)
		}
	}

	data := response_json{
		"paginator" 	: paginator,
		"data"  		: mydatas,
		"search" 		: search,
		"searchStatus"  : searchStatus,
	}

	return c.Render(200, "list_setting_user", data)
}

func AddSettingUser(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.user_1") == false{
		return c.Render(403, "error_403", nil)
	}

	data_grup := GetDataGrup()

	data := response_json{
		"data_grup" : data_grup,
	}

	return c.Render(200, "add_setting_user", data)
}

// == Manipulate
func StoreSettingUser(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()
	
	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.user_1") == false{
		return c.Render(403, "error_403", nil)
	}

	field := new(models.ModelUser)
	if err := c.Bind(field); err != nil {
		return err
	}

	image 	        := FormFile(c, "image")

	// Proses Upload Gambar Single
	if image != "nil"{
		// Path_Destinasi_file
		path 	   := "upload_api/profile_user/"+ strconv.Itoa(data_users.Id_user)
		folderPath := MakeDirectory(path)
		if folderPath == "0" {
			logs.Println("error_create_directory")
		}
		dst, err := os.Create(folderPath + "/" + image)
		if err != nil {
			logs.Println(err)
		}
		defer dst.Close()
		// Eksekusi Simpan FIle
		file, _ 	   := c.FormFile("image")
		file_image, _  := file.Open()
		defer file_image.Close()		
		if _, err = io.Copy(dst, file_image); err != nil {
			logs.Println(err)
		}
	}
	// insert_user
	user := models.SettingUser{
		Full_name 	: field.Full_name,
		Username 	: field.Username,
		Email 		: field.Email,
		Telephone 	: field.Telephone,
		Address 	: field.Address,
		Gender 		: field.Gender,
		Password 	: HashPassword(field.Password),
		Status 		: field.Status,
		Image 		: image,
		CreatedAt 	: current_time("2006-01-02 15:04:05"),
	}
	if error_insert := db.Create(&user); error_insert.Error != nil {
		logs.Println(error_insert)
		return c.Render(500, "error_500", nil)
	}
	db.NewRecord(user)

	// Insert Group 
	group := models.SettingUserGrup{
		Id_setting_user 	: user.ID,
		Id_setting_grup 	: ConvertStringToInt(field.Id_setting_grup),
		Status 				: field.Status,
		CreatedAt 			: current_time("2006-01-02 15:04:05"),
	}
	if error_insert := db.Create(&group); error_insert.Error != nil {
		logs.Println(error_insert)
		return c.Render(500, "error_500", nil)
	}
	db.NewRecord(group)


	return c.Redirect(301, "/lib/setting/user/")
}