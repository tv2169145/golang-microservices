package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-microservices/src/api/clients/restclient"
	"github.com/tv2169145/golang-microservices/src/api/domain/repositories"
	"github.com/tv2169145/golang-microservices/src/api/services"
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
	"github.com/tv2169145/golang-microservices/src/api/utils/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	funcCreateRepo func (request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func (request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
)

type repoServiceMock struct {

}

func (s *repoServiceMock) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo(request)
}
func (s *repoServiceMock) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	return funcCreateRepos(request)
}

func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func (request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return &repositories.CreateRepoResponse{
			Name: "mock service",
			Owner: "golang",
			Id: 333,
		}, nil
	}

	response := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"/repositories",
		strings.NewReader(`{"name": "controller testing"}`),
	)
	c := test_utils.GetMockedContext(request, response)
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123,"name": "golang-rest-api","full_name": "tv2169145/golang-rest-api"}`)),
		},
	})
	CreateRepo(c)
	assert.EqualValues(t, http.StatusCreated, c.Writer.Status())
	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 333, result.Id)
}

func TestCreateRepoErrorFromGithubMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func (request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	response := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"/repositories",
		strings.NewReader(`{"name": "controller testing"}`),
	)
	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, c.Writer.Status())
	fmt.Println(response.Body.String())
	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid repository name", apiErr.Message())
}
