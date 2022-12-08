package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Users GET call")

	db := Connect()
	var user []User
	err := db.Find(&user).Error
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(user)
	db.Close()
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get User By Id GET call")

	db := Connect()
	var user User
	id := mux.Vars(r)["id"]

	err := db.First(&user, id).Error
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(user)
	db.Close()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create User POST call")

	db := Connect()
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	e := db.Create(&user).Error
	if e != nil {
		panic(e)
	}

	json.NewEncoder(w).Encode(user)
	db.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update User by ID PATCH call")
	db := Connect()
	var user User
	var u User
	id := mux.Vars(r)["id"]

	err := db.First(&user, id).Error
	if err != nil {
		panic(err)
	}

	e := json.NewDecoder(r.Body).Decode(&u)
	if e != nil {
		panic(e)
	}
	user.Name = u.Name
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
	db.Close()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete User by ID DELETE call")

	db := Connect()
	var user User
	id := mux.Vars(r)["id"]

	err := db.First(&user, id).Error
	if err != nil {
		panic(err)
	}
	db.Delete(&user)
	json.NewEncoder(w).Encode(user)
	db.Close()
}
