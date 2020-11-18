package controllers

import (
	"receipt/application/models"
	"receipt/database"
	"github.com/labstack/echo"
	"github.com/dikhimartin/beego-v1.12.0/utils/pagination"
)

// == Custom Function
func GetDataGrupById(id_grup string /*convert_to_md5*/) (models.ModelGrup){
	db := database.CreateCon()
	defer db.Close()

	var id, name_grup, status []byte
	row := db.Table("tb_setting_grup").Where("md5(id) = ?", id_grup).Select("id, name_grup, status").Row() 
	err := row.Scan(&id, &name_grup, &status)
	if err != nil{
		logs.Println(err)
	}
	data := models.ModelGrup{
		ID        	  : string(id),
		Name_Grup 	  : string(name_grup),
		Status    	  : string(status),
		Additional    : id_grup,
	}
	return data
}

func GetDataGrup() ([]models.ModelGrup) {
	db := database.CreateCon()
	defer db.Close()

	rows, err := db.Raw("SELECT id, name_grup, status FROM v_get_grup WHERE status = 'Y' ORDER BY name_grup").Rows();
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
func ListSettingGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grup_2") == false{
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

	selected = "SELECT id, name_grup, status"
	if search != "" {
		ors := " FROM tb_setting_grup WHERE name_grup LIKE '%" + search + "%' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "Y" {
		ors := " FROM tb_setting_grup WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "N" {
		ors := " FROM tb_setting_grup WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else {
		whrs = " FROM tb_setting_grup"
	}

	rows, err := db.Raw(selected + whrs + " ORDER BY id ASC").Rows()
	if err != nil {
		logs.Println(err)
		return c.Render(500, "error_500", nil)
	}
	defer rows.Close()

	each := models.ModelGrup{}
	result := []models.ModelGrup{}

	for rows.Next() {
		var id, name_grup, status []byte

		var err = rows.Scan(&id, &name_grup, &status)
		if err != nil {
			logs.Println(err)
			return c.Render(500, "error_500", nil)
		}

		each.ID 	   = ConvertToMD5(string(id))
		each.Name_Grup = string(name_grup)
		each.Status    = string(status)

		result = append(result, each)
	}

	postsPerPage := 10
	paginator 	 = pagination.NewPaginator(c.Request(), postsPerPage, len(result))

	// fetch the next posts "postsPerPage"
	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)
	mydatas := []models.ModelGrup{}
	for _, num := range idrange {
		if num <= len(result)-1 {
			numdata := result[num]
			mydatas = append(mydatas, numdata)
		}
	}

	data := response_json{
		"paginator" 	: paginator,
		"data"  		: mydatas,
		"search" 		: search,
		"searchStatus"  : searchStatus,
	}

	return c.Render(200, "list_setting_grup", data)
}

func AddSettingGrup(c echo.Context) error {
	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grup_1") == false{
		return c.Render(403, "error_403", nil)
	}
	return c.Render(200, "add_setting_grup", nil)
}

func EditSettingGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grup_3") == false{
		return c.Render(403, "error_403", nil)
	}

	requested_id  := c.Param("id")
	data 		  := GetDataGrupById(requested_id)

	response := response_json{
		"data"  : data,
	}

	return c.Render(200, "edit_setting_grup", response)
}

// == Manipulate
func StoreSettingGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grup_3") == false{
		return c.Render(403, "error_403", nil)
	}

	name_grup     := c.FormValue("name_grup")
	status  	  := c.FormValue("status")

	// insert_data
	insert := models.SettingGrup{
		Name_Grup 			: name_grup,
		Status 				: status,
		CreatedAt 			: current_time("2006-01-02 15:04:05"),
	}
	if error_insert := db.Create(&insert); error_insert.Error != nil {
		logs.Println(error_insert)
		return c.Render(500, "error_500", nil)
	}
	db.NewRecord(insert)

	return c.Redirect(301, "/lib/setting/grup/")
}

func UpdateSettingGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grup_3") == false{
		return c.Render(403, "error_403", nil)
	}

	requested_id := c.Param("id")

	name_grup     := c.FormValue("name_grup")
	status  	  := c.FormValue("status")

	// update_data
	var update models.SettingGrup
	update_data := db.Model(&update).Where("md5(id) = ?", requested_id).Updates(map[string]interface{}{
		"name_grup"    :    name_grup,
		"status"       :    status,
		"updated_at"   :    current_time("2006-01-02 15:04:05"),
	})
	if update_data.Error != nil {
		logs.Println(update_data.Error)
		return c.Render(500, "error_500", nil)
	}

	return c.Redirect(301, "/lib/setting/grup/")
}

