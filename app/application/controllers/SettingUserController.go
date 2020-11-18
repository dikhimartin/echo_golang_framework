package controllers

import (	
	"os"
	"io"
	"time"
	"strconv"
	"receipt/application/models"
	"receipt/database"
	"github.com/labstack/echo"
	"github.com/dikhimartin/beego-v1.12.0/utils/pagination"
)

// == Custom Function
func GetDataUserById(id_user string /*convert_to_md5*/) (models.ModelUser){
	db := database.CreateCon()
	defer db.Close()

	var id, id_setting_grup, full_name, gender, email, telephone, address,  username, name_grup, status, image []byte
	row := db.Table("v_get_user").Where("md5(id) = ?", id_user).Select("id, id_setting_grup, full_name, gender, email, telephone, address,  username, name_grup, status, image").Row() 
	err := row.Scan(&id, &id_setting_grup, &full_name, &gender, &email, &telephone, &address,  &username, &name_grup, &status, &image)
	if err != nil{
		logs.Println(err)
	}

	data := models.ModelUser{
		ID 					:  string(id),
		Id_setting_grup 	:  string(id_setting_grup),
		Full_name 			:  string(full_name),
		Gender 				:  string(gender),
		Email 				:  string(email),
		Telephone 			:  string(telephone),
		Address				:  string(address),
		Username 			:  string(username),
		Name_Grup 			:  string(name_grup),
		Status 				:  string(status),
		Image 				:  string(image),
		Additional 			:  id_user,
	}
	return data
}

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

func EditSettingUser(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.user_1") == false{
		return c.Render(403, "error_403", nil)
	}

	requested_id := c.Param("id")

	data_user := GetDataUserById(requested_id)

	data_grup := GetDataGrup()

	// // get_image
	path_image_profile := ""
	if data_user.Image == ""{
		path_image_profile = "/static/images/users/anonymous.png" 
	}else{
		path_image_profile = "/upload/profile_user/"+ data_user.Image +"?="+ strconv.FormatInt(time.Now().UnixNano(), 10) +" "
	}

	data := response_json{
		"data_user"	        	: data_user,
		"data_grup"  		    : data_grup,
		"path_image_profile"  	: path_image_profile,
	}

	return c.Render(200, "edit_setting_user", data)
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
		logs.Println(err)
		return c.Render(500, "error_500", nil)
	}


	// Proses Upload Gambar Single
	image 	        := FormFile(c, "image")
	if image != "nil"{
		// Path_Destinasi_file
		path 	   := "upload/profile_user/"
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
		image = ""
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

func UpdateSettingUser(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.user_1") == false{
		return c.Render(403, "error_403", nil)
	}

	requested_id 		    := c.Param("id")

	id_setting_grup 		:= ConvertStringToInt(c.FormValue("id_setting_grup"))
	full_name 				:= c.FormValue("full_name")
	email 					:= c.FormValue("email")
	telephone 				:= c.FormValue("telephone")
	address 				:= c.FormValue("address")
	gender 					:= c.FormValue("gender")
	username    			:= c.FormValue("username")
	status      			:= c.FormValue("status")
	password_val    		:= c.FormValue("password")
	confirm_password_val    := c.FormValue("confirm_password")

	user := GetDataUserById(requested_id)

	// reupload_image
	image 	      := FormFile(c, "image")
	if image != "nil"{
		path 	   := "upload/profile_user/"
		if user.Image != ""{
			remove_file := RemoveFile(c, path + "/" + user.Image)
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
		image = user.Image
	}

	// Update Data User
	if password_val == "" && confirm_password_val == ""{
		var update models.SettingUser
		update_user := db.Model(&update).Where("md5(id) = ?", requested_id).Updates(map[string]interface{}{
			"full_name"    :    full_name,
			"email"    	   :    email,
			"telephone"    :    telephone,
			"address"      :    address,
			"gender"       :    gender,
			"username"     :    username,
			"image"        :    image,
			"status"       :    status,
			"updated_at"   :    current_time("2006-01-02 15:04:05"),
		})
		if update_user.Error != nil {
			logs.Println(update_user.Error)
			return c.Render(500, "error_500", nil)
		}
	}else if password_val != "" && confirm_password_val != "" {
		var update models.SettingUser
		update_user := db.Model(&update).Where("md5(id) = ?", requested_id).Updates(map[string]interface{}{
			"full_name"    :    full_name,
			"email"    	   :    email,
			"telephone"    :    telephone,
			"address"      :    address,
			"gender"       :    gender,
			"username"     :    username,
			"image"        :    image,
			"status"       :    status,
			"password"     :    HashPassword(password_val),
			"updated_at"   :    current_time("2006-01-02 15:04:05"),
		})
		if update_user.Error != nil {
			logs.Println(update_user.Error)
			return c.Render(500, "error_500", nil)
		}
	}

	// update user group
	var user_group models.SettingUserGrup
	update_user_group := db.Model(&user_group).Where("md5(id_setting_user) = ?", requested_id).Updates(map[string]interface{}{
		"id_setting_grup"    :    id_setting_grup,
		"updated_at"   		 :    current_time("2006-01-02 15:04:05"),
	})
	if update_user_group.Error != nil {
		logs.Println(update_user_group.Error)
		return c.Render(500, "error_500", nil)
	}

	return c.Redirect(301, "/lib/setting/user/")
}