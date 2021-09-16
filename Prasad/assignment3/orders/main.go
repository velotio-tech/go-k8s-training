package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/thisisprasad/go-k8s-training/Prasad/assignment3/orders/controllers"
)

func main() {
	// NewOrderApp().StartApplication()

	router := mux.NewRouter()

	router.HandleFunc("/order", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/user/{userId}", controllers.GetUserOrders).Methods("GET")
	router.HandleFunc("/order/{id}", controllers.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/orders/user/{userId}", controllers.DeleteAllOrdersOfUser).Methods("DELETE")

	port := os.Getenv("ORDER_SERVER_PORT")
	if port == "" {
		port = "8091"
	}

	log.Println("Starting order server port at:", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalln(err)
	}
}
