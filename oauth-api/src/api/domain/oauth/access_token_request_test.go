package oauth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessTokenRequest_ValidateSuccess(t *testing.T) {
	request := &AccessTokenRequest{
		Username: "jimmy",
		Password: "123123",
	}
	err := request.Validate()
	assert.Nil(t, err)
}

func TestAccessTokenRequest_ValidateUsernameIsEmpty(t *testing.T) {
	request := &AccessTokenRequest{
		Password: "123123",
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid username", err.Message())
}

func TestAccessTokenRequest_ValidatePasswordIsEmpty(t *testing.T) {
	request := &AccessTokenRequest{
		Username: "jimmy",
		Password: "  ",
	}
	err := request.Validate()
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid password", err.Message())
}