package services

import (
	"github.com/tv2169145/golang-microservices/src/api/config"
	"github.com/tv2169145/golang-microservices/src/api/domain/github"
	"github.com/tv2169145/golang-microservices/src/api/domain/repositories"
	"github.com/tv2169145/golang-microservices/src/api/providers/github_provider"
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
	"net/http"
	"sync"
)

type reposService struct {}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = new(reposService)
}

func(s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	request := github.CreateRepoRequest{
		Name: input.Name,
		Description: input.Description,
		Private: false,
	}
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	result := repositories.CreateRepoResponse{
		Id: response.Id,
		Name: response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func(s *reposService) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)  {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)
	var wg sync.WaitGroup

	go s.handleRepoResults(input, output, &wg)
	for _, current := range request {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}
	wg.Wait()
	close(input)
	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}
	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(request) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func(s *reposService) handleRepoResults(input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse, wg *sync.WaitGroup) {
	 var results repositories.CreateReposResponse
	 for result := range input {
	 	//repoResult := repositories.CreateRepositoriesResult{
	 	//	Response: result.Response,
	 	//	Error: result.Error,
		//}
		 results.Results = append(results.Results, result)
		 //results.Results = append(results.Results, result)
		 wg.Done()
	 }
	 output <- results
}

func(s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult)  {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	result, err := s.CreateRepo(input)

	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}
}
