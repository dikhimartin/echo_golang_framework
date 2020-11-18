package database

import (
	"fmt"
	lib   "receipt/lib"
     _"github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
)

var logs  	 = lib.RecordLog("SYSTEMS -")
var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
	lib.GetEnv("DB_USER"),
	lib.GetEnv("DB_PASS"),
	lib.GetEnv("DB_HOST"),
	lib.GetEnv("DB_PORT"),
	lib.GetEnv("DB_NAME"),
)

func init() {
	// DropColumn()
	AutoMigrate()
	ViewMigrate()
	// RemoveIndex()
	// ModifyColumn()
	// AddForeignKey()
	DataSeeder()
}

func CreateCon() *gorm.DB{
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logs.Println(err)
		panic(err)
	}
	return db
}
