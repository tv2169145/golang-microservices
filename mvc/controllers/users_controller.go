package controllers

import (
	"encoding/json"
	"github.com/tv2169145/golang-microservices/mvc/services"
	//"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		//http.Error(resp, "fail request", http.StatusBadRequest)
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_id must a number"))
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
