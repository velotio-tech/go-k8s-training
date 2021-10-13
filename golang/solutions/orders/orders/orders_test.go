package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	URL = "http://localhost:8080/"
)

func TestCreateOrder(t *testing.T) {
	userName := "user1"
	o := orderInfo{
		ProductName: "Car",
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(o)
	res, err := http.Post(URL+"users/"+userName+"/orders", "Content-Type: application/json", payload)
	if err != nil {
		t.Fatal("Check if server is down.")
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatal("Cannot create order.")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Body : %s", body)
}

func TestGetOrder(t *testing.T) {
	userName := "user1"
	expectedOrder := order{
		Id: 1,
		OrderInfo: orderInfo{
			ProductName: "Car",
		},
	}

	res, err := http.Get(URL + "users/" + userName + "/orders/" + strconv.Itoa(expectedOrder.Id))
	if err != nil {
		t.Fatal("Check if server is down.")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("Cannot get order.")
	}
	defer res.Body.Close()
	var o order
	err = json.NewDecoder(res.Body).Decode(&o)
	if err != nil {
		t.Fatal("Error parsing json response.", err)
	}
	if !cmp.Equal(expectedOrder, o) {
		t.Fatal("Invalid order returned.")
	}
}

func TestGetAllOrders(t *testing.T) {
	userName := "user1"
	expectedOrder := order{
		Id: 1,
		OrderInfo: orderInfo{
			ProductName: "Car",
		},
	}

	res, err := http.Get(URL + "users/" + userName + "/orders")
	if err != nil {
		t.Fatal("Check if server is down.")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("Cannot get user.")
	}
	defer res.Body.Close()
	var om ordersMap
	err = json.NewDecoder(res.Body).Decode(&om)
	if err != nil {
		t.Fatal("Error parsing json response.", err)
	}
	if len(om) != 1 ||
		expectedOrder.OrderInfo.ProductName != om[expectedOrder.Id].OrderInfo.ProductName {
		t.Fatal("Invalid Order list returned.")
	}
}

func TestDeleteOrder(t *testing.T) {
	userName := "user1"
	expectedOrder := order{
		Id: 1,
		OrderInfo: orderInfo{
			ProductName: "Car",
		},
	}

	client := &http.Client{}
	req, err :=
		http.NewRequest("DELETE",
			URL+"users/"+userName+"/orders/"+strconv.Itoa(expectedOrder.Id), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal("Cannot delete order.")
	}
}

func TestGetAllOrdersEmpty(t *testing.T) {
	userName := "user1"

	res, err := http.Get(URL + "users/" + userName + "/orders")
	if err != nil {
		t.Fatal("Check if server is down.")
	}
	if res.StatusCode != http.StatusOK {
		t.Fatal("Cannot get user.")
	}
	defer res.Body.Close()
	var om ordersMap
	err = json.NewDecoder(res.Body).Decode(&om)
	if err != nil {
		t.Fatal("Error parsing json response.", err)
	}
	fmt.Println(om)
	if len(om) != 0 {
		t.Fatal("Invalid Order list returned.")
	}
}
