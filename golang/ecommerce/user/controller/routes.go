package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	http.ListenAndServe(":8001", router)
}
