package controller

import (
	"main/service"

	"github.com/gin-gonic/gin"
)

func ProductApi(router *gin.Engine) {
	api := router.Group("/product")
	api.GET("/", service.ListProducts)

}
