package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Print("Listening on http://localhost:8001/")
	http.ListenAndServe(":8001", router)
}
