package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	User_id  string `json:"user_id"`
	Username string `json:"username"`
}

func EnableRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/addUser", AddNewUser).Methods("POST")
	r.HandleFunc("/getallorders", GetAllOrders).Methods("POST")
	r.HandleFunc("/deleteorder", DeleteOrders).Methods("POST")
	r.HandleFunc("/updateorder", UpdateOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8090", r))
}
