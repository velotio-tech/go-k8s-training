package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/order/api/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/api/order/{id}", controllers.GetOrder).Methods("GET")
	router.HandleFunc("/api/order/{id}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/api/order/{id}", controllers.DeleteOrder).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
