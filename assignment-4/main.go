package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/practice/db"
	"github.com/practice/service/user_service"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "root"
// 	dbname   = "store"
// )

func main() {
	connectToDB()
	handlers()
}

func connectToDB() {
	dbDesc := "postgres://postgres:root@db:5433/store?sslmode=disable"
	db.ConnectDB(dbDesc)
	fmt.Println("Database connection is successful ...")
}

func handlers() {
	r := mux.NewRouter()

	r.HandleFunc("/", EntryPoint)
	r.HandleFunc("/users", user_service.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{username}", user_service.GetUserByUserName).Methods("GET")
	r.HandleFunc("/user", user_service.AddNewUser).Methods("POST")
	r.HandleFunc("/user", user_service.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{username}", user_service.DeleteUser).Methods("DELETE")

	r.HandleFunc("/user/{username}/order", user_service.CreateNewOrderByUserName).Methods("POST")
	r.HandleFunc("/user/{username}/orders", user_service.GetAllOrdersByUserName).Methods("GET")
	r.HandleFunc("/user/{username}/order/{orderid}", user_service.GetOrderByUserNameOrderID).Methods("GET")
	r.HandleFunc("/user/{username}/order", user_service.UpdateOrderByUserName).Methods("PUT")
	r.HandleFunc("/user/{username}/order/{orderid}", user_service.DeleteOrderByUserName).Methods("DELETE")

	fmt.Println("server started and running on port 8080 ...")
	http.ListenAndServe(":8080", r)
}

func EntryPoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Default Route")
}
