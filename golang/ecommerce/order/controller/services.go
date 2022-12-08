package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	var order []Order
	id := mux.Vars(r)["user_id"]

	err := db.Where("user_id = ?", id).Find(&order).Error
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(order)
	db.Close()
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	var order Order
	id := mux.Vars(r)["user_id"]
	json.NewDecoder(r.Body).Decode(&order)
	order.UserId, _ = strconv.Atoi(id)
	db.Create(&order)
	json.NewEncoder(w).Encode(order)
	db.Close()
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	var order Order
	var payload Order
	user_id := mux.Vars(r)["user_id"]
	id := mux.Vars(r)["id"]
	db.Where("user_id = ?", user_id).Where("id = ?", id).Find(&order)

	json.NewDecoder(r.Body).Decode(&payload)
	order.Name = payload.Name
	order.Price = payload.Price
	db.Save(&order)
	json.NewEncoder(w).Encode(order)
	db.Close()
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	var order Order

	user_id := mux.Vars(r)["user_id"]
	id := mux.Vars(r)["id"]

	db.Where("user_id = ?", user_id).Where("id = ?", id).Find(&order)
	db.Delete(&order)
	json.NewEncoder(w).Encode(order)
	db.Close()
}

func DeleteAllOrdersForUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	var order []Order

	user_id := mux.Vars(r)["user_id"]
	db.Where("user_id = ?", user_id).Delete(&order)
	db.Close()
}
