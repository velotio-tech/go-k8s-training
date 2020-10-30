package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/orders/models"
	"github.com/gorilla/mux"
)

// Orders is a http.Handler
type Orders struct {
	l *log.Logger
}

// NewOrders creates order handler with the given logger which can call all order handlers
func NewOrders(l *log.Logger) *Orders {
	return &Orders{l}
}

// PostUserOrder creates order for specific user - C
func (u *Orders) PostUserOrder(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle POST Order")

	vars := mux.Vars(r)
	userID, _ := strconv.ParseUint(vars["user"], 10, 64)
	body := BodyParser(r)
	var order models.Order
	err := json.Unmarshal(body, &order)

	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = models.CreateUserOrder(order, userID)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	ToJSON(w, "Order added successfully!", http.StatusCreated)

}

// GetUserOrders returns all orders of specific user from the database - R
func (u *Orders) GetUserOrders(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle GET Orders")

	vars := mux.Vars(r)
	userID, _ := strconv.ParseUint(vars["user"], 10, 64)
	orders := models.GetAll(userID)
	ToJSON(w, orders, http.StatusOK)
}

// GetUserOrder gets specific order of specified user from the database - R
func (u *Orders) GetUserOrder(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle GET Order")

	vars := mux.Vars(r)
	userID, _ := strconv.ParseUint(vars["user"], 10, 64)
	orderID, _ := strconv.ParseUint(vars["order"], 10, 64)
	order := models.GetByOrderID(userID, orderID)
	ToJSON(w, order, http.StatusOK)
}

// PutUserOrder updates specific order of specific user - U
func (u *Orders) PutUserOrder(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle PUT Order")

	vars := mux.Vars(r)
	userID, _ := strconv.ParseUint(vars["user"], 10, 64)
	orderID, _ := strconv.ParseUint(vars["order"], 10, 64)
	body := BodyParser(r)
	var order models.Order
	err := json.Unmarshal(body, &order)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	order.UserID = uint64(userID)
	order.OrderID = uint64(orderID)
	rows, err := models.UpdateUserOrder(order, userID, orderID)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	ToJSON(w, rows, http.StatusOK)
}

// DeleteUserOrder deletes specific order of specific user - D
func (u *Orders) DeleteUserOrder(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle DELETE Order")

	vars := mux.Vars(r)
	userID, _ := strconv.ParseUint(vars["user"], 10, 64)
	orderID, _ := strconv.ParseUint(vars["order"], 10, 64)
	_, err := models.DeleteOrder(userID, orderID)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	ToJSON(w, "Order deleted successfully!", http.StatusCreated)
}

// DeleteUserOrders deletes all orders of a specific user from the database - D
func (u *Orders) DeleteUserOrders(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle DELETE Orders")

	vars := mux.Vars(r)
	userID, _ := strconv.ParseUint(vars["user"], 10, 64)
	_, err := models.DeleteOrders(userID)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	ToJSON(w, "Orders deleted successfully!", http.StatusCreated)
}
