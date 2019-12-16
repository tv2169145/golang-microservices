package app

import (
	"fmt"
	"github.com/tv2169145/golang-microservices/mvc/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	fmt.Println("here")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}