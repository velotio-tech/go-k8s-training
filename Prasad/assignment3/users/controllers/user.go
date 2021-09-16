package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/thisisprasad/go-k8s-training/Prasad/assignment3/users/constants"
	"github.com/thisisprasad/go-k8s-training/Prasad/assignment3/users/models"
	"github.com/thisisprasad/go-k8s-training/Prasad/assignment3/users/utils"
)

//	fetches all the users of the system.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userDao := &models.User{}
	users, err := userDao.GetAllUsers()
	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "fail"
		response["message"] = "Error fetching users."
		utils.RespondAsJson(w, http.StatusInternalServerError, response)
	} else {
		utils.RespondAsJson(w, http.StatusOK, users)
	}
}

//	Creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		log.Println(err)
		utils.RespondAsJson(w,
			http.StatusInternalServerError,
			map[string]string{"error": "Error creating user"})
	}

	response := make(map[string]interface{})
	if user.Create() != nil {
		response["result"] = "fail"
		response["message"] = "Error creating user."
		utils.RespondAsJson(w, http.StatusUnprocessableEntity, response) //	HTTP: 422
	} else {
		response["result"] = "success"
		response["message"] = "User created successfully."
		response["user"] = make(map[string]interface{})
		response["user"].(map[string]interface{})["id"] = user.ID
		response["user"].(map[string]interface{})["email"] = user.Email
		utils.RespondAsJson(w, http.StatusOK, response)
	}
}

//	Updates the details of the user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		log.Println(err)
		utils.RespondAsJson(w,
			http.StatusInternalServerError,
			map[string]string{"error": "Error updating user"})
		return
	}

	response := make(map[string]interface{})
	if user.Update() != nil {
		response["result"] = "fail"
		response["message"] = "Error updating user."
		utils.RespondAsJson(w, http.StatusInternalServerError, response)
	} else {
		response["result"] = "success"
		response["message"] = "User updated successfully."
		response["user"] = user
		utils.RespondAsJson(w, http.StatusOK, response)
	}
}

//	Deletes a user from system
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
	}

	err = user.Delete(user.ID)
	response := make(map[string]interface{})
	if err != nil {
		response["result"] = "fail"
		response["message"] = "Error deleting user."
		utils.RespondAsJson(w, http.StatusUnprocessableEntity, response)
	} else {
		response["result"] = "success"
		response["message"] = "User deleted successfully"
		utils.RespondAsJson(w, http.StatusOK, response)
	}
}

//	Checks whether the user exists in the system based on user id.
func UserExists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Println(err)
		utils.RespondAsJson(w, http.StatusBadRequest, "Incorrect data sent.")
		return
	}
	userDao := &models.User{}
	exists := userDao.Exists(uint(userId))

	response := make(map[string]interface{})
	response["exists"] = exists
	utils.RespondAsJson(w, http.StatusOK, response)
}

//	creates an order for user
func CreateUserOrder(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf(os.Getenv("ORDER_SERVER_BASE_URL") + constants.CreateUserOrder)
	request, err := http.NewRequest(http.MethodPost, url, r.Body)
	request.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if resp.Body == nil {
		log.Println("resp.Body is nil")
	}

	body := utils.GetResponseJsonBody(resp)
	utils.RespondAsJson(w, http.StatusOK, body)
}

//	fetches all orders of an user
func GetAllUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf(os.Getenv("ORDER_SERVER_BASE_URL") + constants.FetchAllUserOrdersUrl + vars["userId"])
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body := utils.GetResponseJsonBody(resp)

	utils.RespondAsJson(w, http.StatusOK, body)
}

//	Deletes all orders of an user
func DeleteUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf(os.Getenv("ORDER_SERVER_BASE_URL") + constants.DelteAllUserOrdersUrl + vars["userId"])
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body := utils.GetResponseJsonBody(resp)
	utils.RespondAsJson(w, http.StatusOK, body)
}

//	Deletes an order of the user.
func DeleteUserOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := fmt.Sprintf(os.Getenv("ORDER_SERVER_BASE_URL") + constants.DeleteUserOrdersUrl + vars["orderId"])
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body := utils.GetResponseJsonBody(resp)
	utils.RespondAsJson(w, http.StatusOK, body)
}
