package main

import (
	"e-commerce/database"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting request: /users")

	db := database.OpenDbConnection()
	rows, err := db.Query("SELECT * from users")
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
	log.Print("Completed request: /users")
	log.Print("-------------------------------------------------")
}
