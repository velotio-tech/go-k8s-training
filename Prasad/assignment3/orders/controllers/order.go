package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/thisisprasad/go-k8s-training/Prasad/assignment3/orders/models"
	"github.com/thisisprasad/go-k8s-training/Prasad/assignment3/orders/utils"
)

//	Creates a new order for the user
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := &models.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		log.Println(err)
		utils.RespondAsJson(w,
			http.StatusBadRequest,
			map[string]string{"error": "Error creating order."})
		return
	}

	response := make(map[string]interface{})
	if !userExists(int(order.UserId)) {
		response["result"] = "fail"
		response["message"] = "User not found."
		utils.RespondAsJson(w, http.StatusNotFound, response)
	} else if order.Create() != nil {
		response["result"] = "fail"
		response["message"] = "Error creating order"
		utils.RespondAsJson(w, http.StatusInternalServerError, response)
	} else {
		response["result"] = "success"
		response["message"] = "Order placed successfully."
		response["order"] = make(map[string]interface{})
		response["order"].(map[string]interface{})["id"] = order.ID
		response["order"].(map[string]interface{})["userId"] = order.UserId
		utils.RespondAsJson(w, http.StatusOK, response)
	}
}

//	Private function. Checks whether user exists in system with the given userId.
func userExists(userId int) bool {
	resp, err := http.Get(os.Getenv("USER_SERVER_BASE_URL") + os.Getenv("USER_SERVER_EXISTS_URL") + strconv.Itoa(userId))
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var bodyJson map[string]bool
	json.Unmarshal(body, &bodyJson)

	return bodyJson["exists"]
}

//	Fetches all the orders of the user
func GetUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		utils.RespondAsJson(w, http.StatusBadRequest, "Incorrect data sent.")
		return
	}
	ordersDao := &models.Order{}
	userOrders, err := ordersDao.GetUserOrders(uint(userId))

	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "fail"
		response["message"] = "Error fetching orders for the user."
		utils.RespondAsJson(w, http.StatusInternalServerError, response)
	} else {
		utils.RespondAsJson(w, http.StatusOK, userOrders)
	}
}

//	Deletes a single order of an user.
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		utils.RespondAsJson(w, http.StatusBadRequest, "Incorrect data sent.")
		return
	}
	ordersDao := &models.Order{}
	err = ordersDao.Delete(uint(id))

	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "fail"
		response["message"] = "Error deleting the order"
		utils.RespondAsJson(w, http.StatusInternalServerError, response)
	} else {
		response["result"] = "success"
		response["message"] = "Order deleted successfully."
		utils.RespondAsJson(w, http.StatusOK, response)
	}
}

//	Deletes all the orders of a user
func DeleteAllOrdersOfUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		utils.RespondAsJson(w, http.StatusBadRequest, "Incorrect data sent.")
		return
	}
	ordersDao := &models.Order{}
	err = ordersDao.DeleteUserOrders(uint(userId))

	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "fail"
		response["message"] = "Failed to delete user's orders."
		utils.RespondAsJson(w, http.StatusInternalServerError, response)
	} else {
		response["result"] = "success"
		response["message"] = "User's orders deleted successfully."
		utils.RespondAsJson(w, http.StatusOK, response)
	}
}
