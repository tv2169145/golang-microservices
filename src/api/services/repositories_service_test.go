package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-microservices/src/api/clients/restclient"
	"github.com/tv2169145/golang-microservices/src/api/domain/repositories"
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
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

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	output := make(chan repositories.CreateRepositoriesResult)
	service := &reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.EqualValues(t, "invalid repository name", result.Error.Message())
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
}

func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {
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
		Name: "for testing",
		Description: "",
	}
	output := make(chan repositories.CreateRepositoriesResult)
	service := &reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.EqualValues(t, "Requires authentication", result.Error.Message())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
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
		Name: "for testing",
		Description: "",
	}
	output := make(chan repositories.CreateRepositoriesResult)
	service := &reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.EqualValues(t, 123, result.Response.Id)
	assert.EqualValues(t, "golang-rest-api", result.Response.Name)
}

func TestHandleRepoResults(t *testing.T) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	var wg sync.WaitGroup
	service := &reposService{}
	go service.handleRepoResults(input, output, &wg)

	wg.Add(1)
	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("invalid repository name"),
		}
	}()
	wg.Wait()
	close(input)

	result := <-output
	assert.NotNil(t, result)
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
}

func TestCreateReposInvalidRequest(t *testing.T) {
	request := []repositories.CreateRepoRequest{
		{},
		{
			Name: "  ",
		},
	}
	response, err := RepositoryService.CreateRepos(request)
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, 2, len(response.Results))
	assert.EqualValues(t, http.StatusBadRequest, response.StatusCode)
}

func TestCreateReposSingleSuccess(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123,"name": "golang-rest-api","full_name": "tv2169145/golang-rest-api"}`)),
		},
	})
	request := []repositories.CreateRepoRequest{
		{},
		{
			Name: "for test",
		},
	}
	response, err := RepositoryService.CreateRepos(request)

	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, 2, len(response.Results))
	assert.EqualValues(t, http.StatusPartialContent, response.StatusCode)

	for _, r := range response.Results {
		if r.Response != nil {
			assert.EqualValues(t, 123, r.Response.Id)
			assert.EqualValues(t, "golang-rest-api", r.Response.Name)
			continue
		}
		if r.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, r.Error.Status())
			continue
		}
	}
}

func TestCreateReposAllSuccess(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123,"name": "golang-rest-api","full_name": "tv2169145/golang-rest-api"}`)),
		},
	})
	request := []repositories.CreateRepoRequest{
		{
			Name: "testing",
		},
		{
			Name: "testing",
		},
	}
	response, err := RepositoryService.CreateRepos(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusCreated, response.StatusCode)
	fmt.Println(response)
	//assert.EqualValues(t, http.StatusCreated, response.StatusCode)
}