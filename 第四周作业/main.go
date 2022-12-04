package main

import (
	"LanShan/Week_04/api"
	"LanShan/Week_04/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
