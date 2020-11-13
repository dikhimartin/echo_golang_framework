package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// "html/template"
	// "strings"
	// "sort"
	// "../helpers"
	"strconv"
	"time"
	"net/http"
	_ "database/sql"
	"fmt"
	"../../database"
	"../models"
	"github.com/dikhimartin/beego-v1.12.0/utils/pagination"
	"github.com/labstack/echo"
	"github.com/flosch/pongo2"
)

func GetGrupPrivilege() ([]models.SettingUserGrup, []models.SettingPrivilege) {
	db := database.CreateCon()
	defer db.Close()

	rowsUserGrup, errUserGrup := db.Query(" SELECT id, name_grup, status FROM v_get_setting_grup WHERE status = 'Y' AND id_setting_grup_privilege IS NULL ORDER BY name_grup")
	if errUserGrup != nil {
		fmt.Println(errUserGrup)
	}

	defer rowsUserGrup.Close()

	eachUserGrup := models.SettingUserGrup{}
	resultUserGrup := []models.SettingUserGrup{}

	for rowsUserGrup.Next() {
		var id, name_grup, status []byte

		var err = rowsUserGrup.Scan(&id, &name_grup, &status)

		if err != nil {
			fmt.Println(errUserGrup)
		}

		eachUserGrup.Id = string(id)
		eachUserGrup.Id_setting_grup = string(name_grup)
		eachUserGrup.Status = string(status)

		resultUserGrup = append(resultUserGrup, eachUserGrup)
	}

	rowsPrivilege, errPrivilege := db.Query(" SELECT kode_privilege, name_menu, status FROM tb_setting_privilege ORDER BY kode_privilege")
	if errPrivilege != nil {
		fmt.Println(errPrivilege)
	}

	defer rowsPrivilege.Close()

	eachPrivilege := models.SettingPrivilege{}
	resultPrivilege := []models.SettingPrivilege{}

	for rowsPrivilege.Next() {
		var kode_privilege, name_menu, status []byte

		var err = rowsPrivilege.Scan(&kode_privilege, &name_menu, &status)

		if err != nil {
			fmt.Println(errPrivilege)
		}

		eachPrivilege.Kode_Privilege = string(kode_privilege)
		eachPrivilege.Name_Menu = string(name_menu)
		eachPrivilege.Status = string(status)

		resultPrivilege = append(resultPrivilege, eachPrivilege)
	}

	return resultUserGrup, resultPrivilege
}

func ListSettingGrupPrivilege(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grupprivilege_2'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grupprivilege_2"{
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

	selected = "SELECT id, id_setting_grup, name_grup, status, created_at, updated_at"
	// search biasa
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

	rows, err := db.Query(selected + whrs + " ORDER BY name_grup ASC")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	defer rows.Close()

	each := models.SettingGrupPrivilege{}
	result := []models.SettingGrupPrivilege{}

	for rows.Next() {
		var id string
		var	id_setting_grup, 
		    name_grup,
			status,
			created_at,
			updated_at[]byte

		var err = rows.Scan(&id, &id_setting_grup, &name_grup, &status, &created_at, &updated_at)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}

		var str string = id
		hasher := md5.New()
		hasher.Write([]byte(str))
		converId := hex.EncodeToString(hasher.Sum(nil))

		dtstr1 := string(created_at)
		dt,_ := time.Parse("2006-01-02 15:04:05", dtstr1)
		date_create := dt.Format("02 Jan 2006 at 15:04 PM")

		dtstr2 := string(updated_at)
		dt2,_ := time.Parse("2006-01-02 15:04:05", dtstr2)
		date_update := dt2.Format("02 Jan 2006 at 15:04 PM")

		if date_create == "01 Jan 0001 at 00:00 AM"{
			date_create = "-"
		}
		if date_update == "01 Jan 0001 at 00:00 AM"{
			date_update = "-"
		}

		each.Id = converId
		each.Name_grup = string(name_grup)
		each.Id_setting_grup = string(id_setting_grup)
		each.Status = string(status)
		each.Created_at = date_create
		each.Updated_at = date_update

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
	mydatas := []models.SettingGrupPrivilege{}
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
		"paginator"		:    paginator,
		"posts"			: mydatas,
		"search"		:       search,
		"searchStatus"	: searchStatus}


	return c.Render(http.StatusOK, "list_setting_grup_privilege", data)
}

