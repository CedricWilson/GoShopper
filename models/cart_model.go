package models

type Cart struct {
	Cart_id     int `gorm:"primaryKey" json:"cart_id" `
	Customer_id int `json:"customer_id" gorm:"type:int"`
	Quantity    int `json:"quantity" gorm:"type:int" binding:"required"`
	Product_id  int `json:"product_id" gorm:"type:int" binding:"required"`
}

type FetchCartDTO struct {
	Cart_id    int     `json:"cart_id" gorm:"type:int"`
	Name       string  `json:"name" gorm:"type:varchar(255)"`
	Quantity   int     `json:"quantity" gorm:"type:int"`
	Unit_price float64 `json:"unit_price" gorm:"type:float64"`
	Price      float64 `json:"price" gorm:"type:float64"`
}

// min max number
