package service

import (
	"github.com/gin-gonic/gin"
	"main/di"
	"main/models"
	"main/utils"
)

func ListProducts(c *gin.Context) {

	slice := []models.Product{}

	if di.Db.Find(&slice).Error != nil {
		utils.Failure(c, "No Products Available")
		return
	}

	utils.Success(c, slice)

}
