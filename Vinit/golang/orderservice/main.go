package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"orderservice/handler"
)

func healthCheck(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: homePage", r.URL)
	json.NewEncoder(w).Encode(struct {
		StatusCode int
		Data string }{
		200,
		"Alive",
	})
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", healthCheck)
	myRouter.HandleFunc("/delete", handler.HandleDelete).Methods("GET")
	myRouter.HandleFunc("/order", handler.HandleOrder).Methods("GET")
	myRouter.HandleFunc("/add/{uid}", handler.AddNewOrder).Methods("POST")
	myRouter.HandleFunc("/{random}", handler.NotFoundHandler)
	log.Println("Starting the Server for the order app")
	log.Fatal(http.ListenAndServe(":9009", myRouter))
}

func main() {
	handleRequests()
}

