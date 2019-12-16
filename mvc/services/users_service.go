package services

import "github.com/tv2169145/golang-microservices/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
