package services

import (
	"github.com/tv2169145/golang-microservices/oauth-api/src/api/domain/oauth"
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
	"time"
)

type OauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

type oauthService struct{}

var (
	OauthService OauthServiceInterface
)

func init() {
	OauthService = new(oauthService)
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	zone, _ := time.LoadLocation("Asia/Taipei")
	token := &oauth.AccessToken{
		UserId: user.Id,
		Expires: time.Now().In(zone).Add(24 * time.Hour).Unix(),
	}
	if err := token.Save(); err != nil {
		return nil, err
	}
	return token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}
	return token, err
}
