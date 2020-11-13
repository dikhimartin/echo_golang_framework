package database

import(
	"../application/models"
)

func AutoMigrate(){
	db := CreateCon()
	db.AutoMigrate(
		&models.SettingUser{},
	)
	defer db.Close()
}


