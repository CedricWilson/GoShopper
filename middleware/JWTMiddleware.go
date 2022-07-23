package middleware

import (
	"fmt"
	"main/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("--Before JWT--")
		
		tokenString := strings.Split(c.GetHeader("Authorization"), " ")[1]
		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			fmt.Println("Valid")
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println("InValid")

			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
