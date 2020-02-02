package controllers

import (
	"fmt"
	"net/http"
	"../../database"
	"../models"
	"time"
	data_user "../../api/data"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

// get data login
func (c *MyCustomContext) getDataLogin() (models.GetDataLogin) {

	db := database.CreateCon()
	defer db.Close()

	// get data login
		dt_user, err   := data_user.GetDataJWT(c)
		if err != nil{
			fmt.Println(err)
		}
		id_users 	   := dt_user["id_users"]
		id_group 	   := dt_user["id_group"]
		name_users 	   := dt_user["name_users"]
		name_group 	   := dt_user["name_group"]
		jti 		   := dt_user["jti"]
		exp 		   := dt_user["exp"]

		image 		:= ""
		extension 	:= ""

		data_users := models.GetDataLogin{
		    Id_users 			:    id_users,
		    Id_group 			:    id_group,
		    Name_users 			:    name_users,
		    Name_group 			:    name_group,
		    Jti					:    jti,
		    Exp					:    exp,
		    Image				:    image,
		    Extension			:    extension,
		}
	// end

	return data_users
}

//generator
func NewSlice(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

func Getinfologin(c echo.Context) error{

	db := database.CreateCon()
	defer db.Close()

	// get data login
		dt_user, err   := data_user.GetDataJWT(c)
		if err != nil{
			fmt.Println(err)
		}
		id_users 	   := dt_user["id_users"]
		id_group 	   := dt_user["id_group"]
		name_users 	   := dt_user["name_users"]
		name_group 	   := dt_user["name_group"]
		image 	   	   := dt_user["image"]
		extension 	   := dt_user["extension"]
		jti 		   := dt_user["jti"]
		exp 		   := dt_user["exp"]

		data_users := models.GetDataLogin{
		    Id_users 			:    id_users,
		    Id_group 			:    id_group,
		    Name_users 			:    name_users,
		    Name_group 			:    name_group,
		    Jti					:    jti,
		    Exp					:    exp,
		    Image				:    image,
		    Extension			:    extension,
		}
	// end

	data = pongo2.Context{
		"data_users"				:   data_users,
		"time":      time.Now().UnixNano(),
	}

	return c.JSON(200, data)
}

func GetSidebarPrivilege(c echo.Context) error{

	db := database.CreateCon()
	defer db.Close()

	// get data login
		dt_user, err   := data_user.GetDataJWT(c)
		if err != nil{
			fmt.Println(err)
		}
		id_group 	   := dt_user["id_group"]
	// end

	sample_crud      		:= ""
	setting_grup      		:= ""
	setting_privilege 		:= ""
	setting_user      		:= ""
	setting_grup_privilege  := ""

	// Get Permission User
	permision_user, errMen := db.Query("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions LIKE '%_2%' ORDER BY kode_permissions DESC", id_group)
	if errMen != nil {
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

