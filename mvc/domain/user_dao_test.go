package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "not expect a user with id 0")
	assert.NotNil(t, err, "we were not expecting error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "status code is err")
	assert.EqualValues(t, "user 0 is not found", err.Message)
	assert.EqualValues(t, "not found", err.Code)
	//if user != nil {
	//	t.Error("not expect a user with id 0")
	//}
	//if err == nil {
	//	t.Error("we were not expecting error when user id is 0")
	//}
	//if err.StatusCode != http.StatusNotFound {
	//	t.Error("status code is err")
	//}
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.NotNil(t, user, "user is nil")
	assert.Nil(t, err, "error is not nil")
	assert.EqualValues(t, 123, user.Id, "user 123 id not equal 123")
	assert.EqualValues(t, "jimmy", user.FirstName, "user 123 FirstName not equal jimmy")
	assert.EqualValues(t, "Lin", user.LastName, "user 123 LastName not equal Lin")
	assert.EqualValues(t, "jimmy@gmail.com", user.Email, "user 123 Email not equal jimmy@gmail.com")
}
