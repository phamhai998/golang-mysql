package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var err error

func ConnectDB() {
	DB, err = sqlx.Connect("mysql", "root:phamhai@tcp(127.0.0.1:3307)/golang-mysql")
	if err != nil {
		log.Fatalln(err)
	}
	// defer DB.Close() khong duoc dung cai nay
}
