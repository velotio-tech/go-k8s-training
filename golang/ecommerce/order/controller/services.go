package controller

import (
	"fmt"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("something")
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("something")
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("something")
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("something")
}

func DeleteAllOrdersForUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("something")
}
