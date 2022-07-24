package models

type Product struct {
	Product_id int `gorm:"primaryKey" json:"product_id"`

	Name string `json:"name"`

	Quantity_in_stock int `json:"-"`

	Unit_price float32 `json:"price"`
}
