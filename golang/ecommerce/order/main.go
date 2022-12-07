package main

import (
	"net/http"
	"order/controller"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/users/{user_id}/orders", controller.GetOrders).Methods("GET")
	router.HandleFunc("/users/{user_id}/orders", controller.CreateOrder).Methods("POST")
	router.HandleFunc("/users/{user_id}/orders/{id}", controller.UpdateOrder).Methods("PATCH")
	router.HandleFunc("/users/{user_id}/orders/{id}", controller.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/orders", controller.DeleteAllOrdersForUser).Methods("DELETE")

	http.ListenAndServe(":8001", router)
}
