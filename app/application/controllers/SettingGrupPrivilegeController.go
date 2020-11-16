package controllers

import (
	"../models"
	"../../database"
	"github.com/labstack/echo"
	"github.com/dikhimartin/beego-v1.12.0/utils/pagination"
)

func ListSettingGrupPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	data_users	:= GetDataLogin(c)
	if CheckPrivileges(data_users.Id_group, "setting.user.grupprivilege_2") == false{
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

	selected = "SELECT id, id_setting_grup, name_grup, status, created_at, updated_at"
	if search != "" {
		ors := " FROM v_get_grup_privilege WHERE name_grup LIKE '%" + search + "%' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "Y" {
		ors := " FROM v_get_grup_privilege WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else if searchStatus != "" && searchStatus == "N" {
		ors := " FROM v_get_grup_privilege WHERE status = '" + searchStatus + "' "
		whrs += ors
	} else {
		whrs = " FROM v_get_grup_privilege"
	}

	rows, err := db.Raw(selected + whrs + " ORDER BY name_grup ASC").Rows()
	if err != nil {
		logs.Println(err)
		return c.Render(500, "error_500", nil)
	}

	defer rows.Close()

	each   := models.ModelGrupPrivilege{}
	result := []models.ModelGrupPrivilege{}

	for rows.Next() {
		var	id, id_setting_grup, name_grup, status, created_at, updated_at[]byte

		var err = rows.Scan(&id, &id_setting_grup, &name_grup, &status, &created_at, &updated_at)
		if err != nil {
			logs.Println(err)
			return c.Render(500, "error_500", nil)
		}

		each.ID  			 = ConvertToMD5(string(id))
		each.Id_setting_grup = string(id_setting_grup)
		each.Name_grup  	 = string(name_grup)
		each.Status  		 = string(status)
		each.CreatedAt 	 	 = FormatDate(string(created_at), "02 January 2006 at 15:04 PM")
		each.UpdatedAt 	 	 = FormatDate(string(updated_at), "02 January 2006 at 15:04 PM")

		result = append(result, each)
	}

	postsPerPage := 10
	paginator 	 = pagination.NewPaginator(c.Request(), postsPerPage, len(result))

	// fetch the next posts "postsPerPage"
	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)
	mydatas := []models.ModelGrupPrivilege{}
	for _, num := range idrange {
		if num <= len(result)-1 {
			numdata := result[num]
			mydatas = append(mydatas, numdata)
		}
	}

	data := response_json{
		"paginator"		: paginator,
		"data"			: mydatas,
		"search"		: search,
		"searchStatus"	: searchStatus,
	}

	return c.Render(200, "list_setting_grup_privilege", data)
}


