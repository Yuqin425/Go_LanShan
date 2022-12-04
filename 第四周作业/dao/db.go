package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	dsn := "root:zyq425ZYQ@(localhost)/myDatabase?charset=utf8&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	// 打开mysql驱动
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}
