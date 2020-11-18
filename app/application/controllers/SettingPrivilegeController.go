package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"net/http"
	"strings"
	_ "database/sql"
	"fmt"
	"strconv"
	// "../helpers"
	// "time"
	"../models"
	"../../database"
	"github.com/labstack/echo"
	"github.com/flosch/pongo2"
)

func (c *MyCustomContext) GetMasterPermissions() ([]models.MasterPermissions) {

	db := database.CreateCon()
	defer db.Close()

	//get
	rows, err := db.Query("SELECT id, name FROM tb_master_permission")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	each := models.MasterPermissions{}
	result := []models.MasterPermissions{}

	for rows.Next() {

		var id,
			value []byte

		err2 := rows.Scan(&id, &value)
		if err2 != nil {
			fmt.Println(err2)
		}

		each.Key 	= string(id)
		each.Value  = string(value)

		result = append(result, each)
	}

	return result
}

func ListSettingPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.privilege_2'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.privilege_2"{
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

	selected = "SELECT id, kode_privilege, name_menu, status, keterangan"
	// search biasa
	if search != "" {
		ors := " FROM tb_setting_privilege WHERE name_menu LIKE '%" + search + "%' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "Y" {
		ors := " FROM tb_setting_privilege WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "N" {
		ors := " FROM tb_setting_privilege WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else {
		whrs = " FROM tb_setting_privilege"
	}

	rows, err := db.Query(selected + whrs + " ORDER BY kode_privilege")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	defer rows.Close()

	each := models.SettingPrivilege{}
	// result := []models.SettingPrivilege{}

	html := ""

	var new_parent, new_menu, new_submenu string

	for rows.Next() {
		var id string
		var kode_privilege, name_menu, status, keterangan []byte

		var err = rows.Scan(&id, &kode_privilege, &name_menu, &status, &keterangan)

		if err != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

		var str string = id

		hasher := md5.New()
		hasher.Write([]byte(str))
		converId := hex.EncodeToString(hasher.Sum(nil))

		each.Id = converId
		each.Name_Menu = string(name_menu)
		each.Status = string(status)
		each.Keterangan = string(keterangan)

		// result = append(result, each)

		// SPLIT
		s := strings.Split(string(kode_privilege), ".")

		parent := s[0] // parent
		menu := ""
		if len(s) < 2{
			menu = string(kode_privilege)   // menu
		}else{
			menu = s[1]   // menu			
		}
		submenu := ""		
		if len(s) > 2 {
			submenu = s[2] // submenu
		}

		capitalize := "class='text-capitalize'"

		if parent != new_parent {
			html += "<tr>" +
				"<td " + capitalize + ">" + parent + "</td>"
		} else {
			html += "<tr>" +
				"<td></td>"
		}

		label_status := ""
		if string(status) == "Y" {
			label_status = "<label class='label label-success'>Aktif</label>"
		} else if string(status) == "N" {
			label_status = "<label class='label label-danger'>Non-Aktif</label>"
		} else {
			label_status = "<label class='label label-danger'>Kosong</label>"
		}

		action_edit := "<a href='/lib/setting/privilege/editform/" + converId + "/' class='btn btn-sm btn-info' data-toggle='tooltip' data-placement='top' title='Edit data!'><i class='fa fa-pencil'></i></a>"

		colspan := "1"
		if len(s) < 3 {
			colspan = "2"
		}

		// menu
		if menu != new_menu {
			html += "<td " + capitalize + " colspan=" + colspan + ">" + menu + "</td>"
		} else {
			html += "<td></td>"
		}

		if len(s) > 2 {
			// submenu
			if submenu != new_submenu {
				html += "<td " + capitalize + ">" + submenu + "</td>"
				html += "<td>" + label_status + "</td>"
				html += "<td>" + string(keterangan) + "</td>"
				html += "<td>" + action_edit + "</td>"
			} else {
				html += "<td></td>"
				html += "<td>" + label_status + "</td>"
				html += "<td>" + string(keterangan) + "</td>"
				html += "<td>" + action_edit + "</td>"
			}
			html += "</tr>"
		} else {
			html += "<td>" + label_status + "</td>"
			html += "<td>" + string(keterangan) + "</td>"
			html += "<td>" + action_edit + "</td>"
			html += "</tr>"
		}

		new_parent = parent
		new_menu = menu
		new_submenu = submenu
	}

	data = pongo2.Context{
		"search":       search,
		"searchStatus": searchStatus,
		"getData":      template.HTML(html)}

	return c.Render(http.StatusOK, "list_setting_privilege", data)
}

func AddSettingPrivilege(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.privilege_1'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.privilege_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege


	GetHakAkses			:= cc.GetMasterPermissions()

	data = pongo2.Context{
		"hak_akses":              GetHakAkses}

	return c.Render(http.StatusOK, "add_setting_privilege", data)
}

