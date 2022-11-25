package api

import (
	"LanShan/Week_03/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/changepwd", changepwd)
	r.POST("/findpwd", findpwd)
	r.POST("/delete", delete)
	r.POST("/comment", Comment)

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.Run(":8848")
}
