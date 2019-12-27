package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/golang-microservices/mvc/services"
	"github.com/tv2169145/golang-microservices/mvc/untils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &untils.ApplicationError{
			Message: "user_id must a number",
			StatusCode: http.StatusBadRequest,
			Code: "baa_request",
		}
		untils.RespondError(c, apiErr)
		//c.JSON(apiErr.StatusCode, apiErr)

		//jsonValue, _ := json.Marshal(apiErr)
		//resp.WriteHeader(apiErr.StatusCode)
		//resp.Write(jsonValue)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		//c.JSON(apiErr.StatusCode, apiErr)
		untils.RespondError(c, apiErr)
		return
	}
	untils.Respond(c, http.StatusOK, user)
	//c.JSON(200, user)
	//jsonValue, _ := json.Marshal(user)
	//resp.Write(jsonValue)
}
