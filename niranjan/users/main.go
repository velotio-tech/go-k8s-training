package main

import (
	"log"
	"net/http"
	"os"

	"example.com/users/handlers"
	"example.com/users/models"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "users-api ", log.LstdFlags)

	// Creating new table
	models.TableCreate(l)

	// create user handler
	uh := handlers.NewUsers(l)

	// create order handler
	oh := handlers.NewOrders(l)

	router := mux.NewRouter()

	// users table handlers
	router.HandleFunc("/users", uh.PostUser).Methods("POST")
	router.HandleFunc("/users", uh.GetUsers).Methods("GET")
	router.HandleFunc("/users/{user}", uh.GetUser).Methods("GET")
	router.HandleFunc("/users/{user}", uh.PutUser).Methods("PUT")
	router.HandleFunc("/users/{user}", uh.DeleteUser).Methods("DELETE")

	// orders handlers
	router.HandleFunc("/users/{user}/orders", oh.PostUserOrder).Methods("POST")
	router.HandleFunc("/users/{user}/orders", oh.GetUserOrders).Methods("GET")
	router.HandleFunc("/users/{user}/orders/{order}", oh.GetUserOrder).Methods("GET")
	router.HandleFunc("/users/{user}/orders/{order}", oh.PutUserOrder).Methods("PUT")
	router.HandleFunc("/users/{user}/orders", oh.DeleteUserOrders).Methods("DELETE")
	router.HandleFunc("/users/{user}/orders/{order}", oh.DeleteUserOrder).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
