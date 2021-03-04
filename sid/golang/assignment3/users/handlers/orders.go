package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farkaskid/go-k8s-training/assignment3/users/helpers"
)

func OrderHandler(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getOrderHandler(resp, req)
	case http.MethodPost:
		createOrderHandler(resp, req)
	case http.MethodPut:
		updateOrderHandler(resp, req)
	case http.MethodDelete:
		deleteOrderHandler(resp, req)
	default:
		resp.WriteHeader(http.StatusBadRequest)
	}
}

func createOrderHandler(resp http.ResponseWriter, req *http.Request) {
	msg, err := helpers.CreateOrder(req)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	response := make(map[string]string)
	response["msg"] = msg

	responseJSON, err := json.Marshal(response)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(responseJSON)
}

func getOrderHandler(resp http.ResponseWriter, req *http.Request) {
	orders, err := helpers.GetOrders(req)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	type order struct {
		ID      int    `json:"id"`
		Details string `json:"details"`
	}

	response := make(map[string][]order)
	response["orders"] = make([]order, len(orders))
	index := 0

	for id, details := range orders {
		response["orders"][index] = order{ID: id, Details: details}
		index++
	}

	responseJSON, err := json.Marshal(response)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(responseJSON)
}

func updateOrderHandler(resp http.ResponseWriter, req *http.Request) {
	msg, err := helpers.CreateOrder(req)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	response := make(map[string]string)
	response["msg"] = msg

	responseJSON, err := json.Marshal(response)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(responseJSON)
}

func deleteOrderHandler(resp http.ResponseWriter, req *http.Request) {
	msg, err := helpers.CreateOrder(req)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	response := make(map[string]string)
	response["msg"] = msg

	responseJSON, err := json.Marshal(response)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(responseJSON)
}
