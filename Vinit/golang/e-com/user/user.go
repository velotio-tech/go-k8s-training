package user

import (
	"e-com/database"
	"e-com/helper"
	"e-com/orders"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type User struct {
	Username string `json:"username" ,bson:"username"`
	UserId int `json:"Uid" ,bson:"Uid"`
	Orders []orders.Order `json:"orders" ,bson:"orders"`
}

var IDGEN int
var URI = "http://order:9009"

func ReturnUser(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: return User")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["uid"]
	fmt.Println(userId)
	data := database.GetAllUsers(userId)
	resp := new(interface{})
	json.Unmarshal(data, resp)
	json.NewEncoder(w).Encode(resp)
}

func ReturnAllUsers(w http.ResponseWriter, r *http.Request)  {
	log.Println("Endpoint Hit: Return All Users", r.URL)
	w.Header().Set("Content-Type", "application/json")
	data := database.GetAllUsers("")
	resp := new(interface{})
	json.Unmarshal(data, resp)
	json.NewEncoder(w).Encode(resp)
}

func AddUser (w http.ResponseWriter, r *http.Request) {
	var user User
	var data []byte

	log.Println("Endpoint Hit: ADD User: creating user")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.UserId = IDGEN
	user.Orders = []orders.Order{}
	data, _ = json.Marshal(user)
	rData := database.WriteToDB(data)
	IDGEN += 1
	resp := new(interface{})
	json.Unmarshal(rData, resp)
	json.NewEncoder(w).Encode(resp)
}

func HandleAllOrders (w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: HandleAllOrders")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["uid"]
	operation := vars["category"]
	params := 	"uid=" + url.QueryEscape(userId)
	URL := URI+ "/" + operation
	path := URL+"?"+params
	fmt.Println("Sending request to: ", path)
	data := helper.MakeRequest("GET", path, r.Body)
	resp := new(interface{})
	json.Unmarshal(data, resp)
	json.NewEncoder(w).Encode(resp)
}

func HandleSingleOrder (w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: HandleSingleOrders")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["uid"]
	operation := vars["category"]
	orderId := vars["cid"]
	params := 	"uid=" + url.QueryEscape(userId) + "&" + "oid=" + url.QueryEscape(orderId)
	URL := URI+ "/" + operation
	path := URL+"?"+params
	log.Println("Sending request to: ", path)
	data := helper.MakeRequest("GET", path, r.Body)
	resp := new(interface{})
	json.Unmarshal(data, resp)
	json.NewEncoder(w).Encode(resp)

}

func AddOrder (w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: Add Orders")
	var order orders.Order
	var data []byte

	_ = json.NewDecoder(r.Body).Decode(&order)
	fmt.Println(order)
	vars := mux.Vars(r)
	userId := vars["uid"]

	order.OrderId = rand.Intn(10000000)
	order.LastModified = time.Now().String()
	URL := URI + "/add/" + userId

	postbody, _ := json.Marshal(order)
	reqBody := ioutil.NopCloser(strings.NewReader(string(postbody)))
	w.Header().Set("Content-Type", "application/json")
	log.Println("Sending Request to: ",URL, r.Body)
	data = helper.MakeRequest("POST", URL, reqBody )
	fmt.Println(string(data))
	response := new(interface{})
	json.Unmarshal(data, response)
	json.NewEncoder(w).Encode(response)
}

