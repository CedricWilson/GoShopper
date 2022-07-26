package di

import (
	"fmt"
	"goshopper/models"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
	
    "github.com/rs/zerolog/log"


)

var (
	Db *gorm.DB
	Redis *redis.Client
)

func InitDB() {
	dsn := "root:rootroot@tcp(localhost:3306)/sql_store?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error().Msg(err.Error())
		
		os.Exit(1)
	} else {
		fmt.Println("Database connected successfully!")
		db.AutoMigrate(&models.Customer{})
		Db = db
	}
}

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	Redis = rdb

}
