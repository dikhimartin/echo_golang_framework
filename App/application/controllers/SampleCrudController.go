package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"fmt"
	"time"
	"strings"
	"../models"
	"../../database"
	"github.com/labstack/echo"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2"
)

func ListSampleCrudController(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'samplecrud_2'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "samplecrud_2"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}

	var delete_permission []byte
	delete_permissions, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'samplecrud_4'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = delete_permissions.QueryRow(data_users.Id_group).Scan(&delete_permission)	
	defer delete_permissions.Close()
	//end check_privilege


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
	// search biasa
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

	rows, err := db.Query(selected + whrs + " ORDER BY text_input ASC")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	each 	:= models.SampleCrud{}
	result  := []models.SampleCrud{}

	for rows.Next() {
		var id string 
		var text_input,
		 	text_area, 
		 	created_by, 
		 	created_at, 
		 	updated_at, 
		 	status []byte

		var err = rows.Scan(&id, &text_input, &text_area, &created_by, &created_at, &updated_at, &status)

		if err != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

		var str string = id

		hasher := md5.New()
		hasher.Write([]byte(str))
		converId := hex.EncodeToString(hasher.Sum(nil))

		each.Id 				= converId
		each.Text_input 		= string(text_input)
		each.Text_area 			= string(text_area)
		each.Created_by 		= string(created_by)

		// format_date_created_at
		dt_created_at_conv := string(created_at)
		dt_1,_ := time.Parse("2006-01-02 15:04:05", dt_created_at_conv)
		dt_created_at := dt_1.Format("02 Jan 2006 at 15:04 PM")
		each.Created_at = dt_created_at

		// format_date_updated_at
		dt_updated_at_conv := string(updated_at)
		dt_2,_ := time.Parse("2006-01-02 15:04:05", dt_updated_at_conv)
		dt_updated_at := dt_2.Format("02 Jan 2006 at 15:04 PM")
		each.Updated_at = dt_updated_at

		each.Status 	= string(status)

		result = append(result, each)
	}

	// Lets use the Forbes top 7.
	paginate := result

	// sets paginator with the current offset (from the url query param)
	postsPerPage := 10
	paginator 	 = pagination.NewPaginator(c.Request(), postsPerPage, len(paginate))

	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)
	//create a new page list that shows up on html
	paginates := []models.SampleCrud{}
	for _, num := range idrange {
		//Prevent index out of range errors
		if num <= len(paginate)-1 {
			numdata := paginate[num]
			paginates = append(paginates, numdata)
		}
	}

	data = pongo2.Context{
		"delete_permission"		:    string(delete_permission),
		"paginator"				:    paginator,
		"posts"					:    paginates,
		"search"				:    search,
		"searchStatus"			:    searchStatus}

	return c.Render(http.StatusOK, "list_sample_crud", data)
}

func AddSampleCrudController(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'samplecrud_1'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "samplecrud_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	errorInsert := ""
	if errorFeedback != nil {
		errorInsert = "Field Required can't empty!"
		errorFeedback = nil
	}

	response := models.SampleCrud{
		Id         		  :      "",
		Text_input        :  c.FormValue("text_input")}

	data := pongo2.Context{"response": response, "status": "add", "error": errorInsert}

	return c.Render(http.StatusOK, "add_form_sample_crud", data)
}

func StoreSampleCrudController(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'samplecrud_1'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "samplecrud_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	text_input     := c.FormValue("text_input")
	text_area  	   := c.FormValue("text_area")
	status 		   := c.FormValue("status")


	// INSERT
	sql := "INSERT INTO tb_sample_crud(text_input, text_area, created_by, created_at, status) VALUES(?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(text_input, text_area,  "1", time.Now(), status)
	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	return c.Redirect(301, "/lib/sample_crud/")
}

func EditSampleCrudController(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'samplecrud_3'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "samplecrud_3"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	requested_id := c.Param("id")

	var text_input,
		text_area,
		status []byte

	err := db.QueryRow("SELECT text_input, text_area, status FROM tb_sample_crud WHERE md5(id) = ?", requested_id).Scan(&text_input, &text_area, &status)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
		fmt.Println(err)
	}

	errorUpdate := ""
	if errorFeedback != nil {
		errorUpdate = "Name tidak boleh kosong!"
		errorFeedback = nil
	}
	response := models.SampleCrud{Id: requested_id, Text_input: string(text_input), Text_area: string(text_area),  Status: string(status)}

	data = pongo2.Context{
		"response"					: 	response,
		"error"						:    errorUpdate}

	return c.Render(http.StatusOK, "edit_form_sample_crud", data)
}

func UpdateSampleCrudController(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'samplecrud_3'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "samplecrud_3"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	requested_id := c.Param("id")

	text_input 		 := c.FormValue("text_input")
	text_area      	 := c.FormValue("text_area")
	status 		     := c.FormValue("status")


	update_ex_cost, err := db.Prepare("UPDATE products SET sku=? WHERE id=?")
	defer update_ex_cost.Close()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	update_ex_cost.Exec(text_input, text_area,  1, time.Now(), status, requested_id)

	return c.Redirect(301, "/lib/sample_crud/")
}

func DeleteSampleCrudController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'samplecrud_4'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "samplecrud_4"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege


	emp := new(models.SampleCrud)
	if err := c.Bind(emp); err != nil {
		return err
	}
	requested_id := c.Param("id")

	sql := "DELETE FROM tb_sample_crud WHERE md5(id) = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(requested_id)
	if err2 != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, requested_id)
}

func DeleteAllSampleCrudController(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'samplecrud_4'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "samplecrud_4"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	
	requested_id := c.Param("id")

	// Split on comma.
	result := strings.Split(requested_id, ",")

	// Display all elements.
	for i := range result {

		sql := "DELETE FROM tb_sample_crud WHERE md5(id) = ?"
		stmt, err := db.Prepare(sql)
		if err != nil {
			fmt.Println(err)
		}
		defer stmt.Close()
		_, err2 := stmt.Exec(result[i])
		if err2 != nil {
			fmt.Println(err)
		}
	}

	// Length is 3.
	return c.JSON(http.StatusOK, len(result))
}