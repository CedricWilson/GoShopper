package service

import (
	// "fmt"
	"main/di"
	"main/models"
	"main/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func FetchCart(c *gin.Context) {
	userId, err1 := VerifyUser(strings.Split(c.GetHeader("Authorization"), " ")[1])
	if err1 != nil {
		utils.Failure(c, err1)
		return
	}

	// slice := []models.Cart{}

	// di.Db.Joins("JOIN products ON carts.product_id = products.product_id WHERE carts.customer_id = ?", &userId).Find(&slice)
	//
	// mod := &models.FetchCartDTO{}
	slice := []models.FetchCartDTO{}

	// res :=  di.Db.Raw("select carts.cart_id,products.name,carts.quantity FROM carts JOIN products ON carts.product_id = products.product_id WHERE carts.customer_id = ?", userId)
	//

	di.Db.Raw("call fetchCart(?)", userId).Scan(&slice)

	total := 0.00

	for _, e := range slice {
		total += e.Price
	}

	response := map[string]any{
		"products": slice,
		"total":    total,
	}

	utils.Success(c, response)

}
func AddToCart(c *gin.Context) {

	userId, err1 := VerifyUser(strings.Split(c.GetHeader("Authorization"), " ")[1])
	if err1 != nil {
		utils.Failure(c, err1)
		return
	}

	var addProduct models.Cart

	if err := c.ShouldBind(&addProduct); err != nil {
		utils.Failure(c, err)
		return
	}
	addProduct.Customer_id = userId

	res := di.Db.Create(&addProduct)

	if res.Error != nil {
		utils.Failure(c, res.Error)
		return
	}

	utils.Success(c, "Success")

}
