package oauth

import (
	"fmt"
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() (errors.ApiError) {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserId)
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.ApiError) {
	token := tokens[accessToken]
	if token == nil || token.IsExpired() {
		return nil, errors.NewNotFoundError("no access token found or token is expired")
	}
	return token, nil
}