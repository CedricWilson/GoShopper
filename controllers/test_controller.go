package controller

import (
	"github.com/gin-gonic/gin"
	"main/service"
)

func TestApi(router *gin.Engine) {
	api := router.Group("/test")
	api.GET("/initial", service.Test)
}
