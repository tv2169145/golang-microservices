package github_provider

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	authHeader := GetAuthorizationHeader("jimmy123")
	fmt.Println(authHeader)
	assert.EqualValues(t, "token jimmy123", authHeader)
}

func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("here")
}
