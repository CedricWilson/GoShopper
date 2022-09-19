package controller

import (
	// "fmt"
	// "fmt"
	 "goshopper/service"
	"goshopper/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CartApi(router *gin.Engine) {
	api := router.Group("/cart", handler())
	api.GET("/fetch", service.FetchCart) 
	api.POST("/add", service.AddToCart)

}

func handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.GetHeader("Authorization")
		if len(str) == 0 {
			utils.Failure(c, "Authorization is required!")
			return

		}
		userId, err := service.VerifyUser(strings.Split(c.GetHeader("Authorization"), " ")[1])
		if err != nil {
			utils.Failure(c, "Invalid Authorization!")
			return

		}
		c.Set("userId",userId)
	}
}