func AddSettingGrupPrivilege(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grupprivilege_1'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grupprivilege_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	rowsPrivilege, errPrivilege := db.Query("SELECT id_setting_privilege, name_menu,  kode_privilege, kode_permissions, permissions FROM v_get_privilege GROUP BY id_setting_privilege ORDER BY kode_permissions ASC")
	if errPrivilege != nil {
		fmt.Println(errPrivilege)
	}

	defer rowsPrivilege.Close()

	eachPrivilege   := models.SettingPrivilege{}
	resultPrivilege := []models.SettingPrivilege{}

	for rowsPrivilege.Next() {
		var id_setting_privilege, 
			name_menu, 
			kode_privilege,
			kode_permissions,
			permissions []byte

		var err = rowsPrivilege.Scan(&id_setting_privilege, &name_menu, &kode_privilege, &kode_permissions, &permissions)
		if err != nil {
			fmt.Println(errPrivilege)
		}

				//get hak aksess
				rows, err2 := db.Query("SELECT permissions, kode_permissions FROM v_get_privilege WHERE id_setting_privilege = ?", string(id_setting_privilege))
				if err2 != nil {
					fmt.Println(err2)
				}
				defer rows.Close()

				eachHakAkses := models.MasterPermissions{}
				resultHakAkses := []models.MasterPermissions{}

				for rows.Next() {

					var permissions,
						kode_permissions []byte

					err2 := rows.Scan(&permissions, &kode_permissions)
					if err2 != nil {
						fmt.Println(err2)
					}

					hak_akses := ""
					if string(kode_permissions) == string(kode_privilege)+"_1"{
						hak_akses = "Create"
					}else if string(kode_permissions) == string(kode_privilege)+"_2"{
						hak_akses = "Read/View"
					}else if string(kode_permissions) == string(kode_privilege)+"_3"{
						hak_akses = "Edit"
					}else if string(kode_permissions) == string(kode_privilege)+"_4"{
						hak_akses = "Delete"
					}

					eachHakAkses.Key 		 = string(permissions)
					eachHakAkses.Value  	 = hak_akses
					eachHakAkses.Additional  = string(kode_permissions)

					resultHakAkses = append(resultHakAkses, eachHakAkses)
				}

		eachPrivilege.Id 	 			= string(id_setting_privilege)
		eachPrivilege.Name_Menu  		= string(name_menu)
		eachPrivilege.Kode_Privilege 	= string(kode_privilege)
		eachPrivilege.Kode_Permissions 	= string(kode_permissions)
		eachPrivilege.Permissions 		= resultHakAkses

		resultPrivilege = append(resultPrivilege, eachPrivilege)
	}


	resultUserGrup, _ := GetGrupPrivilege()

	data = pongo2.Context{
		"user_grup"		  : resultUserGrup,
		"resultPrivilege" : resultPrivilege,
	}

	return c.Render(http.StatusOK, "add_setting_grup_privilege", data)
}

func StoreSettingGrupPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grupprivilege_1'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grupprivilege_1"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege


	id_setting_grup := c.FormValue("id_setting_grup")
	remarks 		:= c.FormValue("remarks")
	status 			:= c.FormValue("status")

	t := time.Now()
	current_time := t.Format("2006-01-02 15:04:05")

    // Insert ke db tb_purchase_order_out
	sql := "INSERT INTO tb_setting_grup_privilege(id_setting_grup, remarks, status, created_at) VALUES(?, ?, ?, ?)"
	stmt, err := db.Prepare(sql)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	res, err2 := stmt.Exec(id_setting_grup, remarks, status, current_time)

	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	sum_last_id, _ := res.LastInsertId()
	get_last_id := strconv.FormatInt(int64(sum_last_id), 10)


	// insert permissions
	form, err := c.MultipartForm()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	permissions := form.Value["hak_akses[]"]

	for _, value := range permissions {
		sql := "INSERT INTO tb_setting_grup_privilege_detail(id_setting_grup_privilege, kode_permissions, created_at) VALUES(?, ?, ?)"
		insert, err := db.Prepare(sql)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}
		defer insert.Close()
		_, err2 := insert.Exec(get_last_id, value,  current_time)
		if err2 != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}		
	}

	return c.Redirect(301, "/lib/setting/grup_privilege/")
}

