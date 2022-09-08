package controller

import (
	"main/models"
	service "main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	// Login(c *gin.Context)
	Register(c *gin.Context)
}
type authController struct {
	authService service.AuthService
	// jwtService  service.JWTservice
}

func NewAuthController(auth service.AuthService) AuthController {
	return &authController{
		authService: auth,
		// jwtService:  jwt,
	}
}

func (auth *authController) Register(c *gin.Context) {
	var register models.Register
	err := c.ShouldBind(&register)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	message, err1 := auth.authService.Register(register)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err1.Error())
		return
	}

	c.JSON(http.StatusCreated, message)

}
