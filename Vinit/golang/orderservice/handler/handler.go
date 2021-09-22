package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"orderservice/database"
)

func HandleDelete (w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	oid := r.FormValue("oid")
	rData := database.DeleteOrders(uid, oid)
	resp := new(interface{})
	json.Unmarshal(rData, resp)
	json.NewEncoder(w).Encode(resp)
}

func HandleOrder (w http.ResponseWriter, r *http.Request) {
	log.Println("Handle Order Called")
	uid := r.FormValue("uid")
	oid := r.FormValue("oid")
	rData := database.GetOrders(uid, oid)
	resp := new(interface{})
	json.Unmarshal(rData, resp)
	json.NewEncoder(w).Encode(resp)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		StatusCode int
		Data string }{
		404,
		"Not Found",
	})
}

func AddNewOrder (w http.ResponseWriter, r *http.Request) {
	log.Println("Adding new to order to db, AddNewOrder called")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId := vars["uid"]
	orderDetail, _ := ioutil.ReadAll(r.Body)
	rData := database.AddOrder(userId, orderDetail)
	resp := new(interface{})
	json.Unmarshal(rData, resp)
	json.NewEncoder(w).Encode(resp)
}