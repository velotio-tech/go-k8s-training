package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Users is a http.Handler
type Users struct {
	l *log.Logger
}

// NewUsers creates a users handler with the given logger
func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

// GetUsers returns the users from the database
func (u *Users) GetUsers(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle GET Users")

	w.Header().Set("Content-Type", "application/json")

	// fetch the users from the database
	lu := data.GetUsers()

	// serialize the list to JSON
	err := lu.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetUserOrders returns all orders of specified user from the database
func (u *Users) GetUserOrders(w http.ResponseWriter, r *http.Request) {

	user := mux.Vars(r)["user"]

	u.l.Println("Handle GET User:", user)

	w.Header().Set("Content-Type", "application/json")

	// send request to orders service to get orders of user specified

	// serialize the list to JSON
	err := lu.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetUserOrder returns specified order of specified user from the database
func (u *Users) GetUserOrder(w http.ResponseWriter, r *http.Request) {

	user := mux.Vars(r)["user"]
	order := mux.Vars(r)["order"]

	u.l.Println("Handle GET User:", user)

	w.Header().Set("Content-Type", "application/json")

	// send request to orders service to get specified order of user specified

	// serialize the list to JSON
	err := lu.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
