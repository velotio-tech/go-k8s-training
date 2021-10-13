package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
