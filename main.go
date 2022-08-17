package main

import (
	
	controller "main/controllers"
	"main/di"
	"net/http"

	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	di.InitDB()

	di.InitRedis()
	

	//
	router.LoadHTMLFiles("static/index.tmpl.html")
	
	router.Static("/css", "static/css")
	router.Static("/images", "static/images")


	
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "GoShopper",
		})
	})

	controller.UserApi(router)

	controller.ProductApi(router)
	controller.CartApi(router)

	//
	router.Run(":8080")
}

