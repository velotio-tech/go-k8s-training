package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/addUser", AddUser).Methods("POST")
	r.HandleFunc("/users/{user_id}/orders/{order_id}/{updated_product_id}", UpdateOrder).Methods("POST")
	r.HandleFunc("/users/delete/{order_id}", DeleteOrder).Methods("GET")
	r.HandleFunc("/users/{user_id}/orders", GetAllUserOrders).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
