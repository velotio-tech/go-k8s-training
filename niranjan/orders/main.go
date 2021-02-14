package main

import (
	"log"
	"net/http"
	"os"

	"example.com/orders/handlers"
	"example.com/orders/models"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "orders-api ", log.LstdFlags)

	// Creating new table
	models.TableCreate(l)

	// create order handler
	oh := handlers.NewOrders(l)

	router := mux.NewRouter()

	// orders handlers
	router.HandleFunc("/users/{user}/orders", oh.PostUserOrder).Methods("POST")
	router.HandleFunc("/users/{user}/orders", oh.GetUserOrders).Methods("GET")
	router.HandleFunc("/users/{user}/orders/{order}", oh.GetUserOrder).Methods("GET")
	router.HandleFunc("/users/{user}/orders/{order}", oh.PutUserOrder).Methods("PUT")
	router.HandleFunc("/users/{user}/orders", oh.DeleteUserOrders).Methods("DELETE")
	router.HandleFunc("/users/{user}/orders/{order}", oh.DeleteUserOrder).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":6000", router))
}
