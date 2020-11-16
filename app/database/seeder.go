package database

import(
	"strconv"
	"encoding/json"
	"../application/models"
)

// data_seeder_here
func SeedPermission() []models.SchemePermission{
	concept := 	`[{"id":"1","name":"Create","additional":null}, {"id":"2","name":"Read\/View","additional":null}, {"id":"3","name":"Edit","additional":null}, {"id":"4","name":"Delete","additional":null}]`
	var jsonData = []byte(concept)
	var data_concept []models.SchemePermission
	err := json.Unmarshal(jsonData, &data_concept)
	if err != nil {
		logs.Println(err)
		panic(err)
	}
	return data_concept
}

func SeedGroup() []models.SchemeGroup{
	concept := `[{"id":"1","name_grup":"superadmin","status":"Y","created_at":"2020-11-16 12:13:00","updated_at":"2020-11-16 12:13:00","additional":null}]`
	var jsonData = []byte(concept)
	var data_concept []models.SchemeGroup
	err := json.Unmarshal(jsonData, &data_concept)
	if err != nil {
		logs.Println(err)
		panic(err)
	}
	return data_concept
}


func SeedGrupPrivilege() []models.SchemeGrupPrivilege{
	concept := `[{"id":"1","id_setting_grup":"1","remarks":"superadmin","status":"Y","created_at":"2020-11-16 13:29:28","updated_at":"2020-11-16 13:29:28","additional":null}]`
	var jsonData = []byte(concept)
	var data_concept []models.SchemeGrupPrivilege
	err := json.Unmarshal(jsonData, &data_concept)
	if err != nil {
		logs.Println(err)
		panic(err)
	}
	return data_concept
}

func SeedGrupPrivilegeDetail() []models.SchemeGrupPrivilegeDetail{
	concept := `[{"id":"1","id_setting_grup_privilege":"1","code_permissions":"setting.user.grupprivilege_1","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"2","id_setting_grup_privilege":"1","code_permissions":"setting.user.grupprivilege_2","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"3","id_setting_grup_privilege":"1","code_permissions":"setting.user.grupprivilege_3","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"4","id_setting_grup_privilege":"1","code_permissions":"setting.user.grupprivilege_4","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"5","id_setting_grup_privilege":"1","code_permissions":"setting.user.grup_1","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"6","id_setting_grup_privilege":"1","code_permissions":"setting.user.grup_2","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"7","id_setting_grup_privilege":"1","code_permissions":"setting.user.grup_3","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"8","id_setting_grup_privilege":"1","code_permissions":"setting.user.grup_4","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"9","id_setting_grup_privilege":"1","code_permissions":"setting.user.privilege_1","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"10","id_setting_grup_privilege":"1","code_permissions":"setting.user.privilege_2","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"11","id_setting_grup_privilege":"1","code_permissions":"setting.user.privilege_3","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"12","id_setting_grup_privilege":"1","code_permissions":"setting.user.privilege_4","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"13","id_setting_grup_privilege":"1","code_permissions":"setting.user.user_1","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"14","id_setting_grup_privilege":"1","code_permissions":"setting.user.user_2","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"15","id_setting_grup_privilege":"1","code_permissions":"setting.user.user_3","created_at":"2020-01-26 20:29:38","additional":null}, {"id":"16","id_setting_grup_privilege":"1","code_permissions":"setting.user.user_4","created_at":"2020-01-26 20:29:38","additional":null}]`
	var jsonData = []byte(concept)
	var data_concept []models.SchemeGrupPrivilegeDetail
	err := json.Unmarshal(jsonData, &data_concept)
	if err != nil {
		logs.Println(err)
		panic(err)
	}
	return data_concept
}

func SeedPrivilege() []models.SchemePrivilege{
	concept := `[{"id":"1","code_privilege":"setting.user.grup","name_menu":"Setting Grup","status":"Y","remarks":"","additional":null},{"id":"2","code_privilege":"setting.user.privilege","name_menu":"Setting Privilege","status":"Y","remarks":"","additional":null},{"id":"3","code_privilege":"setting.user.user","name_menu":"Setting User","status":"Y","remarks":"","additional":null},{"id":"4","code_privilege":"setting.user.grupprivilege","name_menu":"Setting Grup Privilege","status":"Y","remarks":"","additional":null}]`
	var jsonData = []byte(concept)
	var data_concept []models.SchemePrivilege
	err := json.Unmarshal(jsonData, &data_concept)
	if err != nil {
		logs.Println(err)
		panic(err)
	}
	return data_concept
}



