package controller

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all users")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get user")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create user")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update user")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete user")
}
