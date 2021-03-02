package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/farkaskid/go-k8s-training/assignment3/db/helpers"
)

type Userdata struct {
	Name string
}

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
	decoder := json.NewDecoder(req.Body)

	var userdata Userdata

	err := decoder.Decode(&userdata)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	defer db.Close()

	err = helpers.CreateUser(db, userdata.Name)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Done!"))
}

func getUserHandler(resp http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(strings.Split(req.URL.Path, "/")[2])

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	users, err := helpers.GetUser(db, id)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	usersJSON, err := json.Marshal(users)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(usersJSON))
}

func updateUserHandler(resp http.ResponseWriter, req *http.Request) {
	type userUpdateData struct {
		Id      int
		NewName string
	}

	var updatedUserData userUpdateData
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&updatedUserData)
	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	defer db.Close()

	err = helpers.UpdateUser(db, updatedUserData.Id, updatedUserData.NewName)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Done!"))
}

func deleteUserHandler(resp http.ResponseWriter, req *http.Request) {

}
