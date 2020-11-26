package api

import (
	"receipt/database"
	controllers "receipt/application/controllers"
	lib      	"receipt/lib"
	"github.com/labstack/echo"
)
var logs 		  			= lib.RecordLog("SYSTEMS -")

// ## Define Type Global
type response_json map[string]interface{}


func GetSidebarPrivilege(c echo.Context) error{
	db := database.CreateCon()
	defer db.Close()

	data_users	:= controllers.GetDataLogin(c)
	id_group    := data_users.Id_group

	// Get Permission User
	permision_user, err := db.Raw("SELECT code_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND code_permissions LIKE '%_2%' ORDER BY code_permissions DESC", id_group).Rows()
	if err != nil {
		logs.Println(err)
	}
	defer permision_user.Close()

	var data []string

	for permision_user.Next() {
		var  permision []byte
		var err = permision_user.Scan(&permision)
		if err != nil {
			logs.Println(err)
		}
		data = append(data, string(permision))
	}

	response := response_json{
		"data" 		 : data,
		"developer"  : data_users.Username,
	}
	
	return c.JSON(200, response)
}


