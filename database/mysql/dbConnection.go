package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Sqlclient() *sql.DB {
	dbDriver := "mysql"
	dbUser := "task"
	dbPass := "12345"
	dbName := "EmployeeManagementSystem"
	dbPort := 3309
	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:"+dbPort+")/"+dbName+"?parseTime=True")
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(0.0.0.0:%v)/%s", dbUser, dbPass, dbPort, dbName))
	if err != nil {
		log.Fatal("error connecting DB : ", err.Error())
	}
	return db
}
