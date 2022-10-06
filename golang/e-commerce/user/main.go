package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")

	log.Print("Listening on http://localhost:8001/")
	http.ListenAndServe(":8001", router)
}
