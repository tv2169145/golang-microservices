package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/golang-microservices/src/api/log/option_a"
	"github.com/tv2169145/golang-microservices/src/api/log/option_b"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	option_a.Info("start", "status:pending", "step:1")
	option_b.Info("start",
		option_b.Field("client_id", 1),
		option_b.Field("status", "pending"))
	mapUrls()
	option_a.Info("go", "status:go", "step:2")
	option_b.Info("go",
		option_b.Field("client_id", 1),
		option_b.Field("status", "go"))
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
