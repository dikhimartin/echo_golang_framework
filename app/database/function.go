package database

import(
	"fmt"
	lib   "receipt/lib"
)

func ChecktableRecord(table_name string) bool{
	db := CreateCon()
	defer db.Close()
	var result int
	db.Table(table_name).Count(&result)
	if result == 0{
		return false
	}
	return true
}

func ChecktableExist(table_name string) bool{
	db := CreateCon()
	defer db.Close()

	var table_check int
	row := db.Table(table_name).Select("COUNT(*)").Row() 
	row.Scan(&table_check)

    if table_check > 0 {
        fmt.Println(lib.Warn("Table '"+table_name+"' already exists"))
        return false
    } else {
        fmt.Println(lib.Info("Created migration view ", table_name))
        return true
    }

	return true
}


func DropView(view_name string) bool{
	db := CreateCon()
	defer db.Close()

	if err := db.Exec("DROP VIEW `"+ view_name +"`;").Error; err != nil {
	  // error handling...
	  fmt.Println(err)
	  return false
	}
	return true
}

func DropViewIfExist(view_name string) bool{
	db := CreateCon()
	defer db.Close()

	if err := db.Exec("DROP VIEW IF EXISTS `"+ view_name +"`;").Error; err != nil {
	  // error handling...
	  fmt.Println(err)
	  return false
	}

	return true
}
