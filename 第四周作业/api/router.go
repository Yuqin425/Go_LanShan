package api

import (
	"LanShan/Week_04/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", Register)
	r.POST("/login", Login)

	r.Run(":7890")
}
