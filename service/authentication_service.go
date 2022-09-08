package service

import (
	"main/models"
	"main/repository"
	"main/utils"
	"time"
)

type AuthService interface {
	Register(user models.Register) (string, error)
}
type authService struct {
	UserRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		UserRepo: userRepo,
	}
}

func (service *authService) Register(register models.Register) (string, error) {
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

	res, error := service.UserRepo.Register(cust)
	return res, error
}
