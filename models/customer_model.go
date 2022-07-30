package models

// import "time"

type Customer struct {
	Customer_id int    `gorm:"primaryKey" json:"customer_id"`
	First_name  string `json:"first_name" gorm:"type:varchar(255)"`
	Last_name   string `json:"last_name" gorm:"type:varchar(255)"`
	Email       string `json:"email" gorm:"type:varchar(255)"`
	Birth_date  string `gorm:"type:time" json:"dob"`
	Phone       string `json:"phone" gorm:"type:varchar(255)"`
	Address     string `json:"address" gorm:"type:varchar(255)"`
	City        string `json:"city" gorm:"type:varchar(255)"`
	State       string `json:"state" gorm:"type:varchar(255)"`
	Points      int    `json:"points" gorm:"type:int"`
	Password    string `json:"-"`
	Token       string `json:"token"`
}

type Login struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type Register struct {
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required,email"`
	Phone      string `json:"phone" binding:"required"`
	Address    string `json:"address" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
}

type UpdateUserNameDTO struct {
	First_name string `json:"first_name"  binding:"required"`
}
