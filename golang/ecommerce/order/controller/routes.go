package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/users/{user_id}/orders", GetOrders).Methods("GET")
	router.HandleFunc("/users/{user_id}/orders", CreateOrder).Methods("POST")
	router.HandleFunc("/users/{user_id}/orders/{id}", UpdateOrder).Methods("PATCH")
	router.HandleFunc("/users/{user_id}/orders/{id}", DeleteOrder).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/orders", DeleteAllOrdersForUser).Methods("DELETE")

	http.ListenAndServe(":8002", router)
}
