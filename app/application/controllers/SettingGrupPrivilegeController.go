package controllers

import (
	"../models"
	"../../database"
	"github.com/labstack/echo"
	"github.com/dikhimartin/beego-v1.12.0/utils/pagination"
)

// == Custom Function
func GetDataGrupPrivilegeById(id_grup_privilege string /*convert_to_md5*/) (models.ModelGrupPrivilege){
	db := database.CreateCon()
	defer db.Close()

	var id, id_setting_grup, name_grup, status, remarks []byte
	row := db.Table("v_get_grup_privilege").Where("md5(id) = ?", id_grup_privilege).Select("id, id_setting_grup, name_grup, status, remarks").Row() 
	err := row.Scan(&id, &id_setting_grup, &name_grup, &status, &remarks)
	if err != nil{
		logs.Println(err)
	}
	data := models.ModelGrupPrivilege{
		ID        	     : string(id),
		Id_setting_grup  : string(id_setting_grup),
		Name_grup  		 : string(name_grup),
		Status    	     : string(status),
		Remarks     	 : string(remarks),
		Additional     	 : ConvertToMD5(string(id)),
	}
	return data
}

func CheckGrupPrivilege() ([]models.ModelGrup) {
	db := database.CreateCon()
	defer db.Close()

	rows, err := db.Raw("SELECT id, name_grup, status FROM v_get_grup WHERE status = 'Y' AND id_setting_grup_privilege IS NULL ORDER BY name_grup").Rows();
	if err != nil {
		logs.Println(err)
	}
	defer rows.Close()
	each   := models.ModelGrup{}
	result := []models.ModelGrup{}

	for rows.Next() {
		var id, name_grup, status []byte
		err = rows.Scan(&id, &name_grup, &status)
		if err != nil {
			logs.Println(err)
		}

		each.ID 			 = string(id)
		each.Name_Grup 		 = string(name_grup)
		each.Status  		 = string(status)

		result = append(result, each)
	}

	return result
}

// == View
func ListSettingGrupPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grupprivilege_2") == false{
		return c.Render(403, "error_403", nil)
	}

	var selected string
	var whrs string
	var search string
	var searchStatus string

	if reqSearch := c.FormValue("search"); reqSearch != "" {
		search = reqSearch
	}

	if reqSearchStatus := c.FormValue("searchStatus"); reqSearchStatus != "" {
		searchStatus = reqSearchStatus
	}

	selected = "SELECT id, id_setting_grup, name_grup, status, created_at, updated_at"
	if search != "" {
		ors := " FROM v_get_grup_privilege WHERE name_grup LIKE '%" + search + "%' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "Y" {
		ors := " FROM v_get_grup_privilege WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "N" {
		ors := " FROM v_get_grup_privilege WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else {
		whrs = " FROM v_get_grup_privilege"
	}

	rows, err := db.Raw(selected + whrs + " ORDER BY name_grup ASC").Rows()
	if err != nil {
		logs.Println(err)
		return c.Render(500, "error_500", nil)
	}

	defer rows.Close()

	each   := models.ModelGrupPrivilege{}
	result := []models.ModelGrupPrivilege{}

	for rows.Next() {
		var	id, id_setting_grup, name_grup, status, created_at, updated_at[]byte

		var err = rows.Scan(&id, &id_setting_grup, &name_grup, &status, &created_at, &updated_at)
		if err != nil {
			logs.Println(err)
			return c.Render(500, "error_500", nil)
		}

		each.ID  			 = ConvertToMD5(string(id))
		each.Id_setting_grup = string(id_setting_grup)
		each.Name_grup  	 = string(name_grup)
		each.Status  		 = string(status)
		each.CreatedAt 	 	 = FormatDate(string(created_at), "02 January 2006 at 15:04 PM")
		each.UpdatedAt 	 	 = FormatDate(string(updated_at), "02 January 2006 at 15:04 PM")

		result = append(result, each)
	}

	postsPerPage := 10
	paginator 	 = pagination.NewPaginator(c.Request(), postsPerPage, len(result))

	// fetch the next posts "postsPerPage"
	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)
	mydatas := []models.ModelGrupPrivilege{}
	for _, num := range idrange {
		if num <= len(result)-1 {
			numdata := result[num]
			mydatas = append(mydatas, numdata)
		}
	}

	data := response_json{
		"paginator"		: paginator,
		"data"			: mydatas,
		"search"		: search,
		"searchStatus"	: searchStatus,
	}

	return c.Render(200, "list_setting_grup_privilege", data)
}

func AddSettingGrupPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grupprivilege_2") == false{
		return c.Render(403, "error_403", nil)
	}

	rowsPrivilege, errPrivilege := db.Raw("SELECT id_setting_privilege, name_menu,  code_privilege, code_permissions, permissions FROM v_get_privilege GROUP BY id_setting_privilege ORDER BY code_permissions ASC").Rows()
	if errPrivilege != nil {
		logs.Println(errPrivilege)
		return c.Render(500, "error_500", nil)
	}
	defer rowsPrivilege.Close()

	eachPrivilege   := models.ModelPrivilege{}
	resultPrivilege := []models.ModelPrivilege{}

	for rowsPrivilege.Next() {
		var id_setting_privilege, 
			name_menu, 
			code_privilege,
			code_permissions,
			permissions []byte

			var err = rowsPrivilege.Scan(&id_setting_privilege, &name_menu, &code_privilege, &code_permissions, &permissions)
			if err != nil {
				logs.Println(errPrivilege)
				return c.Render(500, "error_500", nil)
			}

			// get_permission
			rows, err := db.Raw("SELECT permissions, code_permissions FROM v_get_privilege WHERE id_setting_privilege = ?", string(id_setting_privilege)).Rows()
			if err != nil {
				logs.Println(err)
				return c.Render(500, "error_500", nil)
			}
			defer rows.Close()

			eachPermission   := models.ModelPermission{}
			resultPermission := []models.ModelPermission{}

			for rows.Next() {
				var permissions, code_permissions []byte
				err = rows.Scan(&permissions, &code_permissions)
				if err != nil {
					logs.Println(err)
					return c.Render(500, "error_500", nil)
				}

				permission := ""
				if string(code_permissions) == string(code_privilege)+"_1"{
					permission = "Create"
				}else if string(code_permissions) == string(code_privilege)+"_2"{
					permission = "Read/View"
				}else if string(code_permissions) == string(code_privilege)+"_3"{
					permission = "Edit"
				}else if string(code_permissions) == string(code_privilege)+"_4"{
					permission = "Delete"
				}

				eachPermission.ID 		   = string(permissions)
				eachPermission.Name  	   = permission
				eachPermission.Additional  = string(code_permissions)

				resultPermission = append(resultPermission, eachPermission)
			}

		eachPrivilege.ID 	 			= string(id_setting_privilege)
		eachPrivilege.Name_menu  		= string(name_menu)
		eachPrivilege.Code_privilege 	= string(code_privilege)
		eachPrivilege.Code_permission 	= string(code_permissions)
		eachPrivilege.Permissions 		= resultPermission

		resultPrivilege = append(resultPrivilege, eachPrivilege)
	}

	user_grup := CheckGrupPrivilege()

	data := response_json{
		"user_grup"		  : user_grup,
		"resultPrivilege" : resultPrivilege,
	}	

	return c.Render(200, "add_setting_grup_privilege", data)
}

func EditSettingGrupPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grupprivilege_3") == false{
		return c.Render(403, "error_403", nil)
	}

	requested_id := c.Param("id")

	// get_data_grup_privilege
	data_grup_privilege := GetDataGrupPrivilegeById(requested_id)

	// get_data_grup
	data_grup := GetDataGrup()


	// get permissions menu
	rowsPrivilege, errPrivilege := db.Raw("SELECT id_setting_privilege, name_menu,  code_privilege, code_permissions, permissions FROM v_get_privilege GROUP BY id_setting_privilege ORDER BY code_permissions ASC").Rows()
	if errPrivilege != nil {
		logs.Println(errPrivilege)
	}
	defer rowsPrivilege.Close()

	eachPrivilege   := models.ModelPrivilege{}
	resultPrivilege := []models.ModelPrivilege{}
	for rowsPrivilege.Next() {
		var id_setting_privilege, 
			name_menu, 
			code_privilege,
			code_permissions,
			permissions []byte

				var err = rowsPrivilege.Scan(&id_setting_privilege, &name_menu, &code_privilege, &code_permissions, &permissions)
				if err != nil {
					logs.Println(errPrivilege)
				}

				// get_permission
				rows, err := db.Raw("SELECT permissions, code_permissions FROM v_get_privilege WHERE id_setting_privilege = ?", string(id_setting_privilege)).Rows()
				if err != nil {
					logs.Println(err)
				}
				defer rows.Close()

				eachPermission   := models.ModelPermission{}
				resultPermission := []models.ModelPermission{}

				for rows.Next() {

					var permissions, code_permissions []byte

					err = rows.Scan(&permissions, &code_permissions)
					if err != nil {
						logs.Println(err)
					}

					// check or unchecked
					var checked_or_unchecked []byte
					row := db.Table("tb_setting_grup_privilege_detail").Where("md5(id_setting_grup_privilege) = ? AND code_permissions = ?", requested_id, string(code_permissions)).Select("code_permissions").Row() 
					err := row.Scan(&checked_or_unchecked)
					if err != nil{
						logs.Println(err)
					}

					status_check := ""
					if string(checked_or_unchecked) == ""{
						status_check = "unchecked"
					}else if string(checked_or_unchecked) != ""{
						status_check = "checked"
					}

					permission := ""
					if string(code_permissions) == string(code_privilege)+"_1"{
						permission = "Create"
					}else if string(code_permissions) == string(code_privilege)+"_2"{
						permission = "Read/View"
					}else if string(code_permissions) == string(code_privilege)+"_3"{
						permission = "Edit"
					}else if string(code_permissions) == string(code_privilege)+"_4"{
						permission = "Delete"
					}

					eachPermission.ID 		       = string(permissions)
					eachPermission.Name  	       = permission
					eachPermission.Additional  	   = string(code_permissions)
					eachPermission.CheckOrUncheck  = status_check

					resultPermission = append(resultPermission, eachPermission)
				}

		eachPrivilege.ID 	 			= string(id_setting_privilege)
		eachPrivilege.Name_menu  		= string(name_menu)
		eachPrivilege.Code_privilege 	= string(code_privilege)
		eachPrivilege.Code_permission 	= string(code_permissions)
		eachPrivilege.Permissions 		= resultPermission

		resultPrivilege = append(resultPrivilege, eachPrivilege)
	}


	data := response_json{
		"data_grup"		  			: data_grup,
		"data_grup_privilege"		: data_grup_privilege,
		"resultPrivilege"			: resultPrivilege,
	}

	return c.Render(200, "edit_setting_grup_privilege", data)
}

func ShowSettingGrupPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grupprivilege_2") == false{
		return c.Render(403, "error_403", nil)
	}

	requested_id := c.Param("id")

	// get_data_grup_privilege
	data_grup_privilege := GetDataGrupPrivilegeById(requested_id)

	// get_data_grup
	data_grup := GetDataGrup()

	// get permissions menu
	rowsPrivilege, errPrivilege := db.Raw("SELECT id_setting_privilege, name_menu,  code_privilege, code_permissions, permissions FROM v_get_privilege GROUP BY id_setting_privilege ORDER BY code_permissions ASC").Rows()
	if errPrivilege != nil {
		logs.Println(errPrivilege)
	}
	defer rowsPrivilege.Close()
	eachPrivilege   := models.ModelPrivilege{}
	resultPrivilege := []models.ModelPrivilege{}
	for rowsPrivilege.Next() {
		var id_setting_privilege, 
			name_menu, 
			code_privilege,
			code_permissions,
			permissions []byte

				var err = rowsPrivilege.Scan(&id_setting_privilege, &name_menu, &code_privilege, &code_permissions, &permissions)
				if err != nil {
					logs.Println(errPrivilege)
				}

				// get_permission
				rows, err := db.Raw("SELECT permissions, code_permissions FROM v_get_privilege WHERE id_setting_privilege = ?", string(id_setting_privilege)).Rows()
				if err != nil {
					logs.Println(err)
				}
				defer rows.Close()

				eachPermission   := models.ModelPermission{}
				resultPermission := []models.ModelPermission{}

				for rows.Next() {

					var permissions, code_permissions []byte

					err = rows.Scan(&permissions, &code_permissions)
					if err != nil {
						logs.Println(err)
					}

					// check or unchecked
					var checked_or_unchecked []byte
					row := db.Table("tb_setting_grup_privilege_detail").Where("md5(id_setting_grup_privilege) = ? AND code_permissions = ?", requested_id, string(code_permissions)).Select("code_permissions").Row() 
					err := row.Scan(&checked_or_unchecked)
					if err != nil{
						logs.Println(err)
					}

					status_check := ""
					if string(checked_or_unchecked) == ""{
						status_check = "unchecked"
					}else if string(checked_or_unchecked) != ""{
						status_check = "checked"
					}

					permission := ""
					if string(code_permissions) == string(code_privilege)+"_1"{
						permission = "Create"
					}else if string(code_permissions) == string(code_privilege)+"_2"{
						permission = "Read/View"
					}else if string(code_permissions) == string(code_privilege)+"_3"{
						permission = "Edit"
					}else if string(code_permissions) == string(code_privilege)+"_4"{
						permission = "Delete"
					}

					eachPermission.ID 		       = string(permissions)
					eachPermission.Name  	       = permission
					eachPermission.Additional  	   = string(code_permissions)
					eachPermission.CheckOrUncheck  = status_check

					resultPermission = append(resultPermission, eachPermission)
				}

		eachPrivilege.ID 	 			= string(id_setting_privilege)
		eachPrivilege.Name_menu  		= string(name_menu)
		eachPrivilege.Code_privilege 	= string(code_privilege)
		eachPrivilege.Code_permission 	= string(code_permissions)
		eachPrivilege.Permissions 		= resultPermission

		resultPrivilege = append(resultPrivilege, eachPrivilege)
	}

	data := response_json{
		"data_grup"		  		: data_grup,
		"data_grup_privilege"   : data_grup_privilege,
		"resultPrivilege"		: resultPrivilege,
	}
	return c.Render(200, "show_setting_grup_privilege", data)
}


