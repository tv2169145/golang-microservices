package oauth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserByUsernameAndPassword(t *testing.T) {
	user, err := GetUserByUsernameAndPassword("user", "123")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, "username: user not found", err.Message())

	user, err = GetUserByUsernameAndPassword("jimmy", "123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, "jimmy", user.Username)
	assert.EqualValues(t, 123, user.Id)
}
