package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/thisisprasad/go-k8s-training/Prasad/assignment3/users/controllers"
)

func main() {
	// NewUserApp().StartApplication()

	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET") // /user's'
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/exists/{userId}", controllers.UserExists).Methods("GET")
	//	User order routes
	router.HandleFunc("/user/order", controllers.CreateUserOrder).Methods("POST")
	router.HandleFunc("/user/{userId}/orders", controllers.GetAllUserOrders).Methods("GET")
	router.HandleFunc("/user/{userId}/orders", controllers.DeleteUserOrders).Methods("DELETE")
	router.HandleFunc("/user/{userId}/order/{orderId}", controllers.DeleteUserOrder).Methods("DELETE")

	port := os.Getenv("USER_SERVER_PORT")
	if port == "" {
		port = "8090"
	}

	log.Println("Starting user server at port:", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalln(err)
	}
}
