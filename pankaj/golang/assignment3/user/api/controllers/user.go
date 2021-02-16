package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/user/api/models"
	"github.com/pankaj9310/go-k8s-training/pankaj/golang/assignment3/user/api/utils"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
	utils.Respond(w, resp)
}

var GetUsers = func(w http.ResponseWriter, r *http.Request) {

	var users []models.User
	models.GetDB().Find(&users)
	response := utils.Message(true, "List of all users")
	response["user"] = users
	utils.Respond(w, response)
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	response := utils.Message(true, "User Details")
	response["user"] = models.GetUser(vars["id"])
	utils.Respond(w, response)
}

var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := user.Update(id)
	utils.Respond(w, resp)

}
var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	utils.Respond(w, models.DeleteUser(id))
}
