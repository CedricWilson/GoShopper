package controller

import (
	"main/middleware"
	"main/service"

	"github.com/gin-gonic/gin"
)

var (
	jwtService service.JWTService = service.JWTAuthService()
	userService service.UserService = service.CreateUserService(jwtService)
	
)

func UserApi(router *gin.Engine) {
	api := router.Group("/user")

	api.GET("/getList", service.GetUsers)
	api.GET("/getAll", service.GetAll)
	api.POST("/login", service.Login)
	api.POST("/register", service.Register)

	routes := router.Group("/modify", middleware.AuthorizeJWT())
	{
		routes.POST("/update", userService.UpdateUserName)

	}

}
