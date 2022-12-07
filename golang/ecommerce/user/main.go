package main

import (
	"net/http"
	"user/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":8001", router)
}
