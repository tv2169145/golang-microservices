package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-microservices/oauth-api/src/api/domain/oauth"
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
	"testing"
	"time"
)

var (
	funcCreateAccessToken func (request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	funcGetAccessToken func (accessToken string) (*oauth.AccessToken, errors.ApiError)
)

type oauthServiceMock struct {}

func (s *oauthServiceMock) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	return funcCreateAccessToken(request)
}

func (s *oauthServiceMock) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	return funcGetAccessToken(accessToken)
}

func TestOauthService_CreateAccessTokenInvalidRequest(t *testing.T) {
	//service := new(oauthServiceMock)
	request := oauth.AccessTokenRequest{
		Username: "",
		Password: "123",
	}
	token, err := OauthService.CreateAccessToken(request)
	assert.Nil(t, token)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid username", err.Message())
}

func TestOauthService_CreateAccessTokenNoFoundUser(t *testing.T) {
	request := oauth.AccessTokenRequest{
		Username: "ttt",
		Password: "123",
	}
	token, err := OauthService.CreateAccessToken(request)
	assert.Nil(t, token)
	assert.NotNil(t, err)
	assert.EqualValues(t, "username: ttt not found", err.Message())
}

func TestOauthService_CreateAccessTokenSaveError(t *testing.T) {
	oauth.StartMock()
	request := oauth.AccessTokenRequest{
		Username: "jimmy",
		Password: "123",
	}
	token, err := OauthService.CreateAccessToken(request)
	assert.Nil(t, token)
	assert.NotNil(t, err)
	assert.EqualValues(t, "is testing", err.Message())
}

func TestOauthService_CreateAccessTokenSuccess(t *testing.T) {
	request := oauth.AccessTokenRequest{
		Username: "jimmy",
		Password: "123",
	}
	token, err := OauthService.CreateAccessToken(request)
	assert.Nil(t, err)
	assert.NotNil(t, token)
	assert.EqualValues(t, "USR_123", token.AccessToken)
}

func TestOauthService_GetAccessTokenNoFoundToken(t *testing.T) {
	accessToken := "URS_1234"
	token, err := OauthService.GetAccessToken(accessToken)
	assert.Nil(t, token)
	assert.NotNil(t, err)
	assert.EqualValues(t, "no access token found or token is expired", err.Message())
}

func TestOauthService_GetAccessTokenSuccess(t *testing.T) {
	zone, _ := time.LoadLocation("Asia/Taipei")
	token := oauth.AccessToken{
		UserId: 123,
		Expires: time.Now().In(zone).Add(1 * time.Hour).Unix(),
	}
	token.Save()
	fmt.Println(token.AccessToken)
	getToken, err := OauthService.GetAccessToken(token.AccessToken)
	assert.Nil(t, err)
	assert.NotNil(t, getToken)
	assert.EqualValues(t, "USR_123", token.AccessToken)
}
