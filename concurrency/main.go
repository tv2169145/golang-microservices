package main

import (
	"bufio"
	"fmt"
	"github.com/tv2169145/golang-microservices/src/api/domain/repositories"
	"github.com/tv2169145/golang-microservices/src/api/services"
	"github.com/tv2169145/golang-microservices/src/api/utils/errors"
	"os"
	"sync"
)

var (
	success map[string]string = make(map[string]string)
	failed map[string]errors.ApiError = make(map[string]errors.ApiError)
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result *repositories.CreateRepoResponse
	Error errors.ApiError
}

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)
	file, err := os.Open("/Users/jimmy/Desktop/requests.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}
	return result
}

func main() {
	var wg sync.WaitGroup
	requests := getRequests()
	input := make(chan createRepoResult)
	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))

	buffer := make(chan bool, 10)

	go handleResults(input, &wg)
	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(buffer, request, input)
	}
	wg.Wait()
	close(input)
	fmt.Println(success)
	fmt.Println(failed)
	
}

func handleResults(input chan createRepoResult, wg *sync.WaitGroup) {
	for result := range input {
		wg.Done()
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
			continue
		} else {
			success[result.Request.Name] = result.Result.Name
		}
	}
}

func createRepo(buffer chan bool, request repositories.CreateRepoRequest, output chan createRepoResult) {
	response, err := services.RepositoryService.CreateRepo(request)
	output <- createRepoResult{
		Request: request,
		Result: response,
		Error: err,
	}
	<- buffer
}
