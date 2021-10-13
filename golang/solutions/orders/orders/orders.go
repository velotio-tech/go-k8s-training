package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

	if _, found := allUsersOrders[userName]; !found {
		// Users service will validate user name, orders assumes its validated.
		allUsersOrders[userName] = make(ordersMap, 50)
	}

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
