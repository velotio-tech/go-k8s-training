package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Proxy all requests to orders service.
func proxyRequestToOrdersSvc(url string, method string,
	w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}
	req, err :=
		http.NewRequest(method, url, r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	// Relay the status code as is.
	w.WriteHeader(res.StatusCode)
	if err != nil {
		log.Println("Error processing request: ", r.URL)
		return
	}
	defer res.Body.Close()
	// Relay the response body.
	_, err = io.Copy(w, res.Body)

	if err != nil {
		log.Println("Error processing request: ", r.URL)
		return
	}
}

func respondUserNotFound(userName string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "User %v Not Found\n", userName)
	fmt.Printf("User %v Not Found\n", userName)
}

func listOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]
	fmt.Printf("GET /users/%v/orders called...\n", userName)

	// Validate username locally and then proxy to order.
	if _, found := allUsers[userName]; found {
		url := OrdersEndpoint + "/users/" + userName + "/orders"
		proxyRequestToOrdersSvc(url, "GET", w, r)
	} else {
		respondUserNotFound(userName, w)
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]
	fmt.Printf("POST /users/%v/orders called...", userName)

	// Validate username locally and then proxy to order.
	if _, found := allUsers[userName]; found {
		url := OrdersEndpoint + "/users/" + userName + "/orders"
		proxyRequestToOrdersSvc(url, "POST", w, r)
	} else {
		respondUserNotFound(userName, w)
	}
}

func getOrderInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]
	orderNumber := vars["ordernumber"]
	fmt.Printf("GET /users/%v/orders/%v called...\n", userName, orderNumber)

	// Validate username locally and then proxy to order.
	if _, found := allUsers[userName]; found {
		url := OrdersEndpoint + "/users/" + userName + "/orders/" + orderNumber
		proxyRequestToOrdersSvc(url, "GET", w, r)
	} else {
		respondUserNotFound(userName, w)
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["username"]
	orderNumber := vars["ordernumber"]
	fmt.Printf("DELETE /users/%v/orders/%v called...\n", userName, orderNumber)

	// Validate username locally and then proxy to order.
	if _, found := allUsers[userName]; found {
		url := OrdersEndpoint + "/users/" + userName + "/orders/" + orderNumber
		proxyRequestToOrdersSvc(url, "DELETE", w, r)
	} else {
		respondUserNotFound(userName, w)
	}
}
