package app

import (
	"github.com/tv2169145/golang-microservices/src/api/controllers/polo"
	"github.com/tv2169145/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
