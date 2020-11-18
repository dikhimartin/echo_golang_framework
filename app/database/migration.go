package database

import(
	"receipt/application/models"
)

func AutoMigrate(){
	db := CreateCon()
	db.AutoMigrate(
		&models.SettingUser{},
		&models.SettingGrup{},
		&models.SettingPrivilege{},
		&models.SettingPrivilegeDetail{},
		&models.SettingGrupPrivilege{},
		&models.SettingGrupPrivilegeDetail{},
		&models.SettingUserGrup{},
		&models.Permission{},
	)
	defer db.Close()
}

func ViewMigrate(){
	db := CreateCon()
	defer db.Close()

	// v_get_user
	if DropViewIfExist("v_get_user") != false{
		if ChecktableExist("v_get_user") == true {
			db.Exec("CREATE ALGORITHM=UNDEFINED  SQL SECURITY DEFINER VIEW `v_get_user`  AS  "+ v_get_user() +";")
		}
	}

	// v_get_grup
	if DropViewIfExist("v_get_grup") != false{
		if ChecktableExist("v_get_grup") == true {
			db.Exec("CREATE ALGORITHM=UNDEFINED  SQL SECURITY DEFINER VIEW `v_get_grup`  AS  "+ v_get_grup() +";")
		}
	}

	// v_get_user_grup
	if DropViewIfExist("v_get_user_grup") != false{
		if ChecktableExist("v_get_user_grup") == true {
			db.Exec("CREATE ALGORITHM=UNDEFINED  SQL SECURITY DEFINER VIEW `v_get_user_grup`  AS  "+ v_get_user_grup() +";")
		}
	}

	// v_get_privilege
	if DropViewIfExist("v_get_privilege") != false{
		if ChecktableExist("v_get_privilege") == true {
			db.Exec("CREATE ALGORITHM=UNDEFINED  SQL SECURITY DEFINER VIEW `v_get_privilege`  AS  "+ v_get_privilege() +";")
		}
	}
	
	// v_get_grup_privilege
	if DropViewIfExist("v_get_grup_privilege") != false{
		if ChecktableExist("v_get_grup_privilege") == true {
			db.Exec("CREATE ALGORITHM=UNDEFINED  SQL SECURITY DEFINER VIEW `v_get_grup_privilege`  AS  "+ v_get_grup_privilege() +";")
		}
	}

	// v_get_grup_privilege_detail
	if DropViewIfExist("v_get_grup_privilege_detail") != false{
		if ChecktableExist("v_get_grup_privilege_detail") == true {
			db.Exec("CREATE ALGORITHM=UNDEFINED SQL SECURITY DEFINER VIEW `v_get_grup_privilege_detail`  AS  "+ v_get_grup_privilege_detail() +";")
		}
	}

	
}
