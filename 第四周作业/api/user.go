package api

import (
	"LanShan/Week_04/dao"
	"LanShan/Week_04/model"
	"LanShan/Week_04/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Register(c *gin.Context) {
	userid := c.PostForm("userid")
	username := c.PostForm("username")
	password := c.PostForm("password")

	var id int64
	id, _ = strconv.ParseInt(userid, 10, 64)
	//检查用户是否存在
	flag := dao.SelectUser(id)
	if flag {
		utils.RespFailed(c, "该用户已存在")
		return
	}

	// 创建新用户
	var newuser = model.User{
		UserId:   id,
		Username: username,
		Password: password,
	}
	flag = dao.Adduser(&newuser)
	if !flag {
		utils.RespFailed(c, "注册用户失败")
		return
	}
	utils.RespSuccess(c, "注册用户成功")
}

func Login(c *gin.Context) {
	userid := c.PostForm("userid")
	username := c.PostForm("username")
	password := c.PostForm("password")

	var id int64
	id, _ = strconv.ParseInt(userid, 10, 64)
	//检查用户是否存在
	flag := dao.SelectUser(id)
	if !flag {
		utils.RespFailed(c, "该用户不存在")
		return
	}

	Selectusername, Selectpwd := dao.SelectPasswordFromUserId(id)
	if Selectusername != username {
		utils.RespFailed(c, "用户名不正确")
		return
	}
	if Selectpwd != password {
		utils.RespFailed(c, "密码错误")
		return
	}
	utils.RespSuccess(c, "登录成功")
}
