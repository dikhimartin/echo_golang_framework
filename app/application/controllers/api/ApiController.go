package api

import (
	"fmt"
	"net/http"
	"receipt/application/models"
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

	sample_crud      		:= ""
	setting_grup      		:= ""
	setting_privilege 		:= ""
	setting_user      		:= ""
	setting_grup_privilege  := ""

	// Get Permission User
	permision_user, err := db.Raw("SELECT code_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND code_permissions LIKE '%_2%' ORDER BY code_permissions DESC", id_group).Rows()
	if err != nil {
		logs.Println(err)
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	defer permision_user.Close()

	for permision_user.Next() {
		var  permision []byte
		var err = permision_user.Scan(&permision)
		if err != nil {
			fmt.Println(err)
		}
		// SAMPLE CRUD
		if string(permision) == "samplecrud_2"{
			sample_crud = string(permision)
		}
		
		// SETTING
		if string(permision) == "setting.user.grup_2"{
			setting_grup = string(permision)
		}
		if string(permision) == "setting.user.privilege_2"{
			setting_privilege = string(permision)
		}
		if string(permision) == "setting.user.user_2"{
			setting_user = string(permision)
		}
		if string(permision) 	   == "setting.user.grupprivilege_2"{
			setting_grup_privilege = string(permision)
		}
	}
	data_menus := models.SidebarMenu{
		Setting_grup 			:    string(setting_grup), 
		Setting_privilege 		:    string(setting_privilege), 
		Setting_user 			:    string(setting_user), 
		Setting_grup_privilege 	:    string(setting_grup_privilege),	
		Sample_crud 			: 	 string(sample_crud), 
	}
	return c.JSON(200, data_menus)
}


