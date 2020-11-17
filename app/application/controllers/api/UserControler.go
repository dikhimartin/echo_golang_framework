package api

import(
	"time"
	controllers ".."
	"../../../database"
	"github.com/labstack/echo"
)

func Getinfologin(c echo.Context) error{
	data_users	:= controllers.GetDataLogin(c)
	response := response_json{
		"data_users"				:   data_users,
		"time"						:   time.Now().UnixNano(),
	}
	return c.JSON(200, response)
}


func CheckUsername(c echo.Context) error{
	db := database.CreateCon()
	defer db.Close()

	type_check     := c.FormValue("type_check")
	username   	   := c.FormValue("username")
	username_old   := c.FormValue("username_old")

	var check_username []byte
	if type_check == "1"{
		//cek_data_add_new
		row := db.Table("v_get_user").Where("username = ?", username).Select("username").Row()
		err := row.Scan(&check_username)
		if err != nil{
			logs.Println(err)
		}
	}else if type_check == "2"{
		//cek_data_edit
		row := db.Table("v_get_user").Where("username= ? AND username != ?", username, username_old).Select("username").Row()
		err := row.Scan(&check_username)
		if err != nil{
			logs.Println(err)
		}
	}

	if string(check_username) != "" {
		response := map[string]string{"alert": "Maaf, Username sudah digunakan!", "kode": "1"}
		return c.JSON(200, response)
	} else {
		response := map[string]string{"alert": "Sukses, Username belum digunakan!", "kode": "0"}
		return c.JSON(200, response)
	}	
}
