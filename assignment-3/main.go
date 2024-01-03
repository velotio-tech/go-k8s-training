package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pratikpjain/go-k8s-training/assignment3/db"
	userservice "github.com/pratikpjain/go-k8s-training/assignment3/services/users"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

func main() {
	connectToDB()
	handlers()
}

func connectToDB() {
	dbDesc := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db.ConnectDB(dbDesc)
	fmt.Println("Database connection is successful ...")
}

func handlers() {
	r := mux.NewRouter()

	r.HandleFunc("/", HealthCheck)
	r.HandleFunc("/users", userservice.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{username}", userservice.GetUserByUserName).Methods("GET")
	r.HandleFunc("/user", userservice.AddNewUser).Methods("POST")
	r.HandleFunc("/user", userservice.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{username}", userservice.DeleteUser).Methods("DELETE")

	r.HandleFunc("/user/{username}/order", userservice.CreateNewOrderByUserName).Methods("POST")
	r.HandleFunc("/user/{username}/orders", userservice.GetAllOrdersByUserName).Methods("GET")
	r.HandleFunc("/user/{username}/order/{orderid}", userservice.GetOrderByUserNameOrderID).Methods("GET")
	r.HandleFunc("/user/{username}/order", userservice.UpdateOrderByUserName).Methods("PUT")
	r.HandleFunc("/user/{username}/order/{orderid}", userservice.DeleteOrderByUserName).Methods("DELETE")

	fmt.Println("server started and running on port 8080 ...")
	http.ListenAndServe(":8080", r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Site is running")
}
