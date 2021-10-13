package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type orderInfo struct {
	ProductName string `json:"product_name"`
}

type order struct {
	Id        int       `json:"id"`
	OrderInfo orderInfo `json:"order_details"`
}

// orders with order ids and details
type ordersMap map[int]order

// all users with their orders
type usersOrdersMap map[string]ordersMap

type orderIdType int

func (orderId *orderIdType) increment() {
	(*orderId)++
}

// Global var, not a good practise to share across requests as each can run in parallel.
var allUsersOrders usersOrdersMap
var orderIdCounter = orderIdType(0)

var (
	ordersServiceHostDefault = "0.0.0.0"
	ordersServicePortDefault = "8081"
)

func main() {
	// Initially create space for 50 orders.
	allUsersOrders = make(usersOrdersMap, 50)

	// Setup the http router.
	r := mux.NewRouter()
	r.HandleFunc("/users/{username}/orders", createOrder).Methods("POST")
	r.HandleFunc("/users/{username}/orders", listOrders).Methods("GET")
	r.HandleFunc("/users/{username}/orders/{ordernumber}", getOrderInfo).Methods("GET")
	r.HandleFunc("/users/{username}/orders/{ordernumber}", deleteOrder).Methods("DELETE")

	// Setup the server.
	ordersSvcHostName := os.Getenv("ORDERS_SVC_HOST_NAME")
	if len(ordersSvcHostName) == 0 {
		ordersSvcHostName = ordersServiceHostDefault
	}
	ordersSvcPort := os.Getenv("ORDERS_SVC_PORT")
	if len(ordersSvcPort) == 0 {
		ordersSvcPort = ordersServicePortDefault
	}
	srv := &http.Server{
		Handler: r,
		Addr:    ordersSvcHostName + ":" + ordersSvcPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
