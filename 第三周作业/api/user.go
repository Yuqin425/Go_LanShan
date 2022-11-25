package api

import (
	"LanShan/Week_03/api/middleware"
	"LanShan/Week_03/dao"
	"LanShan/Week_03/model"
	"LanShan/Week_03/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFailed(c, "请求错误")
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")

	flag := dao.SelectUser(username)

	if flag {
		utils.RespFailed(c, "此账号已存在")
		return
	}
	if len(telephone) < 11 {
		utils.RespFailed(c, "手机号必须为11位")
		return
	}
	dao.AddUser(username, password)
	dao.AddTel(username, telephone)
	utils.RespSuccess(c, "注册账号成功")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFailed(c, "请求错误")
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")

	flag := dao.SelectUser(username)

	if !flag {
		utils.RespFailed(c, "此账号不存在")
		return
	}

	selectPassword := dao.SelectPasswordFromUsername(username)

	if selectPassword != password {
		utils.RespFailed(c, "密码错误")
		return
	}

	claim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "ZYQ",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, "登陆成功"+telephone+tokenString)
}

func changepwd(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFailed(c, "请求错误")
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	newpassword := c.PostForm("newpassword")

	flag := dao.SelectUser(username)

	if !flag {
		utils.RespFailed(c, "此账号不存在")
		return
	}

	selectPassword := dao.SelectPasswordFromUsername(username)

	if selectPassword != password {
		utils.RespFailed(c, "原密码错误")
		return
	}

	dao.DeleteUser(username)
	dao.AddUser(username, newpassword)
	utils.RespSuccess(c, "修改密码成功")
}

func findpwd(c *gin.Context) {
	username := c.PostForm("username")
	telephone := c.PostForm("telephone")
	newpassword := c.PostForm("newpassword")

	flag := dao.SelectUser(username)

	if !flag {
		utils.RespFailed(c, "此账号不存在")
		return
	}

	SelectTel := dao.SelectTelephoneFromUsername(username)

	if SelectTel != telephone {
		utils.RespFailed(c, "验证失败")
		return
	}

	dao.DeleteUser(username)
	dao.AddUser(username, newpassword)
	utils.RespSuccess(c, "找回密码成功")
}

func delete(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFailed(c, "请求错误")
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")

	flag := dao.SelectUser(username)

	if !flag {
		utils.RespFailed(c, "此账号不存在")
		return
	}

	selectPassword := dao.SelectPasswordFromUsername(username)
	if selectPassword != password {
		utils.RespFailed(c, "密码错误")
		return
	}

	selecctTel := dao.SelectTelephoneFromUsername(username)
	if selecctTel != telephone {
		utils.RespFailed(c, "验证失败")
	}

	dao.DeleteUser(username)
	dao.DeleteTel(username)
	utils.RespSuccess(c, "注销账号成功")
}

func Comment(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFailed(c, "请输入账号和密码")
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	comment := c.PostForm("comment")

	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFailed(c, "此账号不存在")
		return
	}

	selectPassword := dao.SelectPasswordFromUsername(username)

	if selectPassword != password {
		utils.RespFailed(c, "密码错误")
		return
	}

	dao.AddComment(comment)
	utils.RespSuccess(c, "评论成功")
}
