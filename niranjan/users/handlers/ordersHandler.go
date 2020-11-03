package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// UserOrderRequest response struct to the user
type UserOrderRequest struct {
	BillAmount uint64 `json:"bill_amount"`
}

// Orders is a http.Handler
type Orders struct {
	l *log.Logger
}

// NewOrders creates order handler with the given logger which can call all order handlers
func NewOrders(l *log.Logger) *Orders {
	return &Orders{l}
}

// HOST address & PORT number used by Orders-ms - user req will be forwarded to this URL
const HOST = "http://orders-svc:6000/"

// PostUserOrder creates order for specific user - C
func (o *Orders) PostUserOrder(w http.ResponseWriter, r *http.Request) {

	o.l.Println("Handle POST Order")

	// Generating new URL
	vars := mux.Vars(r)
	userID := vars["user"]
	url := HOST + "users/" + userID + "/orders"

	// Receive request from user & decode it
	var userOrder UserOrderRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&userOrder)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// make new request from user decoded request & Encode it to send to Orders-ms
	userOrder = UserOrderRequest{BillAmount: userOrder.BillAmount}
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(userOrder)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// make POST request to Orders-ms
	res, _ := http.Post(url, "application/json; charset=utf-8", b)
	defer res.Body.Close()

	// Decode the response received from Orders-ms
	var body interface{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	finalRespone := fmt.Sprintf("%v", body)

	// JSONify the received response & send it to the user
	ToJSON(w, finalRespone, http.StatusOK)

}

// GetUserOrder gets specific order of specific user - R
func (o *Orders) GetUserOrder(w http.ResponseWriter, r *http.Request) {

	o.l.Println("Handle GET Order")

	vars := mux.Vars(r)
	userID := vars["user"]
	orderID := vars["order"]

	url := HOST + "users/" + userID + "/orders/" + orderID

	res, err := http.Get(url)

	defer res.Body.Close()

	// Decode the response received from Orders-ms
	var body interface{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// JSONify the received response & send it to the user
	ToJSON(w, body, http.StatusOK)
}

// GetUserOrders gets all orders of specific user - R
func (o *Orders) GetUserOrders(w http.ResponseWriter, r *http.Request) {

	o.l.Println("Handle GET Orders")

	vars := mux.Vars(r)
	userID := vars["user"]

	// Generating new URL
	url := HOST + "users/" + userID + "/orders"

	// sending GET request
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Decode the response received from Orders-ms
	var body interface{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// JSONify the received response & send it to the user
	ToJSON(w, body, http.StatusOK)

}

// PutUserOrder updates specific order of specific user - U
func (o *Orders) PutUserOrder(w http.ResponseWriter, r *http.Request) {

	o.l.Println("Handle PUT Order")

	// Generating new URL
	vars := mux.Vars(r)
	userID := vars["user"]
	orderID := vars["order"]
	url := HOST + "users/" + userID + "/orders/" + orderID

	// Receive request from user & decode it
	var userOrder UserOrderRequest
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&userOrder)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// make new request from user decoded request & Encode it to send to Orders-ms
	userOrder = UserOrderRequest{BillAmount: userOrder.BillAmount}
	b := new(bytes.Buffer)
	err = json.NewEncoder(b).Encode(userOrder)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// make PUT request to Orders-ms
	// initialize http client
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, url, b)
	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	defer res.Body.Close()

	// Decode the response received from Orders-ms
	var body interface{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	finalRespone := fmt.Sprintf("%v", body)

	// JSONify the received response & send it to the user
	ToJSON(w, finalRespone, http.StatusOK)

}

// DeleteUserOrder deletes specific order of specific user - D
func (o *Orders) DeleteUserOrder(w http.ResponseWriter, r *http.Request) {

	o.l.Println("Handle DELETE Order")

	// Generating new URL
	vars := mux.Vars(r)
	userID := vars["user"]
	orderID := vars["order"]
	url := HOST + "users/" + userID + "/orders/" + orderID

	// make DELETE request to Orders-ms
	// initialize http client
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	res, err := client.Do(req)
	defer res.Body.Close()

	// Decode the response received from Orders-ms
	var body interface{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	finalRespone := fmt.Sprintf("%v", body)

	// JSONify the received response & send it to the user
	ToJSON(w, finalRespone, http.StatusOK)

}

// DeleteUserOrders deletes all order of specific user - D
func (o *Orders) DeleteUserOrders(w http.ResponseWriter, r *http.Request) {

	o.l.Println("Handle DELETE Orders")

	// Generating new URL
	vars := mux.Vars(r)
	userID := vars["user"]
	url := HOST + "users/" + userID + "/orders"

	// make DELETE request to Orders-ms
	// initialize http client
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	res, err := client.Do(req)
	defer res.Body.Close()

	// Decode the response received from Orders-ms
	var body interface{}
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	finalRespone := fmt.Sprintf("%v", body)

	// JSONify the received response & send it to the user
	ToJSON(w, finalRespone, http.StatusOK)

}
