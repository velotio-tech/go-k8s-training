package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

var GetUserOrders = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user := models.GetUser(vars["id"])
	if user == nil {
		utils.Respond(w, utils.Message(false, "User does not exits"))
		return
	}
	var orderAPI = models.GetOrderURL()
	resp, err := http.Get(orderAPI + "/api/users/" + userID + "/orders")
	if err != nil {
		utils.Respond(w, utils.Message(false, "Order server is not responding. Try again after sometime"))
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Order server is not responding. Try again after sometime"))
		log.Fatalln(err)
		return
	}

	orders := make(map[string]interface{})

	jsonErr := json.Unmarshal(body, &orders)

	if jsonErr != nil {
		utils.Respond(w, utils.Message(false, "Order server is not responding. Try again after sometime"))
		log.Fatalln(err)
	}

	response := utils.Message(true, "User Orders")
	response["orders"] = orders["order"]
	utils.Respond(w, response)

}

var GetUserOrder = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user := models.GetUser(vars["id"])
	if user == nil {
		utils.Respond(w, utils.Message(false, "User does not exits"))
		return
	}
	var orderAPI = models.GetOrderURL()
	resp, err := http.Get(orderAPI + "/api/users/" + userID + "/orders/" + vars["orderID"])
	if err != nil {
		log.Println(err)
		utils.Respond(w, utils.Message(false, "Order server is not responding. Try again after sometime"))
		return
	}
	log.Println(err)
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		utils.Respond(w, utils.Message(false, "Order server is not responding. Try again after sometime"))

		return
	}

	orders := make(map[string]interface{})

	jsonErr := json.Unmarshal(body, &orders)

	if jsonErr != nil {
		utils.Respond(w, utils.Message(false, "Order server is not responding. Try again after sometime"))

	}

	response := utils.Message(true, "User Orders")
	response["orders"] = orders["order"]
	utils.Respond(w, response)

}
