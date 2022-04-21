package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	db := Connect()
	stmt, err := db.Prepare("INSERT INTO users VALUES (?, ?)")
	if err != nil {
		fmt.Println("unable to prepare statement", err)
	}
	_, err = stmt.Exec(user.User_id, user.Username)
	if err != nil {
		fmt.Fprint(w, "unable to add usr:", err)
	}
	fmt.Fprint(w, user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	TotalUsers := GetUsers(db)
	b, _ := json.Marshal(TotalUsers)
	fmt.Fprint(w, string(b))
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	db := Connect()
	Orders := GetOrders(db, user.User_id)
	B, _ := json.Marshal(Orders)
	fmt.Fprint(w, string(B))
}

func DeleteOrders(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	db := Connect()
	Orders := DeleteOrder(db,user.User_id)
	b, _ := json.Marshal(Orders)
	fmt.Fprint(w, string(b))
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	db := Connect()
	Orders := UpdateOrderHelper(db,user.User_id,user.Username)
	b, _ := json.Marshal(Orders)
	fmt.Fprint(w, string(b))
}