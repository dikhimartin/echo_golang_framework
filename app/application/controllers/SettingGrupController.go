package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"../models"
	"github.com/labstack/echo"
	_ "database/sql"
	"fmt"
	"../../database"
	"github.com/dikhimartin/beego-v1.12.0/utils/pagination"
	"github.com/flosch/pongo2"
)


func (c *MyCustomContext) getDataGrup() ([]models.SettingGrup) {

	db := database.CreateCon()
	defer db.Close()

	//get data grup
	rows, err := db.Query("SELECT id, name_grup FROM tb_setting_grup ORDER BY name_grup")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	each := models.SettingGrup{}
	result := []models.SettingGrup{}

	for rows.Next() {
		var id,
			name_grup []byte

		err := rows.Scan(&id, &name_grup)
		if err != nil {
			fmt.Println(err)
		}

		each.Id = string(id)
		each.Name_Grup = string(name_grup)

		result = append(result, each)
	}

	return result
}

func ListSettingGrup(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grup_2'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grup_2"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
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

	selected = "SELECT id, name_grup, status"
	// search biasa
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

	rows, err := db.Query(selected + whrs + " ORDER BY id ASC")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	defer rows.Close()

	each := models.SettingGrup{}
	result := []models.SettingGrup{}

	for rows.Next() {
		var id string
		var name_grup, status []byte

		var err = rows.Scan(&id, &name_grup, &status)

		if err != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

		var str string = id

		hasher := md5.New()
		hasher.Write([]byte(str))
		converId := hex.EncodeToString(hasher.Sum(nil))

		each.Id = converId
		each.Name_Grup = string(name_grup)
		each.Status = string(status)

		result = append(result, each)
	}

	// Lets use the Forbes top 7.
	mydata := result

	// sets paginator with the current offset (from the url query param)
	postsPerPage := 10
	paginator = pagination.NewPaginator(c.Request(), postsPerPage, len(mydata))

	// fmt.Println(paginator.Offset())
	// fetch the next posts "postsPerPage"
	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)
	//create a new page list that shows up on html
	mydatas := []models.SettingGrup{}
	for _, num := range idrange {
		//Prevent index out of range errors
		if num <= len(mydata)-1 {
			numdata := mydata[num]
			mydatas = append(mydatas, numdata)
		}
	}

	// set the paginator in context
	// also set the page list in context
	// if you also have more data, set it context
	data = pongo2.Context{
		"paginator":    paginator,
		"setting_grup": mydatas,
		"search":       search,
		"searchStatus": searchStatus}

	return c.Render(http.StatusOK, "list_setting_grup", data)
}

func AddSettingGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grup_1'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grup_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege


	return c.Render(http.StatusOK, "add_setting_grup", nil)
}

func StoreSettingGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grup_1'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grup_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	emp := new(models.SettingGrup)
	if err := c.Bind(emp); err != nil {
		return err
	}

	// Insert
	sql := "INSERT INTO tb_setting_grup(name_grup, status) VALUES(?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(emp.Name_Grup, emp.Status)
	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	return c.Redirect(301, "/lib/setting/grup/")
}

func EditSettingGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grup_3'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grup_3"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	requested_id := c.Param("id")
	var name_grup string
	var status string
	err := db.QueryRow("SELECT name_grup, status FROM tb_setting_grup WHERE md5(id) = ?", requested_id).Scan(&name_grup, &status)

	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	response := models.SettingGrup{
		Id:        requested_id,
		Name_Grup: name_grup,
		Status:    status}

	data = pongo2.Context{
		"setting_grup": response}

	return c.Render(http.StatusOK, "edit_setting_grup", data)
}

func UpdateSettingGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grup_3'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grup_3"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege
	
	emp := new(models.SettingGrup)
	if err := c.Bind(emp); err != nil {
		return err
	}

	id := c.Param("id")
	name_grup := c.FormValue("name_grup")
	status := c.FormValue("status")

	selDB, err2 := db.Prepare("UPDATE tb_setting_grup SET name_grup=?, status=? WHERE md5(id)=?")

	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	defer selDB.Close()

	selDB.Exec(name_grup, status, id)

	return c.Redirect(301, "/lib/setting/grup/")
}
