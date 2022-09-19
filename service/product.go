package service

import (
	"github.com/gin-gonic/gin"
	"goshopper/di"
	"goshopper/models"
	"goshopper/utils"
)

func ListProducts(c *gin.Context) {

	slice := []models.Product{}

	if di.Db.Find(&slice).Error != nil {
		utils.Failure(c, "No Products Available")
		return
	}

	utils.Success(c, slice)

}
