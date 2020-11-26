package controllers

import (
	"strings"
	"receipt/application/models"
	"receipt/database"
	"github.com/dikhimartin/beego-v1.12.0/utils/pagination"
	"github.com/labstack/echo"
)


// == Custom Function
func GetDataSampleCrudById(id_sample_crud string /*convert_to_md5*/) (models.ModelSampleCrud){
	db := database.CreateCon()
	defer db.Close()

	var id, text_input, text_area, status []byte
	row := db.Table("tb_sample_crud").Where("md5(id) = ?", id_sample_crud).Select("id, text_input, text_area, status").Row() 
	err := row.Scan(&id, &text_input, &text_area, &status)
	if err != nil{
		logs.Println(err)
	}
	data := models.ModelSampleCrud{
		ID        	  : string(id),
		Text_input 	  : string(text_input),
		Text_area 	  : string(text_area),
		Status    	  : string(status),
		Additional    : id_sample_crud,
	}
	return data
}

func GetDataSampleCrud() ([]models.ModelSampleCrud) {
	db := database.CreateCon()
	defer db.Close()

	rows, err := db.Raw("SELECT id, text_input, text_area, status FROM v_get_grup WHERE status = 'Y' ORDER BY text_input").Rows();
	if err != nil {
		logs.Println(err)
	}
	defer rows.Close()
	each   := models.ModelSampleCrud{}
	result := []models.ModelSampleCrud{}

	for rows.Next() {
		var id, text_input, text_area, status []byte
		err = rows.Scan(&id, &text_input, &text_area, &status)
		if err != nil {
			logs.Println(err)
		}

		each.ID 			 = string(id)
		each.Text_input 	 = string(text_input)
		each.Text_area 		 = string(text_area)
		each.Status  		 = string(status)

		result = append(result, each)
	}

	return result
}


// == View
func ListSampleCrudController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "samplecrud_2") == false{
		return c.Render(403, "error_403", nil)
	}

	delete_permission := CheckPrivileges(data_users.Id_group, "samplecrud_4")
	

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

	selected = "SELECT id, text_input, text_area, created_by, created_at, updated_at, status"
	if search != "" {
		ors := " FROM tb_sample_crud WHERE concat(text_input, text_area) LIKE '%" + search + "%' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "Y" {
		ors := " FROM tb_sample_crud WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "N" {
		ors := " FROM tb_sample_crud WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else {
		whrs = " FROM tb_sample_crud"
	}

	rows, err := db.Raw(selected + whrs + " ORDER BY text_input ASC").Rows()
	if err != nil {
		logs.Println(err)
		return c.Render(500, "error_500", nil)
	}

	each 	:= models.ModelSampleCrud{}
	result  := []models.ModelSampleCrud{}

	for rows.Next() {
		var id, text_input, text_area,  created_by,  created_at,  updated_at,  status []byte
		var err = rows.Scan(&id, &text_input, &text_area, &created_by, &created_at, &updated_at, &status)
		if err != nil {
			logs.Println(err)
			return c.Render(500, "error_500", nil)
		}

		each.ID 	   			= ConvertToMD5(string(id))
		each.Text_input 		= string(text_input)
		each.Text_area 			= string(text_area)
		each.Status 			= string(status)
		each.CreatedAt 	 	 	= FormatDate(string(created_at), "02 January 2006 at 15:04 PM")
		each.UpdatedAt 	 	 	= FormatDate(string(updated_at), "02 January 2006 at 15:04 PM")

		result = append(result, each)
	}


	postsPerPage := 10
	paginator 	 = pagination.NewPaginator(c.Request(), postsPerPage, len(result))

	// fetch the next posts "postsPerPage"
	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)
	mydatas := []models.ModelSampleCrud{}
	for _, num := range idrange {
		if num <= len(result)-1 {
			numdata := result[num]
			mydatas = append(mydatas, numdata)
		}
	}


	data := response_json{
		"delete_permission" 	: delete_permission,
		"paginator" 			: paginator,
		"data"  				: mydatas,
		"search" 				: search,
		"searchStatus"  		: searchStatus,
	}

	return c.Render(200, "list_sample_crud", data)
}

func AddSampleCrudController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "samplecrud_1") == false{
		return c.Render(403, "error_403", nil)
	}

	return c.Render(200, "add_form_sample_crud", "data")
}

func EditSampleCrudController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "samplecrud_3") == false{
		return c.Render(403, "error_403", nil)
	}

	requested_id  := c.Param("id")
	data 		  := GetDataSampleCrudById(requested_id)

	response := response_json{
		"data"  : data,
	}

	return c.Render(200, "edit_form_sample_crud", response)
}

// == Manipulate
func StoreSampleCrudController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "samplecrud_1") == false{
		return c.Render(403, "error_403", nil)
	}

	text_input     := c.FormValue("text_input")
	text_area  	   := c.FormValue("text_area")
	status 		   := c.FormValue("status")

	// insert_data
	insert := models.SampleCrud{
		Text_input 			: text_input,
		Text_area 			: text_area,
		Status 				: status,
		Created_by 			: data_users.Id_user,
		CreatedAt 			: current_time("2006-01-02 15:04:05"),
	}
	if error_insert := db.Create(&insert); error_insert.Error != nil {
		logs.Println(error_insert)
		return c.Render(500, "error_500", nil)
	}
	db.NewRecord(insert)

	return c.Redirect(301, "/lib/sample_crud/")
}

func UpdateSampleCrudController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "samplecrud_3") == false{
		return c.Render(403, "error_403", nil)
	}

	requested_id := c.Param("id")

	text_input 		 := c.FormValue("text_input")
	text_area      	 := c.FormValue("text_area")
	status 		     := c.FormValue("status")

	// update_data
	var update models.SampleCrud
	update_data := db.Model(&update).Where("md5(id) = ?", requested_id).Updates(map[string]interface{}{
		"text_input"   :    text_input,
		"text_area"    :    text_area,
		"updated_by"   :    data_users.Id_user,
		"status"       :    status,
		"updated_at"   :    current_time("2006-01-02 15:04:05"),
	})
	if update_data.Error != nil {
		logs.Println(update_data.Error)
		return c.Render(500, "error_500", nil)
	}


	return c.Redirect(301, "/lib/sample_crud/")
}

func DeleteSampleCrudController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "samplecrud_4") == false{
		return c.JSON(200, false)
	}

	requested_id := c.Param("id")
	var model models.SampleCrud
	delete := db.Unscoped().Where("md5(id) = ?", requested_id).Delete(&model)
	if delete.Error != nil {
		logs.Println(delete.Error)
		return c.JSON(200, false)
	}	

	return c.JSON(200, true)
}

func DeleteAllSampleCrudController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "samplecrud_4") == false{
		return c.JSON(200, false)
	}
	
	requested_id := c.Param("id")

	result := strings.Split(requested_id, ",")

	for i := range result {
		var model models.SampleCrud
		delete := db.Unscoped().Where("md5(id) = ?", result[i]).Delete(&model)
		if delete.Error != nil {
			logs.Println(delete.Error)
			return c.JSON(200, false)
		}	
	}

	return c.JSON(200, true)
}