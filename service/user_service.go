package service

import (
	
	"goshopper/di"
	"goshopper/models"
	"goshopper/utils"
	"strings"
	"time"

	
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Register(c *gin.Context) {
	
	var register models.Register

	err := c.ShouldBind(&register)
	if err != nil {
		utils.Failure(c, err.Error())
		return
	}

	pass, _ := utils.HashPassword(register.Password)

	cust := models.Customer{
		First_name: register.First_name,
		Last_name:  register.Last_name,
		Email:      register.Email,
		Password:   pass,
		State:      register.State,
		City:       register.City,
		Phone:      register.Phone,
		Address:    register.Address,
		Birth_date: time.Now(),
		Points:     0,
	}
	res := di.Db.Create(&cust)
	if res.Error != nil {
		utils.Failure(c, res.Error)
		return
	}
	utils.Success(c, "Success")

}

func Login(c *gin.Context) {
	var login models.Login

	err := c.ShouldBind(&login)
	if err != nil {
		utils.Failure(c, err.Error())
		return
	}

	cust := models.Customer{}
	
	if di.Db.Where("email = ?", login.Email).First(&cust).Error != nil {
		utils.Failure(c, "No user found!")
		return
	}
	

	
	if !utils.CheckPasswordHash(login.Password, cust.Password) {
		utils.Failure(c, "Incorrect Credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":        cust.Customer_id,
		"generatedTime": time.Now(),
	})

	tokenString, err := token.SignedString([]byte(("your-256-bit-secret")))
	if err != nil {
		utils.Failure(c, err.Error())
		return

	}
	cust.Token = tokenString

	utils.Success(c, cust)
}

func UpdateUserName(c *gin.Context) {

	val, err1 := VerifyUser(strings.Split(c.GetHeader("Authorization"), " ")[1])
	if err1 != nil {
		utils.Failure(c, err1)
		return
	}

	var user models.UpdateUserNameDTO

	err := c.ShouldBind(&user)
	if err != nil {
		utils.Failure(c, err.Error())
		return
	}


	di.Db.Exec("call updateUser(?, ?)", user.First_name, val)

	utils.Success(c, "Succedd")

}



func GetUsers(c *gin.Context) {
	slice := models.StaticUsers()

	if c.Query("name") != "" {
		for _, e := range slice {
			if e.Name == c.Query("name") {

				utils.Success(c, e)
				return
			}
		}
		utils.Failure(c, "No Data Found")
		return
	}
	utils.Success(c, slice)

}

func GetAll(c *gin.Context) {
	slice := []models.Customer{}

	if c.Query("name") != "" {
		cust := models.Customer{}

		if di.Db.Where("first_name = ?", c.Query("name")).Take(&cust).Error != nil {
			utils.Failure(c, "No user found!")
			return
		}
		slice = append(slice, cust)
	} else {
		di.Db.Raw("call getCustomers()").Scan(&slice)
	}

	utils.Success(c, slice)
}

