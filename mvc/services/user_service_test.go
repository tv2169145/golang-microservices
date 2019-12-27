package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-microservices/mvc/domain"
	"github.com/tv2169145/golang-microservices/mvc/untils"
	"log"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *untils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *untils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNoFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *untils.ApplicationError) {
		return nil, &untils.ApplicationError{
			Message: "user 111 is not found",
			StatusCode: http.StatusNotFound,
			Code: "not found",
		}
	}
	user, err := UserService.GetUser(111)
	assert.Nil(t, user, "user is not nil")
	assert.NotNil(t, err, "err is nil")
	assert.EqualValues(t, "user 111 is not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *untils.ApplicationError) {
		log.Println("here2")
		return &domain.User{
			Id: 123,
			FirstName:"jimmy",
			LastName:"Lin",
			Email:"jimmy@gmail.com",
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err, "err is not nil")
	assert.NotNil(t, user, "user is nil")
	assert.EqualValues(t, 123, user.Id)
}
