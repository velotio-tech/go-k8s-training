package main

import (
	"database/sql"
	"e-commerce/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting request: GET /users")

	db := database.OpenDbConnection()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var users []User

	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		users = append(users, User{Id: id, Name: name})
	}

	json.NewEncoder(w).Encode(users)

	db.Close()
	log.Print("Completed request: GET /users")
	log.Print("-------------------------------------------------")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting request: GET /users/:id")

	params := mux.Vars(r)
	db := database.OpenDbConnection()
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", params["id"])

	var id int
	var name string

	err := row.Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No record found!", 404)
		} else {
			panic(err)
		}
	} else {
		user := User{Id: id, Name: name}
		json.NewEncoder(w).Encode(user)
	}

	db.Close()
	log.Print("Completed request GET /users/:id")
	log.Print("-------------------------------------------------")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting request: POST /users")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	db := database.OpenDbConnection()

	sqlQuery := "INSERT INTO users (name) VALUES ($1)"
	_, err = db.Exec(sqlQuery, user.Name)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "New user created!")

	db.Close()
	log.Print("Completed request POST /users")
	log.Print("-------------------------------------------------")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting request: PATCH /users/:id")

	params := mux.Vars(r)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	db := database.OpenDbConnection()

	sqlQuery := "UPDATE users set name = $1 where id = $2"
	result, err := db.Exec(sqlQuery, user.Name, params["id"])
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	} else if rowsAffected == 0 {
		http.Error(w, "Failed to update user record!", 412)
	} else {
		fmt.Fprintf(w, "User updated successfully!")
	}

	db.Close()
	log.Print("Completed request PATCH /users/:id")
	log.Print("-------------------------------------------------")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting request: DELETE /users/:id")

	params := mux.Vars(r)
	db := database.OpenDbConnection()

	sqlQuery := "DELETE FROM users where id = $1"
	result, err := db.Exec(sqlQuery, params["id"])
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	} else if rowsAffected == 0 {
		http.Error(w, "Failed to delete user record!", 412)
	} else {
		fmt.Fprintf(w, "User record deleted successfully!")
	}

	db.Close()
	log.Print("Completed request DELETE /users/:id")
	log.Print("-------------------------------------------------")
}