func EditSettingGrupPrivilege(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grupprivilege_3'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grupprivilege_3"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege


	requested_id := c.Param("id")

	var id_setting_grup, status, remarks []byte

	err := db.QueryRow("SELECT id_setting_grup, status, remarks FROM tb_setting_grup_privilege WHERE md5(id) = ?", requested_id).Scan(&id_setting_grup, &status, &remarks)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	response := models.SettingGrupPrivilege{
		Id:             		requested_id,
		Id_setting_grup:        string(id_setting_grup),
		Status:       	  		string(status),
		Keterangan:     		string(remarks)}


	// get permissions menu
	rowsPrivilege, errPrivilege := db.Query("SELECT id_setting_privilege, name_menu,  kode_privilege, kode_permissions, permissions FROM v_get_privilege GROUP BY id_setting_privilege ORDER BY kode_permissions ASC")
	if errPrivilege != nil {
		fmt.Println(errPrivilege)
	}

	defer rowsPrivilege.Close()

	eachPrivilege   := models.SettingPrivilege{}
	resultPrivilege := []models.SettingPrivilege{}

	for rowsPrivilege.Next() {
		var id_setting_privilege, 
			name_menu, 
			kode_privilege,
			kode_permissions,
			permissions []byte

		var err = rowsPrivilege.Scan(&id_setting_privilege, &name_menu, &kode_privilege, &kode_permissions, &permissions)
		if err != nil {
			fmt.Println(errPrivilege)
		}

				//get hak aksess
				rows, err2 := db.Query("SELECT permissions, kode_permissions FROM v_get_privilege WHERE id_setting_privilege = ?", string(id_setting_privilege))
				if err2 != nil {
					fmt.Println(err2)
				}
				defer rows.Close()

				eachHakAkses := models.MasterPermissions{}
				resultHakAkses := []models.MasterPermissions{}

				for rows.Next() {

					var key,
						kode_permissions []byte

					err2 := rows.Scan(&key, &kode_permissions)
					if err2 != nil {
						fmt.Println(err2)
					}
					// check or unchecked
					var checked_or_unchecked []byte
					err := db.QueryRow("SELECT kode_permissions FROM tb_setting_grup_privilege_detail WHERE md5(id_setting_grup_privilege) = ? AND kode_permissions = ?", requested_id, string(kode_permissions)).Scan(&checked_or_unchecked)
					if err != nil {
						fmt.Println(err)
					}
					permissions := ""
					if string(checked_or_unchecked) == ""{
						permissions = "unchecked"
					}else if string(checked_or_unchecked) != ""{
						permissions = "checked"
					}


					hak_akses := ""
					if string(kode_permissions) == string(kode_privilege)+"_1"{
						hak_akses = "Create"
					}else if string(kode_permissions) == string(kode_privilege)+"_2"{
						hak_akses = "Read/View"
					}else if string(kode_permissions) == string(kode_privilege)+"_3"{
						hak_akses = "Edit"
					}else if string(kode_permissions) == string(kode_privilege)+"_4"{
						hak_akses = "Delete"
					}

					eachHakAkses.Key 		     = string(key)
					eachHakAkses.Value  	     = hak_akses
					eachHakAkses.Additional  	 = string(kode_permissions)
					eachHakAkses.CheckOrUncheck  = permissions

					resultHakAkses = append(resultHakAkses, eachHakAkses)
				}

		eachPrivilege.Id 	 			= string(id_setting_privilege)
		eachPrivilege.Name_Menu  		= string(name_menu)
		eachPrivilege.Kode_Privilege 	= string(kode_privilege)
		eachPrivilege.Kode_Permissions 	= string(kode_permissions)
		eachPrivilege.Permissions 		= resultHakAkses

		resultPrivilege = append(resultPrivilege, eachPrivilege)
	}


	// get grup
	rowsUserGrup, errUserGrup := db.Query(" SELECT id, name_grup, status FROM v_get_setting_grup WHERE status = 'Y' ORDER BY name_grup")
	if errUserGrup != nil {
		fmt.Println(errUserGrup)
	}

	defer rowsUserGrup.Close()

	eachUserGrup := models.SettingUserGrup{}
	resultUserGrup := []models.SettingUserGrup{}

	for rowsUserGrup.Next() {
		var id, name_grup, status []byte

		var err = rowsUserGrup.Scan(&id, &name_grup, &status)

		if err != nil {
			fmt.Println(errUserGrup)
		}

		eachUserGrup.Id = string(id)
		eachUserGrup.Id_setting_grup = string(name_grup)
		eachUserGrup.Status = string(status)

		resultUserGrup = append(resultUserGrup, eachUserGrup)
	}


	data = pongo2.Context{
		"user_grup"		  		: resultUserGrup,
		"setting_grup_privilege": response,
		"resultPrivilege"		: resultPrivilege,
	}
	return c.Render(http.StatusOK, "edit_setting_grup_privilege", data)
}

