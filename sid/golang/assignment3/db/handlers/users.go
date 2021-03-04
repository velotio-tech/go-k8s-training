package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/farkaskid/go-k8s-training/assignment3/db/helpers"
)

func UserHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	switch req.Method {
	case http.MethodGet:
		getUserHandler(resp, req, db)
	case http.MethodPost:
		createUserHandler(resp, req, db)
	case http.MethodPut:
		updateUserHandler(resp, req, db)
	case http.MethodDelete:
		deleteUserHandler(resp, req, db)
	default:
		resp.WriteHeader(http.StatusBadRequest)
	}
}

func createUserHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	type CreateUserdata struct {
		Name string
	}

	decoder := json.NewDecoder(req.Body)

	var userdata CreateUserdata

	err := decoder.Decode(&userdata)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	err = helpers.CreateUser(db, userdata.Name)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Done!"))
}

func getUserHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	var users map[int]string
	var err error
	var singleRequest bool

	if strings.HasSuffix(req.URL.Path, "/users") || strings.HasSuffix(req.URL.Path, "/users/") {
		idInts, err := helpers.GetIDs(req)

		if err != nil {
			ErrorHandler(resp, req, err, http.StatusBadRequest)
			return
		}

		if len(idInts) == 0 {
			users, err = helpers.GetAllUsers(db)
		} else {
			users, err = helpers.GetUsers(db, idInts)
		}
	} else {
		singleRequest = true
		id, err := strconv.Atoi(strings.Split(req.URL.Path, "/")[2])

		if err != nil {
			ErrorHandler(resp, req, err, http.StatusBadRequest)
			return
		}

		users, err = helpers.GetUser(db, id)
	}

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	if len(users) == 0 && singleRequest {
		http.NotFound(resp, req)
		return
	}

	usersJSON, err := json.Marshal(users)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(usersJSON))
}

func updateUserHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	type userUpdateData struct {
		ID      int
		NewName string
	}

	var updatedUserData userUpdateData
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&updatedUserData)
	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	err = helpers.UpdateUser(db, updatedUserData.ID, updatedUserData.NewName)

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Done!"))
}

func deleteUserHandler(resp http.ResponseWriter, req *http.Request, db *sql.DB) {
	var err error

	if strings.HasSuffix(req.URL.Path, "/orders") || strings.HasSuffix(req.URL.Path, "/orders/") {
		idInts, err := helpers.GetIDs(req)

		if err != nil {
			ErrorHandler(resp, req, err, http.StatusBadRequest)
			return
		}

		if len(idInts) == 0 {
			err = helpers.DeleteAllUsers(db)
		} else {
			err = helpers.DeleteUsers(db, idInts)
		}
	} else {
		id, err := strconv.Atoi(strings.Split(req.URL.Path, "/")[2])

		if err != nil {
			ErrorHandler(resp, req, err, http.StatusBadRequest)
			return
		}

		err = helpers.DeleteUser(db, id)
	}

	if err != nil {
		ErrorHandler(resp, req, err, http.StatusInternalServerError)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("Done!"))
}
