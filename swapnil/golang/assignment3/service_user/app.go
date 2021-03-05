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

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}

func (a *App) Initialize(host, user, password, dbname, port string) {
	connectionString :=
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	var err error
	a.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
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
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var u User
	a.DB.First(&u, id)
	if uint(id) != u.ID {
		respondWithError(w, http.StatusNotFound, "user id is not found")
		return
	}
	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	a.DB.Find(&users)
	fmt.Println("{}", users)
	respondWithJSON(w, http.StatusOK, users)
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	fmt.Println(u)
	if len(u.Email) == 0 || len(u.Name) == 0 {
		respondWithError(w, http.StatusBadRequest, "name and email is required")
		return
	}
	result := a.DB.Create(&u)
	if result.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(result.Error))
		return
	}
	respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var u User
	a.DB.First(&u, id)
	if uint(id) != u.ID {
		respondWithError(w, http.StatusNotFound, "user id is not found")
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	if len(u.Name) == 0 {
		respondWithError(w, http.StatusBadRequest, "name not provided")
		return
	}

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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	result := a.DB.Delete(&User{}, id)
	if result.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(result.Error))
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) getUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	remoteURL := fmt.Sprintf("http://service_order:8011/orders?userid=%v", id)
	fmt.Println(remoteURL)
	resp, err := http.Get(remoteURL)
	fmt.Printf("%+v\n", resp)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	var u map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	// adding user id in body
	u["userid"] = id
	encoded, _ := json.Marshal(u)
	d := bytes.NewReader(encoded)
	remoteURL := "http://service_order:8011/orders"
	fmt.Println(remoteURL)
	client := &http.Client{}
	req, err := http.NewRequest("POST", remoteURL, d)
	fmt.Printf("%+v\n", req)
	if err != nil {
		fmt.Println("request creating")
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("==>other api")
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
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
	remoteURL := fmt.Sprintf("http://service_order:8011/orders/%v/%v", id, orderid)
	fmt.Println(remoteURL)
	resp, err := http.Get(remoteURL)
	fmt.Printf("%+v\n", resp)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	remoteURL := fmt.Sprintf("http://service_order:8011/orders/%v", id)
	fmt.Println(remoteURL)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", remoteURL, nil)
	fmt.Printf("%+v\n", req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
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

	remoteURL := fmt.Sprintf("http://service_order:8011/orders/%v/%v", id, orderid)
	fmt.Println(remoteURL)
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", remoteURL, nil)
	fmt.Printf("%+v\n", req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
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

	remoteURL := fmt.Sprintf("http://service_order:8011/orders/%v/%v", id, orderid)
	fmt.Println(remoteURL)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", remoteURL, r.Body)
	fmt.Printf("%+v\n", req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint(err))
		return
	}
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	respondWithJSON(w, resp.StatusCode, data)
}

func (a *App) initializeRoutes() {
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
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
