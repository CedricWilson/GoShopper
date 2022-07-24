package di

import (
	"fmt"
	"main/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDB() {
	dsn := "root:rootroot@tcp(localhost:3306)/sql_store?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Database connected successfully!")
		db.AutoMigrate(&models.Customer{})
		Db = db
	}
}