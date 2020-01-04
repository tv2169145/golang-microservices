package polo

import (
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-microservices/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestMarco(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/marco", strings.NewReader(``))
	c := test_utils.GetMockedContext(request, response)
	Marco(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
