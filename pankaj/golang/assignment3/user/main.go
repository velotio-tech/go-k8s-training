package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/user/api/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/user/{id}/orders", controllers.GetUserOrders).Methods("GET")
	router.HandleFunc("/api/user/{id}/orders/{orderID}", controllers.GetUserOrder).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
