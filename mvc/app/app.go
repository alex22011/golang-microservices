package app

import (
	"net/http"

	"github.com/alex22011/golang-microservices/mvc/controllers"
)
//export this
func StartAPP() {
	http.HandleFunc("/users", controllers.GetUser)

	err :=http.ListenAndServe(":8080", nil)
	if err!=nil {
		panic(err)
	}
}
