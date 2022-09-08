package repository

import (
	"main/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user models.Customer) (string, error)
}

type UserConnection struct {
	Conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserConnection{
		Conn: db,
	}
}

func (db *UserConnection) Register(user models.Customer) (string, error) {


	res := db.Conn.Create(&user)
	if res.Error != nil {

		return "Failure", res.Error
	}
	return "Success", nil

}
