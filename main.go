package main

import (
	// "fmt"
	// "main/utils"
	controller "main/controllers"
	"main/di"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	di.InitDB()

	//

	controller.UserApi(router)

	controller.ProductApi(router)
	controller.CartApi(router)

	//
	router.Run(":8080")
}


