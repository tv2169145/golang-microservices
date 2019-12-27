package app

import "github.com/tv2169145/golang-microservices/mvc/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
	//http.HandleFunc("/users", controllers.GetUser)
}