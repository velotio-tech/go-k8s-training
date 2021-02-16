package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/order/api/models"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/order/api/utils"
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {

	order := &models.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := order.Create()
	utils.Respond(w, resp)
}

var GetOrders = func(w http.ResponseWriter, r *http.Request) {

	var orders []models.Order
	models.GetDB().Find(&orders)
	response := utils.Message(true, "List of all orders")
	response["order"] = orders
	utils.Respond(w, response)
}

var GetOrder = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	response := utils.Message(true, "Order Details")
	response["order"] = models.GetOrder(vars["id"])
	utils.Respond(w, response)
}

var UpdateOrder = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	order := &models.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := order.Update(id)
	utils.Respond(w, resp)

}
var DeleteOrder = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	utils.Respond(w, models.DeleteOrder(id))
}
