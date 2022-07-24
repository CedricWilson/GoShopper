package controller

import (
	"main/service"

	"github.com/gin-gonic/gin"
)

func CartApi(router *gin.Engine) {
	api := router.Group("/cart")
	api.GET("/fetch", service.FetchCart)
	api.POST("/add", service.AddToCart)

}
