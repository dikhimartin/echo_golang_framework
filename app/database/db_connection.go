package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "db_receipt_golang_crud"
    DB_USER = "root"
    DB_PASS = "root"
)

/*Create mysql connection*/
func CreateCon() *sql.DB {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", DB_USER, DB_PASS, DB_HOST, DB_NAME))
	if err != nil {
		fmt.Println(err)
	}

	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println("db is not connected")
	}

	return db
}
