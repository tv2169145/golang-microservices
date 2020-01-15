package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/golang-microservices/src/api/log"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Info("start", "status:pending", "step:1")
	mapUrls()
	log.Info("go", "status:go", "step:2")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
