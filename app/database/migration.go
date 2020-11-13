package database

import(
	"../application/models"
)

func AutoMigrate(){
	db := CreateCon()
	db.AutoMigrate(
		&models.SettingUser{},
		&models.SettingGrup{},
		&models.SettingPrivilege{},
		&models.SettingPrivilegeDetail{},
		&models.SettingGrupPrivilege{},
		&models.SettingUserGrup{},
		&models.Permission{},
		
		&models.SampleCrud{},
	)
	defer db.Close()
}


