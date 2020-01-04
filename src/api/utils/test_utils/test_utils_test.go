package test_utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMockedContext(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/something", nil)
	request.Header = http.Header{"X-Mock": {"true"}}
	c := GetMockedContext(request, response)
	assert.NotNil(t, c)
	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "8080", c.Request.URL.Port())
	assert.EqualValues(t, "/something", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))
}