// == Manipulate
func StoreSettingGrupPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grupprivilege_2") == false{
		return c.Render(403, "error_403", nil)
	}


	id_setting_grup := ConvertStringToInt(c.FormValue("id_setting_grup"))
	remarks 		:= c.FormValue("remarks")
	status 			:= c.FormValue("status")

	// insert_grup_privilege
	grup_privilege := models.SettingGrupPrivilege{
		Id_setting_grup 	: id_setting_grup,
		Remarks 			: remarks,
		Status 				: status,
		CreatedAt 			: current_time("2006-01-02 15:04:05"),
	}
	if error_insert := db.Create(&grup_privilege); error_insert.Error != nil {
		logs.Println(error_insert)
		return c.Render(500, "error_500", nil)
	}
	db.NewRecord(grup_privilege)


	// insert permissions
	form, _ := c.MultipartForm()
	permissions := form.Value["permission[]"]
	for _, value := range permissions {
		// insert_grup_privilege_detail
		grup_privilege_detail := models.SettingGrupPrivilegeDetail{
			Id_setting_grup_privilege 	: grup_privilege.ID,
			Code_permissions 			: value,
			CreatedAt 					: current_time("2006-01-02 15:04:05"),
		}
		if error_insert := db.Create(&grup_privilege_detail); error_insert.Error != nil {
			logs.Println(error_insert)
			return c.Render(500, "error_500", nil)
		}
		db.NewRecord(grup_privilege_detail)
	}



	return c.Redirect(301, "/lib/setting/grup_privilege/")
}

func UpdateSettingGrupPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grupprivilege_3") == false{
		return c.Render(403, "error_403", nil)
	}


	requested_id := c.Param("id")

	id_setting_grup := ConvertStringToInt(c.FormValue("id_setting_grup"))
	remarks 		:= c.FormValue("remarks")
	status 			:= c.FormValue("status")


	// update_data
	var grup_privilege models.SettingGrupPrivilege
	update_grup_privilege := db.Model(&grup_privilege).Where("md5(id) = ?", requested_id).Updates(map[string]interface{}{
		"id_setting_grup"    :  id_setting_grup,
		"status"       		 :  status,
		"remarks"            :  remarks,
		"updated_at"         :  current_time("2006-01-02 15:04:05"),
	})
	if update_grup_privilege.Error != nil {
		logs.Println(update_grup_privilege.Error)
		return c.Render(500, "error_500", nil)
	}

	//  get_data_grup_privilege
	data_grup_privilege := GetDataGrupPrivilegeById(requested_id)


	// update permissions
	if DeleteSettingGrupPrivilegeDetail(requested_id) == false{
		return c.Render(500, "error_500", nil)
	}

	form, _ := c.MultipartForm()
	permissions := form.Value["permission[]"]
	for _, value := range permissions {
		// insert_grup_privilege_detail
		grup_privilege_detail := models.SettingGrupPrivilegeDetail{
			Id_setting_grup_privilege 	: ConvertStringToInt(data_grup_privilege.ID),
			Code_permissions 			: value,
			CreatedAt 					: current_time("2006-01-02 15:04:05"),
		}
		if error_insert := db.Create(&grup_privilege_detail); error_insert.Error != nil {
			logs.Println(error_insert)
			return c.Render(500, "error_500", nil)
		}
		db.NewRecord(grup_privilege_detail)
	}

	return c.Redirect(301, "/lib/setting/grup_privilege/")
}

func DeleteSettingGrupPrivilege(id_grup_privilege string /*convert_to_md5*/) bool{
	db := database.CreateCon()
	defer db.Close()

	var data models.SettingGrupPrivilege
	delete := db.Unscoped().Where("md5(id) = ?", id_grup_privilege).Delete(&data)
	if delete.Error != nil {
		logs.Println(delete.Error)
		return false
	}
	return true
}

func DeleteSettingGrupPrivilegeDetail(id_grup_privilege string /*convert_to_md5*/) bool{
	db := database.CreateCon()
	defer db.Close()

	var data models.SettingGrupPrivilegeDetail
	delete := db.Unscoped().Where("md5(id_setting_grup_privilege) = ?", id_grup_privilege).Delete(&data)
	if delete.Error != nil {
		logs.Println(delete.Error)
		return false
	}
	return true
}
