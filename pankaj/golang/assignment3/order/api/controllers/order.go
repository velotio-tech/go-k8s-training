package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/order/api/models"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/order/api/utils"
)

var CreateOrder = func(w http.ResponseWriter, r *http.Request) {

	order := &models.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	log.Println(err)
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
	orderId, err := strconv.Atoi(vars["id"])
	if err != nil {

		utils.Respond(w, utils.Message(false, "Invalid request"))
		return

	}
	response := utils.Message(true, "Order Details")

	response["order"] = models.GetOrder(orderId)
	utils.Respond(w, response)
}

var UpdateOrder = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, err := strconv.Atoi(vars["id"])
	if err != nil {

		utils.Respond(w, utils.Message(false, "Invalid request"))
		return

	}
	order := &models.Order{}
	err = json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := order.Update(orderId)
	utils.Respond(w, resp)

}
var DeleteOrder = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, err := strconv.Atoi(vars["id"])
	if err != nil {

		utils.Respond(w, utils.Message(false, "Invalid request"))
		return

	}
	utils.Respond(w, models.DeleteOrder(orderId))
}
var GetUserOrders = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {

		utils.Respond(w, utils.Message(false, "Invalid request"))
		return

	}
	response := utils.Message(true, "Order Details")
	var orders []models.Order
	// models.GetDB().Find(&orders)
	err = models.GetDB().Where("user_id = ?", userID).Find(&orders).Error
	if err != nil { //Order not found!
		response["order"] = nil
		utils.Respond(w, response)
		return
	}
	// log.Println(orders)
	response["order"] = orders
	utils.Respond(w, response)

}

var GetUserOrder = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	orderID, err := strconv.Atoi(vars["orderID"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	response := utils.Message(true, "Order Details")
	var order models.Order
	err = models.GetDB().Where("user_id = ? AND id = ?", userID, orderID).Find(&order).Error
	if err != nil { //Order not found!
		response["order"] = nil
		utils.Respond(w, response)
		return
	}
	response["order"] = order
	utils.Respond(w, response)

}
