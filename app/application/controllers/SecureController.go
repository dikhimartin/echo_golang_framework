package controllers

import (
	"../../database"
)


func CheckPrivileges(id_grup int, code_permissions string) bool{
	db := database.CreateCon()
	defer db.Close()

	// check_privilege
	var check_privilege []byte
	row := db.Table("v_get_grup_privilege_detail").Where("id_setting_grup = ? AND code_permissions = ?", id_grup, code_permissions).Select("code_permissions").Row() // (*sql.Row)
	err := row.Scan(&check_privilege)
	if err != nil{
		logs.Println(err)
	}
	if string(check_privilege) == ""{
		return false
	}
	return true
}