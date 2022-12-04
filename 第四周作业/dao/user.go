package dao

import (
	"LanShan/Week_04/model"
)

func Adduser(u *model.User) bool {
	sqlStr := "insert into user(username,password) values (?,?)"
	_, err := db.Exec(sqlStr, u.Username, u.Password)
	if err != nil {
		return false
	}
	return true
}

func SelectUser(userId int64) bool {
	sqlStr := "select id, username, password from user where id = ?"
	rows, err := db.Query(sqlStr, userId)
	if err != nil {
		panic(err)
	}
	rows.Next()
	var u model.User
	err = rows.Scan(&u.UserId, &u.Username, &u.Password)
	if err != nil {
		return false
	}
	if u.Username == "" {
		return false
	}
	return true
}

func SelectPasswordFromUserId(userId int64) (username string, password string) {
	sqlStr := "select id, username, password from user where id = ?"
	rows, err := db.Query(sqlStr, userId)
	if err != nil {
		return "", ""
	}
	rows.Next()
	var u model.User
	err = rows.Scan(&u.UserId, &u.Username, &u.Password)
	if err != nil {
		return "", ""
	}
	return u.Username, u.Password
}
