package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-microservices/src/api/clients/restclient"
	"github.com/tv2169145/golang-microservices/src/api/domain/repositories"
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
	"github.com/tv2169145/golang-microservices/src/api/utils/test_utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(
		http.MethodPost,
		"/repositories",
		strings.NewReader(``),
	)
	c := test_utils.GetMockedContext(request, response)
	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	fmt.Println(response.Body.String())

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(
		http.MethodPost,
		"/repositories",
		strings.NewReader(`{"name": "controller testing"}`),
		)
	c := test_utils.GetMockedContext(request, response)
	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
		},
	})

	CreateRepo(c)
	assert.EqualValues(t, http.StatusUnauthorized, response.Code)

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requires authentication", apiErr.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(
		http.MethodPost,
		"/repositories",
		strings.NewReader(`{"name": "controller testing"}`),
	)
	c := test_utils.GetMockedContext(request, response)
	c.Request = request
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
	assert.EqualValues(t, 123, result.Id)
}