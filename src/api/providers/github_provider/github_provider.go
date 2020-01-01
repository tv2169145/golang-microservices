package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/tv2169145/golang-microservices/src/api/clients/restclient"
	"github.com/tv2169145/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo = "https://api.github.com/user/repos"
)

func GetAuthorizationHeader(token string) string {
	return fmt.Sprintf(headerAuthorizationFormat, token)
}

func CreateRepo(token string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	// Authorization: token abcc3bf6925c146cdc2ea93d772644ea129a5b50
	headers := http.Header{}
	headers.Set("Authorization", GetAuthorizationHeader(token))
	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("err when try to creating new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "invalid response body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		//if err := json.NewDecoder(response.Body).Decode(&errResponse); err != nil {
		//	return nil, &errResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
		//}

		if err := json.Unmarshal(jsonBytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message: "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode

		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		log.Println(fmt.Sprintf("err when trying to unmarshal create repo successful response %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "error unmarshal successful response",
		}
	}

	return &result, nil
}
