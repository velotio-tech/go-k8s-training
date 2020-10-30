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

// PostOrder creates order for specific user - C
func (u *Users) PostOrder(w http.ResponseWriter, r *http.Request) {
}

// GetOrder gets specific order of specific user - R
func (u *Users) GetOrder(w http.ResponseWriter, r *http.Request) {
}

// GetOrders gets all orders of specific user - R
func (u *Users) GetOrders(w http.ResponseWriter, r *http.Request) {
}

// PutOrder updates specific order of specific user - U
func (u *Users) PutOrder(w http.ResponseWriter, r *http.Request) {
}

// DeleteOrder deletes specific order of specific user - D
func (u *Users) DeleteOrder(w http.ResponseWriter, r *http.Request) {
}

// DeleteOrders deletes all order of specific user - D
func (u *Users) DeleteOrders(w http.ResponseWriter, r *http.Request) {
}
