package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func listOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]
	fmt.Printf("GET /users/%v/orders called...\n", userName)

	if orders, ok := allUsersOrders[userName]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(orders)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User %v Not Found\n", userName)
		fmt.Printf("User %v Not Found\n", userName)
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]
	fmt.Printf("POST /users/%v/orders called...", userName)

	orderIdCounter.increment()
	newOrderId := orderIdCounter

	// TODO validate user with users service.

	var o orderInfo
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the orders map for given user
	if orders, ok := allUsersOrders[userName]; ok {
		id := int(newOrderId)
		newOrder := order{
			Id:        id,
			OrderInfo: o,
		}
		orders[id] = newOrder
		// Write added objects json to response
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newOrder)
		fmt.Println(newOrder)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User %v Not Found\n", userName)
		fmt.Printf("User %v Not Found\n", userName)
	}
}

func getOrderInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]
	orderNumber, err := strconv.Atoi(vars["ordernumber"])
	if err != nil {
		fmt.Fprintf(w, "Invalid order number format %v\n", vars["ordernumber"])
		fmt.Printf("Invalid order number format %v\n", vars["ordernumber"])
		return
	}

	fmt.Printf("GET /users/%v/orders/%v called...\n", userName, orderNumber)

	// Find the orders map for given user
	if orders, ok := allUsersOrders[userName]; ok {
		// Find specific order
		if o, ok := orders[orderNumber]; ok {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(o)
			fmt.Println(o)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Order %v Not Found\n", orderNumber)
			fmt.Printf("Order %v Not Found\n", orderNumber)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User %v Not Found\n", userName)
		fmt.Printf("User %v Not Found\n", userName)
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]
	orderNumber, err := strconv.Atoi(vars["ordernumber"])
	if err != nil {
		fmt.Fprintf(w, "Invalid order number format %v\n", vars["ordernumber"])
		fmt.Printf("Invalid order number format %v\n", vars["ordernumber"])
		return
	}

	fmt.Printf("DELETE /users/%v/orders/%v called...\n", userName, orderNumber)

	// Find the orders map for given user
	if orders, ok := allUsersOrders[userName]; ok {
		// Find specific order
		if _, ok := orders[orderNumber]; ok {
			delete(orders, orderNumber)
			w.WriteHeader(http.StatusOK)
			fmt.Printf("DELETED order number %v for user %v.\n", orderNumber, userName)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Order %v Not Found\n", orderNumber)
			fmt.Printf("Order %v Not Found\n", orderNumber)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User %v Not Found\n", userName)
		fmt.Printf("User %v Not Found\n", userName)
	}
}

func main() {
	// Initially create space for 50 orders.
	allUsersOrders = make(usersOrdersMap, 50)

	// TBD : test data, remove later
	allUsersOrders["user1"] = make(ordersMap, 50)

	// Setup the http router.
	r := mux.NewRouter()
	r.HandleFunc("/users/{username}/orders", createOrder).Methods("POST")
	r.HandleFunc("/users/{username}/orders", listOrders).Methods("GET")
	r.HandleFunc("/users/{username}/orders/{ordernumber}", getOrderInfo).Methods("GET")
	r.HandleFunc("/users/{username}/orders/{ordernumber}", deleteOrder).Methods("DELETE")

	// Setup the server.
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
