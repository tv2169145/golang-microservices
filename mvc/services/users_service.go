package services

import (
	"github.com/tv2169145/golang-microservices/mvc/domain"
	"github.com/tv2169145/golang-microservices/mvc/untils"
)

func GetUser(userId int64) (*domain.User, *untils.ApplicationError) {
	return domain.GetUser(userId)
}
