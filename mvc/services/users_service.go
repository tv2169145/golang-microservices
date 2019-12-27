package services

import (
	"github.com/tv2169145/golang-microservices/mvc/domain"
	"github.com/tv2169145/golang-microservices/mvc/untils"
	"log"
)

type userService struct {}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *untils.ApplicationError) {
	log.Println("here")
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
