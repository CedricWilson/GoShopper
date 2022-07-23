package service

import (
	"fmt"
	"main/di"
	"main/models"
	"main/utils"
	"strings"
	"time"

	// "strings"
	// "time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// type encrypt struct {
// 	hasher utils.Hasher
// }

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
	// start1 := time.Now()
	if di.Db.Where("email = ?", login.Email).First(&cust).Error != nil {
		utils.Failure(c, "No user found!")
		return
	}
	// elapsed1 := time.Since(start1)
	// fmt.Println("DB: "+shortDur(elapsed1))

	// pass := encrypt{}
	// pass.hasher.CheckPasswordHash(login.Password, cust.Password)

	// start2 := time.Now()
	if !utils.CheckPasswordHash(login.Password, cust.Password) {
		utils.Failure(c, "Incorrect Credentials")
		return
	}
	// elapsed2 := time.Since(start2)
	// fmt.Println("Hash: "+shortDur(elapsed2))

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

	// fmt.Println(user.First_name)
	// fmt.Println(val)

	di.Db.Exec("call updateUser(?, ?)", user.First_name, val)

	utils.Success(c, "Succedd")

}

func VerifyUser(tokenString string) (string, interface{}) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(("your-256-bit-secret")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["userId"])

		return fmt.Sprintf("%v", claims["userId"]), nil
	} else {
		// fmt.Println(err)
		return "", err

	}

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

	// fmt.Println(slice)
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

// func shortDur(d time.Duration) string {
//     s := d.String()
//     if strings.HasSuffix(s, "m0s") {
//         s = s[:len(s)-2]
//     }
//     if strings.HasSuffix(s, "h0m") {
//         s = s[:len(s)-2]
//     }
//     return s
// }
