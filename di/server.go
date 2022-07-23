package di

import (
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func InitilizeServer() {
	Router = gin.Default()
	InitDB()
}
func StartServer() {
	Router.Run()
}
