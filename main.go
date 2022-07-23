package main

import (
	// "fmt"
	// "main/utils"
	"main/controllers"
	"main/di"
)

func main() {
	 di.InitilizeServer()

	//

	controller.UserApi(di.Router)
	controller.TestApi(di.Router)

	//
	di.StartServer()
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
