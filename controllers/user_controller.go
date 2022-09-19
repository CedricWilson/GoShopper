package controller

import (
	"goshopper/service"

	"github.com/gin-gonic/gin"
)

func UserApi(router *gin.Engine) {
	api := router.Group("/user")
	api.GET("/getList", service.GetUsers)
	api.GET("/getAll", service.GetAll)
	api.POST("/login", service.Login)
	api.POST("/register", service.Register)
	api.POST("/update", service.UpdateUserName)

	

}