func StoreSettingPrivilege(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.privilege_1'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.privilege_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	emp := new(models.SettingPrivilege)
	if err := c.Bind(emp); err != nil {
		return err
	}

	// Insert tb_setting_privilege
	sql := "INSERT INTO tb_setting_privilege(kode_privilege, name_menu, status, keterangan) VALUES(?, ?, ?, ?)"
	insert, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	res, err2 := insert.Exec(emp.Kode_Privilege, emp.Name_Menu, emp.Status, emp.Keterangan)
	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
		// fmt.Println(err2)
	}

	sum_last_id, _ := res.LastInsertId()
	get_last_id := strconv.FormatInt(int64(sum_last_id), 10)

	// insert tb_setting_privilege_details
	form, err := c.MultipartForm()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	permissions 			:= form.Value["permissions[]"]
	for _, value := range permissions {
		sql := "INSERT INTO tb_setting_privilege_detail(id_setting_privilege, permissions) VALUES(?, ?)"
		insert_detail, err := db.Prepare(sql)
		if err != nil {
			// fmt.Println(err)
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}
		defer insert_detail.Close()
		_, err2 := insert_detail.Exec(get_last_id, value)
		if err2 != nil {
			// fmt.Println(err2)
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

	}


	return c.Redirect(301, "/lib/setting/privilege/")
}

func EditSettingPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.privilege_3'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.privilege_3"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege


	requested_id := c.Param("id")
	var kode_privilege, name_menu, status, keterangan string

	err := db.QueryRow("SELECT kode_privilege, name_menu, status, keterangan FROM tb_setting_privilege WHERE md5(id) = ?", requested_id).Scan(&kode_privilege, &name_menu, &status, &keterangan)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	response := models.SettingPrivilege{
		Id:             requested_id,
		Kode_Privilege: kode_privilege,
		Name_Menu:      name_menu,
		Status:         status,
		Keterangan:     keterangan}


	//get hak aksess
	rows, err2 := db.Query("SELECT id, name FROM tb_master_permission")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer rows.Close()

	eachHakAkses := models.MasterPermissions{}
	resultHakAkses := []models.MasterPermissions{}

	for rows.Next() {

		var id_master_permissions,
			value []byte

		err2 := rows.Scan(&id_master_permissions, &value)
		if err2 != nil {
			fmt.Println(err2)
		}

		// check or unchecked
		var checked_or_unchecked []byte
		err := db.QueryRow("SELECT permissions FROM tb_setting_privilege_detail WHERE md5(id_setting_privilege) = ? AND permissions = ?", requested_id, string(id_master_permissions)).Scan(&checked_or_unchecked)
		if err != nil {
			fmt.Println(err)
		}
		permissions := ""
		if string(checked_or_unchecked) == ""{
			permissions = "unchecked"
		}else if string(checked_or_unchecked) != ""{
			permissions = "checked"
		}

		eachHakAkses.Key 		 = string(id_master_permissions)
		eachHakAkses.Value  	 = string(value)
		eachHakAkses.Additional  = permissions

		resultHakAkses = append(resultHakAkses, eachHakAkses)
	}

	data := pongo2.Context{
		"setting_privilege" :  response,
		"resultHakAkses"	:  resultHakAkses,
	}

	return c.Render(http.StatusOK, "edit_setting_privilege", data)
}

func UpdateSettingPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.privilege_3'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.privilege_3"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege
	
	emp := new(models.SettingPrivilege)
	if err := c.Bind(emp); err != nil {
		return err
	}

	id := c.Param("id")
	kode_privilege := c.FormValue("kode_privilege")
	name_menu := c.FormValue("name_menu")
	status := c.FormValue("status")
	keterangan := c.FormValue("keterangan")

	selDB, err2 := db.Prepare("UPDATE tb_setting_privilege SET kode_privilege=?, name_menu=?, status=?, keterangan=? WHERE md5(id)=?")
	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	defer selDB.Close()
	selDB.Exec(kode_privilege, name_menu, status, keterangan, id)


	// get id_setting_privilege 
	var id_setting_privilege []byte
	err := db.QueryRow("SELECT id FROM tb_setting_privilege WHERE md5(id) = ?", id).Scan(&id_setting_privilege)
	if err != nil {
		fmt.Println(err)
	}

	// Delete
	sql2 := "DELETE FROM tb_setting_privilege_detail WHERE md5(id_setting_privilege) = ?"
	stmt, err := db.Prepare(sql2)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err3 := stmt.Exec(id)
	if err3 != nil {
		fmt.Println(err)
	}

	// update tb_setting_privilege_details
	form, err := c.MultipartForm()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	permissions := form.Value["permissions[]"]
	for _, value := range permissions {
		sql := "INSERT INTO tb_setting_privilege_detail(id_setting_privilege, permissions) VALUES(?, ?)"
		insert_detail, err := db.Prepare(sql)
		if err != nil {
			fmt.Print(err.Error())
		}
		defer insert_detail.Close()
		_, err4 := insert_detail.Exec(string(id_setting_privilege), value)
		if err4 != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

	}


	return c.Redirect(301, "/lib/setting/privilege/")
}
