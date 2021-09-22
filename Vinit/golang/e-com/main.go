package main

import (
	"e-com/user"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: homePage", r.URL)
	json.NewEncoder(w).Encode(struct {
		StatusCode int
		Data string }{
		200,
		"E-com Backend alive",
	})
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", healthCheck)
	myRouter.HandleFunc("/users", user.ReturnAllUsers).Methods("GET")
	myRouter.HandleFunc("/users", user.AddUser).Methods("POST")
	myRouter.HandleFunc("/users/{uid}", user.ReturnUser).Methods("GET")
	myRouter.HandleFunc("/users/{uid}/{category}", user.HandleAllOrders).Methods("GET")
	myRouter.HandleFunc("/users/{uid}/{category}/{cid}", user.HandleSingleOrder).Methods("GET")
	myRouter.HandleFunc("/users/{uid}/order", user.AddOrder).Methods("POST")
	log.Println("Starting the Server for the e-com app")
	log.Fatal(http.ListenAndServe(":8009", myRouter))
}

func main() {
	user.IDGEN = 0
	handleRequests()
}
