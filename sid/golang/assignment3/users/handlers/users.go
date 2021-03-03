package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farkaskid/go-k8s-training/assignment3/users/helpers"
)

func UserHandler(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getUserHandler(resp, req)
	case http.MethodPost:
		createUserHandler(resp, req)
	case http.MethodPut:
		updateUserHandler(resp, req)
	case http.MethodDelete:
		deleteUserHandler(resp, req)
	default:
		resp.WriteHeader(http.StatusBadRequest)
	}
}

func createUserHandler(resp http.ResponseWriter, req *http.Request) {
	msg, err := helpers.CreateUser(req)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	response := make(map[string]string)
	response["msg"] = msg

	responseJSON, err := json.Marshal(response)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(responseJSON)
}

func getUserHandler(resp http.ResponseWriter, req *http.Request) {
	users, err := helpers.GetUsers(req)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	type user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	response := make(map[string][]user)
	response["users"] = make([]user, len(users))
	index := 0

	for id, name := range users {
		response["users"][index] = user{ID: id, Name: name}
		index++
	}

	responseJSON, err := json.Marshal(response)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(responseJSON)
}

func updateUserHandler(resp http.ResponseWriter, req *http.Request) {
	msg, err := helpers.CreateUser(req)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	response := make(map[string]string)
	response["msg"] = msg

	responseJSON, err := json.Marshal(response)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(responseJSON)
}

func deleteUserHandler(resp http.ResponseWriter, req *http.Request) {
	msg, err := helpers.CreateUser(req)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
	}

	response := make(map[string]string)
	response["msg"] = msg

	responseJSON, err := json.Marshal(response)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(responseJSON)
}
