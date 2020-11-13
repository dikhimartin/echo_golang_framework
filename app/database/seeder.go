package database

import(
	"strconv"
	"encoding/json"
	"../application/models"
)

// data_seeder_here
func SeedPermission() []models.SchemePermission{
	concept := 	`[{"id":"1","name":"Create","additional":null}, {"id":"2","name":"Read\/View","additional":null}, {"id":"3","name":"Edit","additional":null}, {"id":"4","name":"Delete","additional":null}]`
	var jsonData = []byte(concept)
	var data_concept []models.SchemePermission
	err := json.Unmarshal(jsonData, &data_concept)
	if err != nil {
		logs.Println(err)
		panic(err)
	}
	return data_concept
}


// import_data
func DataSeeder(){
	db := CreateCon()
	defer db.Close()
	tb_permission 		 := SeedPermission()
	if ChecktableRecord("tb_permission") == false{
		for _, v := range tb_permission {
			id, _ := strconv.Atoi(v.ID)
			seed := models.Permission{
				ID    : id,
				Name  : v.Name,
			}
			if error_insert := db.Create(&seed); error_insert.Error != nil {
				logs.Println(error_insert)
				panic(error_insert)
			}
			db.NewRecord(seed)
		}
	}
}


