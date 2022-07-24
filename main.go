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

// func main(){
// 	password := "Abcd@1234"
//     hash, _ := utils.HashPassword(password) // ignore error for the sake of simplicity

//     fmt.Println("Password:", password)
//     fmt.Println("Hash:    ", hash)
//     fmt.Println("Length:    ", len(hash))

//     match := utils.CheckPasswordHash(password, hash)
//     fmt.Println("Match:   ", match)
// }
