package main

import (
	"LanShan/Week_04/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func initDB() {
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

func main() {
	initDB()
	u1 := model.Student{Id: 1, Name: "a", Sex: "男", Age: 1}
	u2 := model.Student{Id: 1, Name: "b", Sex: "女", Age: 2}
	u3 := model.Student{Id: 1, Name: "c", Sex: "男", Age: 3}
	u4 := model.Student{Id: 1, Name: "d", Sex: "女", Age: 4}
	u5 := model.Student{Id: 1, Name: "e", Sex: "男", Age: 5}
	u6 := model.Student{Id: 1, Name: "f", Sex: "女", Age: 6}
	u7 := model.Student{Id: 1, Name: "g", Sex: "男", Age: 7}
	u8 := model.Student{Id: 1, Name: "h", Sex: "女", Age: 8}
	u9 := model.Student{Id: 1, Name: "i", Sex: "男", Age: 9}
	u10 := model.Student{Id: 1, Name: "j", Sex: "女", Age: 10}

	InsertStudent(u1)
	InsertStudent(u2)
	InsertStudent(u3)
	InsertStudent(u4)
	InsertStudent(u5)
	InsertStudent(u6)
	InsertStudent(u7)
	InsertStudent(u8)
	InsertStudent(u9)
	InsertStudent(u10)

	queryMultiRowDemo()
}

// InsertStudent 插入数据
func InsertStudent(st model.Student) {
	sqlStr := "insert into student(name,sex,age) values (?,?,?)"
	_, err := db.Exec(sqlStr, st.Name, st.Sex, st.Age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, sex, age from student where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u model.Student
		err := rows.Scan(&u.Id, &u.Name, &u.Sex, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s sex: %s age:%d\n", u.Id, u.Name, u.Sex, u.Age)
	}
}
