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
	router.HandleFunc("/users", uh.GetUsers)
	router.HandleFunc("/users/{user}/orders", uh.GetUserOrders)
	router.HandleFunc("/users/{user}/orders/{order}", uh.GetUserOrder)

	http.ListenAndServe(":5000", router)
}
