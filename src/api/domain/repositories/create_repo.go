package repositories

import (
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
	"strings"
)

type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func(r *CreateRepoRequest) Validate() errors.ApiError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}

type CreateRepoResponse struct {
	Id int	`json:"id"`
	Owner string `json:"owner"`
	Name string `json:"name"`
}

//----------------併發用
type CreateReposResponse struct {
	 StatusCode int `json:"status"`
	 Results []CreateRepositoriesResult `json:"results"`
}

type  CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error errors.ApiError `json:"error"`
}
