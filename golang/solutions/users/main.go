package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type user struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Global var, not a good practise to share across requests as each can run in parallel.
type usersMap map[string]user

var allUsers usersMap

func listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /users called...")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allUsers)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /users called...")

	var u user
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	allUsers[u.Name] = u

	// Write added objects json to response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)

	fmt.Println(allUsers)
}

func checkUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["name"]
	fmt.Printf("HEAD /users/%v called...\n", userName)

	if _, ok := allUsers[userName]; ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["name"]
	fmt.Printf("GET /users/%v called...\n", userName)

	if u, ok := allUsers[userName]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["name"]
	fmt.Printf("DELETE /users/%v called...\n", userName)
	if u, ok := allUsers[userName]; ok {
		delete(allUsers, userName)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
		fmt.Printf("DELETED user %v.\n", userName)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	// Initially create space for 50 users.
	allUsers = make(usersMap, 50)

	// Setup the http router.
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", listUsers).Methods("GET")
	r.HandleFunc("/users/{name}", getUserInfo).Methods("GET")
	r.HandleFunc("/users/{name}", checkUser).Methods("HEAD")
	r.HandleFunc("/users/{name}", deleteUser).Methods("DELETE")

	// Setup the server.
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
