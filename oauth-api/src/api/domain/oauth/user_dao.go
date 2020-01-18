package oauth

import (
	"fmt"
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username = ? and password = ?;"
)

var (
	users = map[string]*User {
		"jimmy": &User{123, "jimmy"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("username: %s not found", username))
	}
	return user, nil
}
