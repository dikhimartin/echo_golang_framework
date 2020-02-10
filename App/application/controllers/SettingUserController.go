package controllers
import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"
	"fmt"
	"os"
	"io"
	"path/filepath"
	_ "database/sql"
	"../../database"
	"../models"
	"github.com/labstack/echo"
	"github.com/astaxie/beego/utils/pagination"
    "github.com/disintegration/imaging"
	"github.com/flosch/pongo2"
)

func ListSettingUser(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.user_2'")
	if errPriv != nil {
		logs.Println(errPriv)
		return c.Render(http.StatusInternalServerError, "error_505", nil)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.user_2"{
		logs.Println(errPriv)
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	var selected 	 string
	var whrs 		 string
	var search 		 string
	var searchStatus string

	if reqSearch := c.FormValue("search"); reqSearch != "" {
		search = reqSearch
	}
	if reqSearchStatus := c.FormValue("searchStatus"); reqSearchStatus != "" {
		searchStatus = reqSearchStatus
	}



	selected = "SELECT id, id_setting_grup, full_name, email, username, name_grup, status, image, extension"
	// search 
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

	rows, err := db.Query(selected + whrs + " ORDER BY full_name ASC")
	if err != nil {
		logs.Println(err)
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	defer rows.Close()


	each 	:= models.SettingUser{}
	result  := []models.SettingUser{}

	for rows.Next() {
		var id string
		var id_setting_grup, full_name, email, username, name_grup, status, image, extension []byte

		var err = rows.Scan(&id, &id_setting_grup, &full_name, &email, &username, &name_grup, &status, &image, &extension)

		if err != nil {
			logs.Println(err)
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

		var str string = id
		hasher 		:= md5.New()
		hasher.Write([]byte(str))
		converId 	:= hex.EncodeToString(hasher.Sum(nil))

		each.Id 		= converId
		each.Id_group 	= string(id_setting_grup)
		each.Full_name 	= string(full_name)
		each.Email 		= string(email)
		each.Username 	= string(username)
		each.Name_grup 	= string(name_grup)
		each.Status 	= string(status)
		each.Image 		= string(image)
		each.Additional = string(extension)

		result = append(result, each)
	}


	// Lets use the Forbes top 7.
	mydata := result

	// sets paginator with the current offset (from the url query param)
	postsPerPage := 10
	paginator = pagination.NewPaginator(c.Request(), postsPerPage, len(mydata))

	// fmt.Println(paginator.Offset())
	// fetch the next posts "postsPerPage"
	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)
	//create a new page list that shows up on html
	mydatas := []models.SettingUser{}
	for _, num := range idrange {
		//Prevent index out of range errors
		if num <= len(mydata)-1 {
			numdata := mydata[num]
			mydatas = append(mydatas, numdata)
		}
	}

	// set the paginator in context
	// also set the page list in context
	// if you also have more data, set it context
	data = pongo2.Context{
		"paginator":    paginator,
		"setting_user": mydatas,
		"search":       search,
		"searchStatus": searchStatus}

	return c.Render(http.StatusOK, "list_setting_user", data)
}

func AddSettingUser(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.user_1'")
	if errPriv != nil {
		logs.Println(errPriv)
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.user_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	resultGrup	:= cc.getDataGrup()

	data = pongo2.Context{
		"resultGrup"        : resultGrup,
	}

	return c.Render(http.StatusOK, "add_setting_user", data)
}

func StoreSettingUser(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.user_1'")
	if errPriv != nil {
		logs.Println(errPriv)
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.user_1"{
		logs.Println(errPriv)
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege


	// get_current_time
	t 				:= time.Now()
	current_time 	:= t.Format("2006-01-02")

	emp := new(models.SettingUser)
	if err := c.Bind(emp); err != nil {
		return err
	}

	// SALT
	jam 	:= strconv.Itoa(now.Hour())
	menit 	:= strconv.Itoa(now.Minute())
	detik 	:= strconv.Itoa(now.Second())
	var str_salt string = jam + menit + detik
	hasher_salt := md5.New()
	hasher_salt.Write([]byte(str_salt))
	salt 		:= hex.EncodeToString(hasher_salt.Sum(nil))


	// MD5 PASSWORD
	var str_password string = emp.Password
	hasher_password := md5.New()
	hasher_password.Write([]byte(str_password))
	md5password 	:= hex.EncodeToString(hasher_password.Sum(nil))

	// PASSWORD FINAL
	var salt_password string = salt + md5password
	hasher_salt_password := md5.New()
	hasher_salt_password.Write([]byte(salt_password))
	password 			:= hex.EncodeToString(hasher_salt_password.Sum(nil))


	// Define File to upload
	file, err := c.FormFile("image")
	if err != nil {
		errorEntryData = "error_image_null"
		SaveSettingUserWithoutImage(c)
		return c.Redirect(301, "/lib/setting/user/")
	}
	src, err := file.Open()
	if err != nil {
		errorEntryData = "error_image_null"
		SaveSettingUserWithoutImage(c)
		return c.Redirect(301, "/lib/setting/user/")

	}
	defer src.Close()
	// end Define File to upload


	var FileNamePost string
	var defaultname = "martin_profile_"
	var extension 	= filepath.Ext(file.Filename)

	// Insert user 
	sql := "INSERT INTO tb_setting_user(full_name, username, email, telephone, address, gender, password, salt, add_date, status, additional) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	insert_user, err := db.Prepare(sql)
	if err != nil {
		logs.Println(err)
		return c.Render(http.StatusInternalServerError, "error_505", nil)	
	}
	defer insert_user.Close()
	res, err2 := insert_user.Exec(emp.Full_name, emp.Username, emp.Email, emp.Telephone, emp.Address, emp.Gender, password, salt, current_time, emp.Status, extension)

	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	sum_last_id, _ := res.LastInsertId()
	get_last_id := strconv.FormatInt(int64(sum_last_id), 10)


	// Insert Group 
	sql2 := "INSERT INTO tb_setting_user_grup(id_setting_user, id_setting_grup, status, add_date) VALUES(?, ?, ?, ?)"
	insert_group, err := db.Prepare(sql2)
	if err != nil {
		logs.Println(err)
		return c.Render(http.StatusInternalServerError, "error_505", nil)	
	}
	defer insert_group.Close()
	_, err3 := insert_group.Exec(get_last_id, emp.Id_group, emp.Status, current_time)
	if err3 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}


	// UPLOAD FILE

	// hash_name_file
		hasher := md5.New()
		hasher.Write([]byte(defaultname + get_last_id))
		EncryptName := hex.EncodeToString(hasher.Sum(nil))
		FileNamePost = EncryptName + extension

	// Path_Destinasi_file
		dst, err := os.Create("upload/profile_user/" + FileNamePost)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}
		defer dst.Close()

	// Eksekusi Simpan FIle
		if _, err = io.Copy(dst, src); err != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

	// Memotong Gambar
         // load original image
         img, err := imaging.Open("./upload/profile_user/"+FileNamePost)
         if err != nil {
			logs.Println(err)
			return c.Render(http.StatusInternalServerError, "error_505", nil)	
            os.Exit(1)
         }

         // crop from center
         centercropimg := imaging.CropCenter(img, 300, 300)

         // save cropped image
         err = imaging.Save(centercropimg, "./upload/profile_user/"+FileNamePost)
         if err != nil {
                 return c.Render(http.StatusInternalServerError, "error_500", nil)
                 os.Exit(1)
         }

	// Simpan nama file ke database
		insert_nama_image, err2 := db.Prepare("UPDATE tb_setting_user SET image = ? WHERE id = ?")
		if err2 != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}
		defer insert_nama_image.Close()
		insert_nama_image.Exec(EncryptName, get_last_id)


	return c.Redirect(301, "/lib/setting/user/")
}

func SaveSettingUserWithoutImage(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.user_1'")
	if errPriv != nil {
		logs.Println(errPriv)
		return c.Render(http.StatusInternalServerError, "error_505", nil)	
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.user_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	// get_current_time
	t 				:= time.Now()
	current_time 	:= t.Format("2006-01-02")

	emp := new(models.SettingUser)
	if err := c.Bind(emp); err != nil {
		return err
	}

	// SALT
	jam 	:= strconv.Itoa(now.Hour())
	menit 	:= strconv.Itoa(now.Minute())
	detik 	:= strconv.Itoa(now.Second())
	var str_salt string = jam + menit + detik
	hasher_salt := md5.New()
	hasher_salt.Write([]byte(str_salt))
	salt 		:= hex.EncodeToString(hasher_salt.Sum(nil))


	// MD5 PASSWORD
	var str_password string = emp.Password
	hasher_password := md5.New()
	hasher_password.Write([]byte(str_password))
	md5password 	:= hex.EncodeToString(hasher_password.Sum(nil))

	// PASSWORD FINAL
	var salt_password string = salt + md5password
	hasher_salt_password := md5.New()
	hasher_salt_password.Write([]byte(salt_password))
	password 			:= hex.EncodeToString(hasher_salt_password.Sum(nil))

	// Insert user 
	sql := "INSERT INTO tb_setting_user(full_name, username, email, telephone, address, gender, password, salt, add_date, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	insert_user, err := db.Prepare(sql)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	defer insert_user.Close()
	res, err2 := insert_user.Exec(emp.Full_name, emp.Username, emp.Email, emp.Telephone, emp.Address,  emp.Gender, password, salt, current_time, emp.Status)

	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	sum_last_id, _ := res.LastInsertId()
	get_last_id := strconv.FormatInt(int64(sum_last_id), 10)


	// Insert Group 
	sql2 := "INSERT INTO tb_setting_user_grup(id_setting_user, id_setting_grup, status, add_date) VALUES(?, ?, ?, ?)"
	insert_group, err := db.Prepare(sql2)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	defer insert_group.Close()
	_, err3 := insert_group.Exec(get_last_id, emp.Id_group, emp.Status, current_time)
	if err3 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	return c.Render(http.StatusOK, "list_setting_user", data)
}

func EditSettingUser(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.user_3'")
	if errPriv != nil {
		logs.Println(errPriv)
		return c.Render(http.StatusInternalServerError, "error_505", nil)	
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.user_3"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	requested_id := c.Param("id")

	var id_setting_grup, full_name, gender, email, telephone, address, username, name_grup, status, image, extension []byte
	err := db.QueryRow("SELECT id_setting_grup, full_name, gender, email, telephone, address,  username, name_grup, status, image, extension FROM v_get_user WHERE md5(id) = ?", requested_id).Scan(&id_setting_grup, &full_name, &gender, &email, &telephone, &address, &username, &name_grup, &status, &image, &extension)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	errorUpdate := ""
	if errorDuplikatData == "error_update_username" {
		errorUpdate 	  = "Sorry , Username is already used"
		errorDuplikatData = ""
	}

	// get_image
	t   			   := time.Now().UnixNano()
	current_time 	   := strconv.FormatInt(t, 10)
	Path_image_profile := ""
	if string(image) == ""{
		Path_image_profile = "/static/images/users/anonymous.png" 
	}else{
		Path_image_profile = "/upload/profile_user/"+ string(image)+string(extension)+"?="+ current_time +" "
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

	resultGrup			:= cc.getDataGrup()

	data = pongo2.Context{
		"error"			        : errorUpdate,
		"setting_user"	        : response,
		"resultGrup"  		    : resultGrup,
		"Path_image_profile"  	: Path_image_profile,
	}

	return c.Render(http.StatusOK, "edit_setting_user", data)
}

func UpdateSettingUser(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.user_3'")
	if errPriv != nil {
		logs.Println(errPriv)
		return c.Render(http.StatusInternalServerError, "error_505", nil)	
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.user_3"{
		logs.Println(errPriv)
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	// end check_privilege
	
	requested_id 		    := c.Param("id")
	id_group 				:= c.FormValue("id_group")
	full_name 				:= c.FormValue("full_name")
	email 					:= c.FormValue("email")
	telephone 				:= c.FormValue("telephone")
	address 				:= c.FormValue("address")
	gender 					:= c.FormValue("gender")
	username    			:= c.FormValue("username")
	status      			:= c.FormValue("status")
	password_val    		:= c.FormValue("password")
	confirm_password_val    := c.FormValue("confirm_password")
	remove_image    		:= c.FormValue("remove_image")

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

	t 				:= time.Now()
	current_time 	:= t.Format("2006-01-02")


	// Proses Simpan gambar
	file, image_check := c.FormFile("image")

	// if file null
	if image_check != nil {
		fmt.Println("Image Empty")
		// remove image
		if remove_image == "1"{
			// get_data_image
			var image_user, additional []byte
			get_images, err := db.Prepare("SELECT image, additional FROM tb_setting_user WHERE md5(id) = ?")
			if err != nil {
				logs.Println(err)
				return c.Render(http.StatusInternalServerError, "error_500", nil)			
			}
			err = get_images.QueryRow(requested_id).Scan(&image_user, &additional)	
			defer get_images.Close()

			// update_data
			Update_image, err := db.Prepare("UPDATE tb_setting_user SET image=?, additional=? WHERE md5(id)=?")
			if err != nil{
				return c.Render(http.StatusInternalServerError, "error_500", nil)
			}
			_, err2 :=Update_image.Exec(nil, nil, requested_id)
			if err2 != nil {
				return EditSettingUser(c)
			}
			defer Update_image.Close()

			// remove_file
			err = os.Remove("upload/profile_user/" + string(image_user) + string(additional))
			if err != nil {
				return c.Render(http.StatusInternalServerError, "error_500", nil)
			}
		}
		// Update Data User
		if password_val == "" && confirm_password_val == ""{

			UpdateUser, err := db.Prepare("UPDATE tb_setting_user SET full_name=?, email=?, telephone = ?, address = ?, gender = ?, username = ?, status = ?, update_date=? WHERE md5(id)=?")
			if err != nil{
				logs.Println(err)
				return c.Render(http.StatusInternalServerError, "error_500", nil)			
			}
			_, err2 :=UpdateUser.Exec(full_name, email, telephone, address,  gender, username, status, current_time, requested_id)
			if err2 != nil {
				errorDuplikatData = "error_update_username"
				return EditSettingUser(c)
			}
			defer UpdateUser.Close()
		}else if password_val != "" && confirm_password_val != "" {

			UpdateUser, err := db.Prepare("UPDATE tb_setting_user SET full_name=?, email=?, telephone = ?, address = ?, gender = ?, username = ?, password=?, salt = ?,  status=?, update_date=? WHERE md5(id)=?")
			if err != nil{
				logs.Println(err)
				return c.Render(http.StatusInternalServerError, "error_500", nil)			
			}
			_, err2 :=UpdateUser.Exec(full_name, email, telephone, address, gender, username, password, salt, status, current_time, requested_id)
			if err2 != nil {
				logs.Println(err)
				errorDuplikatData = "error_update_username"
				return EditSettingUser(c)
			}
			defer UpdateUser.Close()
		}

		// update group
		UpdateUserGroup, err := db.Prepare("UPDATE tb_setting_user_grup SET id_setting_grup=?, update_date=? WHERE md5(id_setting_user)=?")
		if err != nil{
			logs.Println(err)
			return c.Render(http.StatusInternalServerError, "error_500", nil)			
		}
		UpdateUserGroup.Exec(id_group, current_time, requested_id)
		defer UpdateUserGroup.Close()

		return c.Redirect(301, "/lib/setting/user/")

	// if file not null
	}else{

		fmt.Println("Image Not Empty")

		//if file not null
		src, err := file.Open()
		if err != nil {
			logs.Println(err)
			return c.Render(http.StatusInternalServerError, "error_500", nil)			
		}
		defer src.Close()

		var FileNamePost string
		var defaultname = "martin_profile_"
		var extension 	= filepath.Ext(file.Filename)

		//gets id_user
		var id_user, image_user, additional []byte
		err = db.QueryRow("SELECT id, image, extension FROM v_get_user WHERE md5(id) = ?", requested_id).Scan(&id_user, &image_user, &additional)
		if err != nil {
			logs.Println(err)
			return c.Render(http.StatusInternalServerError, "error_500", nil)			
		}
		// Update Data User
		if password_val == "" && confirm_password_val == ""{

			UpdateUser, err := db.Prepare("UPDATE tb_setting_user SET full_name=?, email=?, telephone = ?, address = ?, gender = ?, username = ?, status = ?, update_date=?, additional = ? WHERE md5(id)=?")
			if err != nil{
				logs.Println(err)
				return c.Render(http.StatusInternalServerError, "error_500", nil)			
			}
			_, err2 :=UpdateUser.Exec(full_name, email, telephone, address, gender, username, status, current_time, extension, requested_id)
			if err2 != nil {
				errorDuplikatData = "error_update_username"
				return EditSettingUser(c)
			}
			defer UpdateUser.Close()
		}else if password_val != "" && confirm_password_val != "" {

			UpdateUser, err := db.Prepare("UPDATE tb_setting_user SET full_name=?, email=?, telephone = ?, address = ?, gender = ?, username = ?, password=?, salt = ?,  status=?, update_date=?, additional = ? WHERE md5(id)=?")
			if err != nil{
				logs.Println(err)
				return c.Render(http.StatusInternalServerError, "error_500", nil)			
			}
			_, err2 :=UpdateUser.Exec(full_name, email, telephone, address, gender, username, password, salt, status, current_time, extension, requested_id)
			if err2 != nil {
				logs.Println(err)
				errorDuplikatData = "error_update_username"
				return EditSettingUser(c)
			}
			defer UpdateUser.Close()
		}

		// update group
		UpdateUserGroup, err := db.Prepare("UPDATE tb_setting_user_grup SET id_setting_grup=?, update_date=? WHERE md5(id_setting_user)=?")
		if err != nil{
			logs.Println(err)
			return c.Render(http.StatusInternalServerError, "error_500", nil)			
		}
		UpdateUserGroup.Exec(id_group, current_time, requested_id)
		defer UpdateUserGroup.Close()


		// HASH TO MD5
		hasher := md5.New()
		hasher.Write([]byte(defaultname + string(id_user)))
		EncryptName := hex.EncodeToString(hasher.Sum(nil))
		FileNamePost = EncryptName + extension

		//remove file before update
		if string(image_user) != ""{
			err = os.Remove("upload/profile_user/" + string(image_user) + string(additional))
			if err != nil {
				logs.Println(err)
				return c.Render(http.StatusInternalServerError, "error_500", nil)			
			}
		}

		// Lokasi File
		dst, err := os.Create("upload/profile_user/" + FileNamePost)
		if err != nil {
			logs.Println(err)
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}
		defer dst.Close()

		// Eksekusi File
		if _, err = io.Copy(dst, src); err != nil {
			logs.Println(err)
			return c.Render(http.StatusInternalServerError, "error_500", nil)		
		}

		// Memotong Gambar
		     // load original image
		     img, err := imaging.Open("./upload/profile_user/"+FileNamePost)
		     if err != nil {
					 logs.Println(err)
					 return c.Render(http.StatusInternalServerError, "error_500", nil)
		             os.Exit(1)
		     }

		     // crop from center
		     centercropimg := imaging.CropCenter(img, 300, 300)

		     // save cropped image
		     err = imaging.Save(centercropimg, "./upload/profile_user/"+FileNamePost)
		     if err != nil {
				logs.Println(err)
				return c.Render(http.StatusInternalServerError, "error_500", nil)		             
				os.Exit(1)
		     }
		// Memotong Gambar

		// Simpan nama file ke database
		insert_nama_image, err2 := db.Prepare("UPDATE tb_setting_user SET image = ? WHERE md5(id) = ?")
		if err2 != nil {
			logs.Println(err)
			return c.Render(http.StatusInternalServerError, "error_500", nil)			
		}
		defer insert_nama_image.Close()
		insert_nama_image.Exec(EncryptName, requested_id)

	}

	return c.Redirect(301, "/lib/setting/user/")
}

func CheckUsername(c echo.Context) error{

	db := database.CreateCon()
	defer db.Close()

	var check_username []byte
	type_check     := c.FormValue("type_check")
	username   	   := c.FormValue("username")
	username_old   := c.FormValue("username_old")

	if type_check == "1"{

		//cek_data_add_new
		query, err := db.Prepare("SELECT username FROM v_get_user WHERE username= ?")
		if err != nil {
			logs.Println(err)
		}
		defer query.Close()
		err = query.QueryRow(username).Scan(&check_username)

	}else if type_check == "2"{

		//cek_data_edit
		query, err := db.Prepare("SELECT username FROM v_get_user WHERE username= ? AND username != ?")
		if err != nil {
			logs.Println(err)
		}
		defer query.Close()
		err = query.QueryRow(username, username_old).Scan(&check_username)
	}

	if string(check_username) != "" {
		response := map[string]string{"alert": "Maaf, Username sudah digunakan!", "kode": "1"}
		return c.JSON(http.StatusOK, response)
	} else {
		response := map[string]string{"alert": "Sukses, Username belum digunakan!", "kode": "0"}
		return c.JSON(http.StatusOK, response)
	}	
}
