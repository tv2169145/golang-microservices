package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-microservices/src/api/clients/restclient"
	"github.com/tv2169145/golang-microservices/src/api/domain/repositories"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{
		Name: " ",
		Description: "for testing",
	}
	response, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, response, "response is not nil")
	assert.NotNil(t, err, "err is nil")
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "golang test",
		Description: "for testing",
	}
	result, err := RepositoryService.CreateRepo(request)
	fmt.Println(err.Status(), err.Message())
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123,"name": "golang-rest-api","full_name": "tv2169145/golang-rest-api"}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "golang test",
		Description: "for testing",
	}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
}