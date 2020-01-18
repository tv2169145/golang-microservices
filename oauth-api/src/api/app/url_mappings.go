package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/golang-microservices/oauth-api/src/api/controllers/oauth"
	"github.com/tv2169145/golang-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}

func index(r *gin.Context) {

}
