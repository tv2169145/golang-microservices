package domain

import (
	"errors"
	"fmt"
)

var users = map[int64]*User{
	123: {Id:123, FirstName:"jimmy", LastName:"Lin", Email:"jimmy@gmail.com"},

}

func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("user %v is not found", userId))
}
