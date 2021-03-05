// app.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// struct to hold the service router and the db
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// user struct with default gorm model such as ID, CreatedAt ...
type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}

func (a *App) Initialize(host, user, password, dbname, port string) {
	// connect to database
	connectionString :=
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	var err error
	a.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// ensure db schema is right
	a.initialMigration()
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get single user hit\n")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	// get userid from vars
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// fetch user row with id
	var u User
	a.DB.First(&u, id)
	// row was not found
	if uint(id) != u.ID {
		respondWithError(w, http.StatusNotFound, "user id is not found")
		return
	}
	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	// query all users
	a.DB.Find(&users)
	fmt.Println("{}", users)
	respondWithJSON(w, http.StatusOK, users)
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	// decode body and populate the user struct
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	fmt.Println(u)
	// ensure email and name is provided
	if len(u.Email) == 0 || len(u.Name) == 0 {
		respondWithError(w, http.StatusBadRequest, "name and email is required")
		return
	}
	// create user row in db
	result := a.DB.Create(&u)
	if result.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(result.Error))
		return
	}
	respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get userid from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var u User
	// fetch user row with pk as userid
	a.DB.First(&u, id)
	if uint(id) != u.ID {
		respondWithError(w, http.StatusNotFound, "user id is not found")
	}
	// populate user struct with data from r body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	// ensure name is provided
	if len(u.Name) == 0 {
		respondWithError(w, http.StatusBadRequest, "name not provided")
		return
	}
	// update the user row in db
	res := a.DB.Save(&u)
	if res.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(res.Error))
		return
	}
	fmt.Println("res", res.RowsAffected, res.Error)

	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get user id from from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// delete user from db with pk
	result := a.DB.Delete(&User{}, id)
	if result.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(result.Error))
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) getUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get userid from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// create order service api url for getting user's orders
	remoteURL := fmt.Sprintf("http://service_order:8011/orders?userid=%v", id)
	fmt.Println(remoteURL)
	resp, err := http.Get(remoteURL)
	fmt.Printf("%+v\n", resp)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	// decode response body
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	respondWithJSON(w, http.StatusOK, data)
}

func (a *App) createOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get user id from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// decode request body into variable which represent order
	var u map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	// adding user id in body
	u["userid"] = id
	// encode with json
	encoded, _ := json.Marshal(u)
	// create object which supports reader interface
	d := bytes.NewReader(encoded)
	remoteURL := "http://service_order:8011/orders"
	fmt.Println(remoteURL)
	client := &http.Client{}
	// create new post request
	req, err := http.NewRequest("POST", remoteURL, d)
	fmt.Printf("%+v\n", req)
	if err != nil {
		fmt.Println("request creating")
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	// start request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("==>other api")
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	// decode the response body and pass same onto the client
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("parsing response body")
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	fmt.Println("success", data)
	respondWithJSON(w, resp.StatusCode, data)
}

func (a *App) getUserOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get user id from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// get order id from vars
	orderid, err := strconv.Atoi(vars["orderid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	// set service order api url
	remoteURL := fmt.Sprintf("http://service_order:8011/orders/%v/%v", id, orderid)
	fmt.Println(remoteURL)
	// request to order service
	resp, err := http.Get(remoteURL)
	fmt.Printf("%+v\n", resp)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	// decode response body and pass on to client
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	respondWithJSON(w, http.StatusOK, data)
}

func (a *App) deleteUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get user id from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// remote api url to delete orders
	remoteURL := fmt.Sprintf("http://service_order:8011/orders/%v", id)
	fmt.Println(remoteURL)
	// create new http clint
	client := &http.Client{}
	// create new request
	req, err := http.NewRequest("DELETE", remoteURL, nil)
	fmt.Printf("%+v\n", req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	// start the request
	resp, err := client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	// decode response body and pass on to client
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	respondWithJSON(w, resp.StatusCode, data)
}

func (a *App) deleteUserOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get user id from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// get order id from vars
	orderid, err := strconv.Atoi(vars["orderid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	// service order api url
	remoteURL := fmt.Sprintf("http://service_order:8011/orders/%v/%v", id, orderid)
	fmt.Println(remoteURL)
	// create new http client
	client := &http.Client{}
	// create new request
	req, err := http.NewRequest("DELETE", remoteURL, nil)
	fmt.Printf("%+v\n", req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	// start request with client created
	resp, err := client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	// decode the response body and pass onto client
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	respondWithJSON(w, resp.StatusCode, data)
}

func (a *App) updateUserOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get user id from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	orderid, err := strconv.Atoi(vars["orderid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	// service order api url
	remoteURL := fmt.Sprintf("http://service_order:8011/orders/%v/%v", id, orderid)
	fmt.Println(remoteURL)
	client := &http.Client{}
	// create req to service order
	req, err := http.NewRequest("PUT", remoteURL, r.Body)
	fmt.Printf("%+v\n", req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	// start the request
	resp, err := client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	// decode the response body and pass onto client
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	respondWithJSON(w, resp.StatusCode, data)
}

func (a *App) initializeRoutes() {

	// register handlers for paths
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/users", a.createUser).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.deleteUser).Methods("DELETE")
	a.Router.HandleFunc("/users/{id:[0-9]+}/orders", a.getUserOrders).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}/orders", a.createOrder).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}/orders/{orderid:[0-9]+}", a.getUserOrder).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}/orders", a.deleteUserOrders).Methods("DELETE")
	a.Router.HandleFunc("/users/{id:[0-9]+}/orders/{orderid:[0-9]+}", a.deleteUserOrder).Methods("DELETE")
	a.Router.HandleFunc("/users/{id:[0-9]+}/orders/{orderid:[0-9]+}", a.updateUserOrder).Methods("PUT")
}

func (a *App) initialMigration() {
	// Migrate the schema
	a.DB.AutoMigrate(&User{})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	// common function to return error response
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// common function to return json response
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
