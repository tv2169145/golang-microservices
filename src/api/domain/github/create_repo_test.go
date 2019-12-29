package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		"golang api test",
		"test api for golang",
		"https://github.com",
		true,
		false,
		true,
		false,
	}
	jsonValue, err := json.Marshal(request)
	assert.Nil(t, err, "err not nil")
	assert.NotNil(t, jsonValue, "jsonValue is nil")
	assert.EqualValues(t, `{"name":"golang api test","description":"test api for golang","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`, string(jsonValue))
	fmt.Println(string(jsonValue))

	var target CreateRepoRequest
	err = json.Unmarshal(jsonValue, &target)
	assert.Nil(t, err, "err not nil")
	assert.EqualValues(t, request.Name, target.Name)
	fmt.Println(target)

}
