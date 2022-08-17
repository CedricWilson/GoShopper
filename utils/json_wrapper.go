package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, val any) {

	response := map[string]any{
		"error":  "0",
		"result": val,
	}

	c.JSON(http.StatusOK, response)
	c.Abort()

}

func Failure(c *gin.Context, val any) {
	response := map[string]any{
		"error":  "1",
		"result": val,
	}
	c.JSON(http.StatusInternalServerError, response)
	c.Abort()
}