// import_data
func DataSeeder(){
	db := CreateCon()
	defer db.Close()

	tb_permission 		 := SeedPermission()
	if ChecktableRecord("tb_permission") == false{
		for _, v := range tb_permission {
			id, _ := strconv.Atoi(v.ID)
			seed := models.Permission{
				ID    : id,
				Name  : v.Name,
			}
			if error_insert := db.Create(&seed); error_insert.Error != nil {
				logs.Println(error_insert)
				panic(error_insert)
			}
			db.NewRecord(seed)
		}
	}

	tb_setting_privilege 	:= SeedPrivilege()
	if ChecktableRecord("tb_setting_privilege") == false{
		for _, v := range tb_setting_privilege {
			id, _ 			   			 := strconv.Atoi(v.ID)
			seed := models.SettingPrivilege{
				ID    	   		 : id,
				Code_privilege   : v.Code_privilege,
				Name_menu        : v.Name_menu,
				Remarks          : v.Remarks,
				Status     	     : v.Status,
			}
			if error_insert := db.Create(&seed); error_insert.Error != nil {
				logs.Println(error_insert)
				panic(error_insert)
			}
			db.NewRecord(seed)
		}
	}

	if ChecktableRecord("tb_setting_privilege_detail") == false{
		for _, vprivilege := range tb_setting_privilege {
			for _, vpermission := range tb_permission{
				id_setting_privilege, _   := strconv.Atoi(vprivilege.ID)
				permission, _ 			  := strconv.Atoi(vpermission.ID)
				seed := models.SettingPrivilegeDetail{
					Id_setting_privilege     : id_setting_privilege,
					Permissions    	   		 : permission,
				}
				if error_insert := db.Create(&seed); error_insert.Error != nil {
					logs.Println(error_insert)
					panic(error_insert)
				}
				db.NewRecord(seed)
			}
		}
	}

	tb_setting_grup 		 := SeedGroup()
	if ChecktableRecord("tb_setting_grup") == false{
		for _, v := range tb_setting_grup {
			id, _ := strconv.Atoi(v.ID)
			seed := models.SettingGrup{
				ID    	   : id,
				Name_Grup  : v.Name_Grup,
				Status     : v.Status,
			}
			if error_insert := db.Create(&seed); error_insert.Error != nil {
				logs.Println(error_insert)
				panic(error_insert)
			}
			db.NewRecord(seed)
		}
	}

	tb_setting_grup_privilege 	:= SeedGrupPrivilege()
	if ChecktableRecord("tb_setting_grup_privilege") == false{
		for _, v := range tb_setting_grup_privilege {
			id, _ 			   := strconv.Atoi(v.ID)
			id_setting_grup, _ := strconv.Atoi(v.Id_setting_grup)
			seed := models.SettingGrupPrivilege{
				ID    	   		 : id,
				Id_setting_grup  : id_setting_grup,
				Remarks     	 : v.Remarks,
				Status     		 : v.Status,
			}
			if error_insert := db.Create(&seed); error_insert.Error != nil {
				logs.Println(error_insert)
				panic(error_insert)
			}
			db.NewRecord(seed)
		}
	}

	tb_setting_grup_privilege_detail 	:= SeedGrupPrivilegeDetail()
	if ChecktableRecord("tb_setting_grup_privilege_detail") == false{
		for _, v := range tb_setting_grup_privilege_detail {
			id, _ 			   			 := strconv.Atoi(v.ID)
			id_setting_grup_privilege, _ := strconv.Atoi(v.Id_setting_grup_privilege)
			seed := models.SettingGrupPrivilegeDetail{
				ID    	   		 		   : id,
				Id_setting_grup_privilege  : id_setting_grup_privilege,
				Code_permissions     	   : v.Code_permissions,
			}
			if error_insert := db.Create(&seed); error_insert.Error != nil {
				logs.Println(error_insert)
				panic(error_insert)
			}
			db.NewRecord(seed)
		}
	}

}


