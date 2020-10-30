package handlers

import (
	"log"
	"net/http"
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
func (u *Users) PostUserOrder(w http.ResponseWriter, r *http.Request) {
	rows := ""
	// make Post request to Order microservices
	ToJSON(w, rows, http.StatusOK)
}

// GetUserOrder gets specific order of specific user - R
func (u *Users) GetUserOrder(w http.ResponseWriter, r *http.Request) {
}

// GetUserOrders gets all orders of specific user - R
func (u *Users) GetUserOrders(w http.ResponseWriter, r *http.Request) {
}

// PutUserOrder updates specific order of specific user - U
func (u *Users) PutUserOrder(w http.ResponseWriter, r *http.Request) {
}

// DeleteUserOrder deletes specific order of specific user - D
func (u *Users) DeleteUserOrder(w http.ResponseWriter, r *http.Request) {
}

// DeleteUserOrders deletes all order of specific user - D
func (u *Users) DeleteUserOrders(w http.ResponseWriter, r *http.Request) {
}