func UpdateSettingGrupPrivilege(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grupprivilege_3'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grupprivilege_3"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	t := time.Now()
	current_time := t.Format("2006-01-02 15:04:05")
	
	
	emp := new(models.SettingGrupPrivilege)
	if err := c.Bind(emp); err != nil {
		return err
	}
	id := c.Param("id")

	id_setting_grup := c.FormValue("id_setting_grup")
	status := c.FormValue("status")
	remarks := c.FormValue("remarks")

	selDB, err2 := db.Prepare("UPDATE tb_setting_grup_privilege SET id_setting_grup=?, status=?, remarks=?, updated_at = ? WHERE md5(id)=?")
	if err2 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	defer selDB.Close()
	selDB.Exec(id_setting_grup, status, remarks, current_time, id)


	// get id_setting_grup_privilege 
	var id_setting_grup_privilege []byte
	err := db.QueryRow("SELECT id FROM tb_setting_grup_privilege WHERE md5(id) = ?", id).Scan(&id_setting_grup_privilege)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	fmt.Println(string(id_setting_grup_privilege))

	// Delete
	sql2 := "DELETE FROM tb_setting_grup_privilege_detail WHERE md5(id_setting_grup_privilege) = ?"
	stmt, err := db.Prepare(sql2)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	defer stmt.Close()
	_, err3 := stmt.Exec(id)
	if err3 != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	// update permissions
	form, err := c.MultipartForm()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}
	permissions := form.Value["hak_akses[]"]

	for _, value := range permissions {

		sql := "INSERT INTO tb_setting_grup_privilege_detail(id_setting_grup_privilege, kode_permissions, created_at) VALUES(?, ?, ?)"
		insert, err := db.Prepare(sql)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}
		defer insert.Close()
		_, err2 := insert.Exec(string(id_setting_grup_privilege), value,  current_time)
		if err2 != nil {
			return c.Render(http.StatusInternalServerError, "error_500", nil)
		}		
	}

	return c.Redirect(301, "/lib/setting/grup_privilege/")
}

