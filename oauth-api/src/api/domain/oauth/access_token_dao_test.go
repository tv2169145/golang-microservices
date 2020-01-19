package oauth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessToken_Save(t *testing.T) {
	zone, _ := time.LoadLocation("Asia/Taipei")
	newToken := &AccessToken{
		UserId: 123,
		Expires: time.Now().In(zone).Unix(),
	}
	err := newToken.Save()

	assert.Nil(t, err)
	assert.NotNil(t, newToken.AccessToken)
	assert.EqualValues(t, newToken, tokens[newToken.AccessToken])
}

func TestGetAccessTokenByTokenGetError(t *testing.T) {
	token, err := GetAccessTokenByToken("USR_123")
	assert.Nil(t, token)
	assert.NotNil(t, err)
	assert.EqualValues(t, "no access token found or token is expired", err.Message())
}

func TestGetAccessTokenByTokenSuccess(t *testing.T) {
	zone, _ := time.LoadLocation("Asia/Taipei")
	newToken := &AccessToken{
		UserId: 123,
		Expires: time.Now().In(zone).Add(1 * time.Hour).Unix(),
	}
	err := newToken.Save()
	fmt.Println(newToken)
	assert.Nil(t, err)
	token, err := GetAccessTokenByToken(newToken.AccessToken)
	fmt.Println(token)
	assert.Nil(t, err)
	assert.NotNil(t, token)
	assert.EqualValues(t, "USR_123", token.AccessToken)
}
