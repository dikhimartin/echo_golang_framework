package controllers

import (
	"strings"
	"html/template"
	"../models"
	"../../database"
	"github.com/labstack/echo"
)

func GetMasterPermissions() ([]models.Permission) {
	db := database.CreateCon()
	defer db.Close()

	rows, err := db.Raw("SELECT id, name FROM tb_permission").Rows()
	if err != nil {
		logs.Println(err)
	}
	defer rows.Close()

	each   := models.Permission{}
	result := []models.Permission{}

	for rows.Next() {
		var id, name []byte
		err = rows.Scan(&id, &name)
		if err != nil {
			logs.Println(err)
		}
		each.ID 	= ConvertStringToInt(string(id))
		each.Name   = string(name)
		result 	    = append(result, each)
	}

	return result
}


// == View
func ListSettingPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()
	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.privilege_2") == false{
		return c.Render(403, "error_403", nil)
	}

	var selected     string
	var whrs         string
	var search       string
	var searchStatus string

	if reqSearch := c.FormValue("search"); reqSearch != "" {
		search = reqSearch
	}
	if reqSearchStatus := c.FormValue("searchStatus"); reqSearchStatus != "" {
		searchStatus = reqSearchStatus
	}

	selected = "SELECT id, code_privilege, name_menu, status, remarks"
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

	rows, err := db.Raw(selected + whrs + " ORDER BY code_privilege").Rows()
	if err != nil {
		logs.Println(err)
		return c.Render(500, "error_500", nil)
	}
	defer rows.Close()

	each := models.ModelPrivilege{}
	html := ""

	var new_parent, new_menu, new_submenu string

	for rows.Next() {
		var id, code_privilege, name_menu, status, remarks []byte
		var err = rows.Scan(&id, &code_privilege, &name_menu, &status, &remarks)
		if err != nil {
			logs.Println(err)
			return c.Render(500, "error_500", nil)
		}

		each.ID         = ConvertToMD5(string(id))
		each.Name_menu  = string(name_menu)
		each.Status     = string(status)
		each.Remarks    = string(remarks)

		// SPLIT
		s := strings.Split(string(code_privilege), ".")

		parent := s[0] // parent
		menu := ""
		if len(s) < 2{
			menu = string(code_privilege)   // menu
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
		}

		action_edit := "<a href='/lib/setting/privilege/editform/" + ConvertToMD5(string(id)) + "/' class='btn btn-sm btn-info' data-toggle='tooltip' data-placement='top' title='Edit data!'><i class='fa fa-pencil'></i></a>"

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
				html += "<td>" + string(remarks) + "</td>"
				html += "<td>" + action_edit + "</td>"
			} else {
				html += "<td></td>"
				html += "<td>" + label_status + "</td>"
				html += "<td>" + string(remarks) + "</td>"
				html += "<td>" + action_edit + "</td>"
			}
			html += "</tr>"
		} else {
			html += "<td>" + label_status + "</td>"
			html += "<td>" + string(remarks) + "</td>"
			html += "<td>" + action_edit + "</td>"
			html += "</tr>"
		}

		new_parent = parent
		new_menu = menu
		new_submenu = submenu
	}

	data := response_json{
		"search"       : search,
		"searchStatus" : searchStatus,
		"getData" 	   : template.HTML(html),
	}

	return c.Render(200, "list_setting_privilege", data)
}

func AddSettingPrivilege(c echo.Context) error {
	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.privilege_1") == false{
		return c.Render(403, "error_403", nil)
	}

	permission			:= GetMasterPermissions()

	data := response_json{
		"permission" :  permission,
	}
	return c.Render(200, "add_setting_privilege", data)
}

func EditSettingPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.privilege_3") == false{
		return c.Render(403, "error_403", nil)
	}
	requested_id := c.Param("id")
	data 		 := GetDataGrupByIdMd5(requested_id)
	
	response := response_json{
		"data"  : data,
	}
	return c.Render(200, "edit_setting_privilege", response)
}