func ShowSettingGrupPrivilege(c echo.Context) error {

	db := database.CreateCon()
	defer db.Close()

	cc := &MyCustomContext{c}
	data_users			:= cc.getDataLogin()

	// check_privilege
	var check_privilege []byte
	check_privileges, errPriv := db.Prepare("SELECT kode_permissions FROM v_get_grup_privilege_detail WHERE id_setting_grup = ? AND kode_permissions = 'setting.user.grupprivilege_2'")
	if errPriv != nil {
		fmt.Printf("%s", errPriv)
	}
	errPriv = check_privileges.QueryRow(data_users.Id_group).Scan(&check_privilege)	
	defer check_privileges.Close()
	if string(check_privilege) != "setting.user.grupprivilege_2"{
		return c.Render(http.StatusInternalServerError, "error_403", nil)
	}
	//end check_privilege

	requested_id := c.Param("id")

	var name_grup, status, remarks []byte

	err := db.QueryRow("SELECT name_grup, status, remarks FROM v_get_grup_privilege WHERE md5(id) = ?", requested_id).Scan(&name_grup, &status, &remarks)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "error_500", nil)
	}

	response := models.SettingGrupPrivilege{
		Id:             		requested_id,
		Name_grup:        string(name_grup),
		Status:       	  		string(status),
		Keterangan:     		string(remarks)}


	// get permissions menu
	rowsPrivilege, errPrivilege := db.Query("SELECT id_setting_privilege, name_menu,  kode_privilege, kode_permissions, permissions FROM v_get_privilege GROUP BY id_setting_privilege ORDER BY kode_permissions ASC")
	if errPrivilege != nil {
		fmt.Println(errPrivilege)
	}

	defer rowsPrivilege.Close()

	eachPrivilege   := models.SettingPrivilege{}
	resultPrivilege := []models.SettingPrivilege{}

	for rowsPrivilege.Next() {
		var id_setting_privilege, 
			name_menu, 
			kode_privilege,
			kode_permissions,
			permissions []byte

		var err = rowsPrivilege.Scan(&id_setting_privilege, &name_menu, &kode_privilege, &kode_permissions, &permissions)
		if err != nil {
			fmt.Println(errPrivilege)
		}

				//get hak aksess
				rows, err2 := db.Query("SELECT permissions, kode_permissions FROM v_get_privilege WHERE id_setting_privilege = ?", string(id_setting_privilege))
				if err2 != nil {
					fmt.Println(err2)
				}
				defer rows.Close()

				eachHakAkses := models.MasterPermissions{}
				resultHakAkses := []models.MasterPermissions{}

				for rows.Next() {

					var key,
						kode_permissions []byte

					err2 := rows.Scan(&key, &kode_permissions)
					if err2 != nil {
						fmt.Println(err2)
					}
					// check or unchecked
					var checked_or_unchecked []byte
					err := db.QueryRow("SELECT kode_permissions FROM tb_setting_grup_privilege_detail WHERE md5(id_setting_grup_privilege) = ? AND kode_permissions = ?", requested_id, string(kode_permissions)).Scan(&checked_or_unchecked)
					if err != nil {
						fmt.Println(err)
					}
					permissions := ""
					if string(checked_or_unchecked) == ""{
						permissions = "unchecked"
					}else if string(checked_or_unchecked) != ""{
						permissions = "checked"
					}


					hak_akses := ""
					if string(kode_permissions) == string(kode_privilege)+"_1"{
						hak_akses = "Create"
					}else if string(kode_permissions) == string(kode_privilege)+"_2"{
						hak_akses = "Read/View"
					}else if string(kode_permissions) == string(kode_privilege)+"_3"{
						hak_akses = "Edit"
					}else if string(kode_permissions) == string(kode_privilege)+"_4"{
						hak_akses = "Delete"
					}

					eachHakAkses.Key 		     = string(key)
					eachHakAkses.Value  	     = hak_akses
					eachHakAkses.Additional  	 = string(kode_permissions)
					eachHakAkses.CheckOrUncheck  = permissions

					resultHakAkses = append(resultHakAkses, eachHakAkses)
				}

		eachPrivilege.Id 	 			= string(id_setting_privilege)
		eachPrivilege.Name_Menu  		= string(name_menu)
		eachPrivilege.Kode_Privilege 	= string(kode_privilege)
		eachPrivilege.Kode_Permissions 	= string(kode_permissions)
		eachPrivilege.Permissions 		= resultHakAkses

		resultPrivilege = append(resultPrivilege, eachPrivilege)
	}


	// get grup
	rowsUserGrup, errUserGrup := db.Query(" SELECT id, name_grup, status FROM v_get_setting_grup WHERE status = 'Y' ORDER BY name_grup")
	if errUserGrup != nil {
		fmt.Println(errUserGrup)
	}

	defer rowsUserGrup.Close()

	eachUserGrup := models.SettingUserGrup{}
	resultUserGrup := []models.SettingUserGrup{}

	for rowsUserGrup.Next() {
		var id, name_grup, status []byte

		var err = rowsUserGrup.Scan(&id, &name_grup, &status)

		if err != nil {
			fmt.Println(errUserGrup)
		}

		eachUserGrup.Id = string(id)
		eachUserGrup.Id_setting_grup = string(name_grup)
		eachUserGrup.Status = string(status)

		resultUserGrup = append(resultUserGrup, eachUserGrup)
	}


	data = pongo2.Context{
		"user_grup"		  		: resultUserGrup,
		"setting_grup_privilege": response,
		"resultPrivilege"		: resultPrivilege,
	}
	return c.Render(http.StatusOK, "show_setting_grup_privilege", data)
}

