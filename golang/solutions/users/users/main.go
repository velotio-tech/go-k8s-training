package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type user struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Global var, not a good practise to share across requests as each can run in parallel.
type usersMap map[string]user

var allUsers usersMap

var (
	usersServiceHostDefault  = "0.0.0.0"
	usersServicePortDefault  = "8080"
	ordersServiceHostDefault = "orderswebappsvc"
	ordersServicePortDefault = "8081"
	OrdersEndpoint           = "http://" + ordersServiceHostDefault + ":" + ordersServicePortDefault
)

func main() {
	// Initially create space for 50 users.
	allUsers = make(usersMap, 50)

	// Setup the http router.
	r := mux.NewRouter()

	// Users Resource served by current service.
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", listUsers).Methods("GET")
	r.HandleFunc("/users/{name}", getUserInfo).Methods("GET")
	r.HandleFunc("/users/{name}", checkUser).Methods("HEAD")
	r.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")

	// Order Resource - Served by Orders service.
	r.HandleFunc("/users/{username}/orders", createOrder).Methods("POST")
	r.HandleFunc("/users/{username}/orders", listOrders).Methods("GET")
	r.HandleFunc("/users/{username}/orders/{ordernumber}", getOrderInfo).Methods("GET")
	r.HandleFunc("/users/{username}/orders/{ordernumber}", deleteOrder).Methods("DELETE")

	// Setup the server.
	usersSvcHostName := os.Getenv("USERS_SVC_HOST_NAME")
	if len(usersSvcHostName) == 0 {
		usersSvcHostName = usersServiceHostDefault
	}
	usersSvcPort := os.Getenv("USERS_SVC_PORT")
	if len(usersSvcPort) == 0 {
		usersSvcPort = usersServicePortDefault
	}

	ordersSvcHostName := os.Getenv("ORDERS_SVC_HOST_NAME")
	if len(ordersSvcHostName) == 0 {
		ordersSvcHostName = ordersServiceHostDefault
	}
	ordersSvcPort := os.Getenv("ORDERS_SVC_PORT")
	if len(ordersSvcPort) == 0 {
		ordersSvcPort = ordersServicePortDefault
	}
	OrdersEndpoint = "http://" + ordersSvcHostName + ":" + ordersSvcPort

	srv := &http.Server{
		Handler: r,
		Addr:    usersSvcHostName + ":" + usersSvcPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
