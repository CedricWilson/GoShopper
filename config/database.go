package config

import (
	"fmt"
	// "github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/models"
	"os"
)

func InitDB() *gorm.DB {
	dsn := "root:rootroot@tcp(localhost:3306)/sql_store?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Database connected successfully!")
		db.AutoMigrate(&models.Customer{})

	}
	return db
}

// func InitRedis() {
// 	rdb := redis.NewClient(&redis.Options{
//         Addr:     "127.0.0.1:6379",
//         Password: "", // no password set
//         DB:       0,  // use default DB
//     })
// 	Redis = rdb

// }
