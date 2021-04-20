package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/foo")
	if err != nil {
		panic(err)
	}
}
