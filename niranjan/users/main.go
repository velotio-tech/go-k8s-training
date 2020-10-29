package main

import (
	"log"
	"net/http"
	"os"

	"example.com/users/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "users-api ", log.LstdFlags)

	// create the handlers
	uh := handlers.NewUsers(l)

	router := mux.NewRouter()

	// users table handlers
	router.HandleFunc("/users", uh.PostUser).Methods("POST")
	router.HandleFunc("/users", uh.GetUsers).Methods("GET")
	router.HandleFunc("/users/{user}", uh.GetUser).Methods("GET")
	router.HandleFunc("/users/{user}", uh.PutUser).Methods("PUT")
	router.HandleFunc("/users/{user}", uh.DeleteUser).Methods("DELETE")

	// orders handlers
	router.HandleFunc("/users/{user}/orders", uh.PostUserOrder).Methods("POST")
	router.HandleFunc("/users/{user}/orders", uh.GetUserOrders).Methods("GET")
	router.HandleFunc("/users/{user}/orders/{order}", uh.GetUserOrder).Methods("GET")
	router.HandleFunc("/users/{user}/orders/{order}", uh.PutUserOrder).Methods("PUT")
	router.HandleFunc("/users/{user}/orders/", uh.DeleteUserOrders).Methods("DELETE")
	router.HandleFunc("/users/{user}/orders/{order}", uh.DeleteUserOrder).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
