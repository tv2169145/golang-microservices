package controllers

import (
	"encoding/json"
	"github.com/tv2169145/golang-microservices/mvc/services"
	"github.com/tv2169145/golang-microservices/mvc/untils"

	//"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &untils.ApplicationError{
			Message: "user_id must a number",
			StatusCode: http.StatusBadRequest,
			Code: "baa_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		//http.Error(resp, "fail request", http.StatusBadRequest)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
