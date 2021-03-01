package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/farkaskid/go-k8s-training/assignment3/db/helpers"
)

type Userdata struct {
	Name string
}

func CreateUserHandler(resp http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var userdata Userdata

	err := decoder.Decode(&userdata)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("Failed to decode user data: " + err.Error()))

		return
	}

	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("Failed to establish db connection"))

		return
	}

	defer db.Close()

	err = helpers.CreateUser(db, userdata.Name)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("Failed to create user"))

		return
	}

	resp.WriteHeader(200)
	resp.Write([]byte("Done!"))
}
