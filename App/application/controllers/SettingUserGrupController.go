package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"
	"strings"
	"../models"
	_ "database/sql"
	"fmt"
	// "time"
	"../../database"
	"github.com/labstack/echo"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2"
)

func GetUserGrup() ([]models.SettingUser, []models.SettingGrup) {
	db := database.CreateCon()
	defer db.Close()

	rowsUser, errUser := db.Query(" SELECT id, full_name FROM v_get_user WHERE status = 'Y' ORDER BY full_name ASC")
	if errUser != nil {
		fmt.Println(errUser)
	}

	defer rowsUser.Close()

	eachUser := models.SettingUser{}
	resultUser := []models.SettingUser{}

	for rowsUser.Next() {
		var id, full_name []byte

		var errUser = rowsUser.Scan(&id, &full_name)

		if errUser != nil {
			fmt.Println(errUser)
		}

		eachUser.Id = string(id)

		resultUser = append(resultUser, eachUser)
	}

	rowsGrup, errGrup := db.Query(" SELECT id, name_grup FROM tb_setting_grup WHERE status = 'Y' ORDER BY name_grup ASC")
	if errGrup != nil {
		fmt.Println(errGrup)
	}

	defer rowsGrup.Close()

	eachGrup := models.SettingGrup{}
	resultGrup := []models.SettingGrup{}

	for rowsGrup.Next() {
		var id, name_grup []byte

		var errGrup = rowsGrup.Scan(&id, &name_grup)

		if errGrup != nil {
			fmt.Println(errGrup)
		}

		eachGrup.Id = string(id)
		eachGrup.Name_Grup = string(name_grup)

		resultGrup = append(resultGrup, eachGrup)
	}

	return resultUser, resultGrup

}

func ListSettingUserGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	var selected string
	var whrs string
	var search_user string
	var search_grup string
	var search_status string

	if reqSearchUser := c.FormValue("id_user"); reqSearchUser != "" {
		search_user = reqSearchUser
	}

	if reqSearchGrup := c.FormValue("id_grup"); reqSearchGrup != "" {
		search_grup = reqSearchGrup
	}

	if reqSearchStatus := c.FormValue("status"); reqSearchStatus != "" {
		search_grup = reqSearchStatus
	}

	selected = "SELECT id, full_name, name_grup, status"
	// search biasa
	if search_user != "" || search_grup != "" {
		ors := " FROM v_get_user_grup WHERE id_user = '" + search_user + "' OR id_setting_grup = '" + search_grup + "' "
		whrs += ors
	} else if search_status != "" && search_status == "Y" {
		ors := " FROM v_get_user_grup WHERE status = '" + search_status + "' "
		whrs += ors
	} else if search_status != "" && search_status == "N" {
		ors := " FROM v_get_user_grup WHERE status = '" + search_status + "' "
		whrs += ors
	} else {
		whrs = " FROM v_get_user_grup"
	}

	rows, err := db.Query(selected + whrs + " ORDER BY full_name, name_grup ASC")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	defer rows.Close()

	each := models.SettingUserGrup{}
	result := []models.SettingUserGrup{}

	for rows.Next() {
		var id string
		var full_name, name_grup,status []byte

		var err = rows.Scan(&id, &full_name, &name_grup, &status)

		if err != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

		var str string = id

		hasher := md5.New()
		hasher.Write([]byte(str))
		converId := hex.EncodeToString(hasher.Sum(nil))

		each.Id = converId
		each.Id_setting_user = string(full_name)
		each.Id_setting_grup = string(name_grup)
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
	mydatas := []models.SettingUserGrup{}
	for _, num := range idrange {
		//Prevent index out of range errors
		if num <= len(mydata)-1 {
			numdata := mydata[num]
			mydatas = append(mydatas, numdata)
		}
	}

	resultUser, resultGrup := GetUserGrup()

	// set the paginator in context
	// also set the page list in context
	// if you also have more data, set it context
	data = pongo2.Context{
		"paginator":         paginator,
		"setting_user_grup": mydatas,
		"search_user":       search_user,
		"search_grup":       search_grup,
		"search_status":     search_status,
		"resultUser":        resultUser,
		"resultGrup":        resultGrup}

	return c.Render(http.StatusOK, "list_setting_user_grup", data)
}

func AddSettingUserGrup(c echo.Context) error {
	resultUser, resultGrup := GetUserGrup()

	errorInsert := ""
	if errorFeedback != nil {
		if strings.Contains(errorFeedback.Error(), "user dan grup") {
			errorInsert = "Nama user pada grup tersebut sudah ada!"
			errorFeedback = nil
		}
	}

	response := models.SettingUserGrup{
		Id_setting_user: c.FormValue("id_user"), 
		Id_setting_grup: c.FormValue("id_grup"), 
		Status: c.FormValue("status")}

	data = pongo2.Context{
		"error": errorInsert,
		"response": response,
		"resultUser": resultUser,
		"resultGrup": resultGrup}

	return c.Render(http.StatusOK, "add_setting_user_grup", data)
}

func StoreSettingUserGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	emp := new(models.SettingUserGrup)
	if err := c.Bind(emp); err != nil {
		return err
	}

	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")

	// Insert
	sql := "INSERT INTO tb_setting_user_grup(id_setting_user, id_setting_grup, add_date, status) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(emp.Id_setting_user, emp.Id_setting_grup, today, emp.Status)

	if err2 != nil {
		if strings.Contains(err2.Error(), "user dan grup") {
			errorFeedback = err2
			return AddSettingUserGrup(c)
		} else {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}
	}

	return c.Redirect(301, "/setting/user_grup/")
}

func EditSettingUserGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	resultUser, resultGrup := GetUserGrup()

	requested_id := c.Param("id")
	var full_name, name_grup, status string

	err := db.QueryRow("SELECT full_name, name_grup, status FROM v_get_user_grup WHERE md5(id) = ?", requested_id).Scan(&full_name, &name_grup, &status)

	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	errorUpdate := ""
	if errorFeedback != nil {
		if strings.Contains(errorFeedback.Error(), "user dan grup") {
			errorUpdate = "Nama user pada grup tersebut sudah ada!"
			errorFeedback = nil
		}
	}

	response := models.SettingUserGrup{
		Id:      requested_id,
		Id_setting_user: full_name,
		Id_setting_grup: name_grup,
		Status:  status}

	data = pongo2.Context{
		"error": errorUpdate,
		"resultUser":        resultUser,
		"resultGrup":        resultGrup,
		"setting_user_grup": response}

	return c.Render(http.StatusOK, "edit_setting_user_grup", data)
}

func UpdateSettingUserGrup(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()
	
	id := c.Param("id")
	id_setting_user := c.FormValue("id_setting_user")
	id_setting_grup := c.FormValue("id_setting_grup")
	status := c.FormValue("status")

	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")

	selDB, err2 := db.Prepare("UPDATE tb_setting_user_grup SET id_setting_user=?, id_setting_grup=?, status=?, update_date=? WHERE md5(id)=?")

	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	defer selDB.Close()

	_, err3 := selDB.Exec(id_setting_user, id_setting_grup, status, today, id)

	if err3 != nil {
		errorFeedback = err3
		return EditSettingUserGrup(c)
	}

	return c.Redirect(301, "/setting/user_grup/")
}
