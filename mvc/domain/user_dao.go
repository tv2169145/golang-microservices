package domain

import (
	"fmt"
	"github.com/tv2169145/golang-microservices/mvc/untils"
	"net/http"
)

var (
	users = map[int64]*User{
	123: {Id:123, FirstName:"jimmy", LastName:"Lin", Email:"jimmy@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *untils.ApplicationError)
}

type userDao struct {}

func (u *userDao) GetUser(userId int64) (*User, *untils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &untils.ApplicationError{
		Message: fmt.Sprintf("user %v is not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not found",
	}
}
