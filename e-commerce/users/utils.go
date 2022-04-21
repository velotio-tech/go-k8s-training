package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	User_id  string `json:"user_id"`
	Username string `json:"username"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8090/users"
	res := makeGetRequest(url)
	fmt.Fprint(w, res)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	values := map[string]string{"user_id": user.User_id, "username": user.Username}
	jsonValue, _ := json.Marshal(values)
	url := "http://localhost:8090/addUser"

	res := makePostRequest(url, jsonValue)
	fmt.Fprint(w, string(res))
}

func GetAllUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	values := map[string]string{"user_id": user_id}
	jsonValue, _ := json.Marshal(values)
	url := "http://localhost:8090/getallorders"

	res := makePostRequest(url, jsonValue)
	fmt.Fprint(w, string(res))

}
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	//
	vars := mux.Vars(r)
	order_id := vars["order_id"]
	new_product_id := vars["updated_product_id"]

	values := map[string]string{"user_id": order_id, "product_id": new_product_id}
	jsonValue, _ := json.Marshal(values)

	url := "http://localhost:8090/updateorder"

	res := makePostRequest(url, jsonValue)
	fmt.Fprint(w, string(res))
}
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	order_id := vars["order_id"]

	values := map[string]string{"order_id": order_id}
	jsonValue, _ := json.Marshal(values)

	url := "http://localhost:8090/deleteorder"

	res := makePostRequest(url, jsonValue)
	fmt.Fprint(w, string(res))

}

func makePostRequest(url string, payload []byte) string {
	res, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return ""
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		return string(body)
	}
}

func makeGetRequest(url string) string {
	resp, err := http.Get("http://localhost:8090/users")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
