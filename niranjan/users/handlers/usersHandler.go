package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/users/models"
	"github.com/gorilla/mux"
)

// Users is a http.Handler
type Users struct {
	l *log.Logger
}

// NewUsers creates a users handler with the given logger which can call all user handlers
func NewUsers(l *log.Logger) *Users {
	return &Users{l}
}

// PostUser creates user in the database - C
func (u *Users) PostUser(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle POST Users")

	body := BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = models.CreateUser(user)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	ToJSON(w, "User added successfully!", http.StatusCreated)
}

// GetUsers returns all users from the database - R
func (u *Users) GetUsers(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle GET Users")

	users := models.GetAll()
	ToJSON(w, users, http.StatusOK)

	/*
		w.Header().Set("Content-Type", "application/json")

		// fetch the users from the database
		lu := models.GetUsers()

		// serialize the list to JSON
		err := lu.ToJSON(w)
		if err != nil {
			http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		}*/
}

// GetUser gets specified user from the database - R
func (u *Users) GetUser(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle GET User")

	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["user"], 10, 64)
	user := models.GetByID(id)
	ToJSON(w, user, http.StatusOK)

	/*user := mux.Vars(r)["user"]

	u.l.Println("Handle GET User:", user)

	w.Header().Set("Content-Type", "application/json")

	// send request to orders service to get orders of user specified

	// serialize the list to JSON
	err := lu.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}*/
}

// PutUser updates user name &/or email in the database - U
func (u *Users) PutUser(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle PUT User")

	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["user"], 10, 32)
	body := BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user.ID = uint32(id)
	rows, err := models.UpdateUser(user)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	ToJSON(w, rows, http.StatusOK)
}

// DeleteUser deletes user from the database - D
func (u *Users) DeleteUser(w http.ResponseWriter, r *http.Request) {

	u.l.Println("Handle DELETE User")

	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["user"], 10, 64)
	_, err := models.Delete(id)
	if err != nil {
		ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
