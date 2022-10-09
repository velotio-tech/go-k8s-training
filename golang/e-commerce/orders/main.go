package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users/{user_id}/orders", GetOrders).Methods("GET")
	router.HandleFunc("/users/{user_id}/orders", CreateOrder).Methods("POST")
	router.HandleFunc("/users/{user_id}/orders/{id}", UpdateOrder).Methods("PATCH")
	router.HandleFunc("/users/{user_id}/orders/{id}", DeleteOrder).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/orders", DeleteAllOrdersForUser).Methods("DELETE")

	log.Print("Listening on http://localhost:8002/")
	http.ListenAndServe(":8002", router)
}
