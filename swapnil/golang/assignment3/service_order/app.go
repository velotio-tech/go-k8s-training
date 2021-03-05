// app.go

package main

import (
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

type Order struct {
	gorm.Model
	Name     string `gorm:"not null"` // name of the order item
	UserID   uint   `gorm:"not null"`
	Quantity int    `gorm:"default:1"`
	Unit     string `gorm:"default:'units'`
}

func (a *App) Initialize(host, user, password, dbname, port string) {
	connectionString :=
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	var err error
	a.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// create tables and columns if not already there in the database
	a.initialMigration()
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get order id from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	// get userid from vars
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	var u Order
	// fetch the order
	a.DB.First(&u, id)
	// u.ID is not equals to what user has requested
	if uint(id) != u.ID {
		respondWithError(w, http.StatusNotFound, "order id is not found")
		return
	}
	// if order is there but user id is different returns error
	if uint(userid) != u.UserID {
		respondWithError(w, http.StatusUnauthorized, "order does not belong to user")
		return
	}
	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) getOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%+v\n", vars)
	// get userid from querystring
	userID := r.FormValue("userid")
	var orders []Order
	if len(userID) > 0 {
		// filter for orders related to that user only
		fmt.Printf("user id is provided")
		a.DB.Where("user_id=?", userID).Find(&orders)
	} else {
		// no filter all orders will be returned
		fmt.Printf("user id is not provided")
		a.DB.Find(&orders)
	}

	fmt.Println("{}", orders)
	// creating map to hold the data array of orders
	data := map[string][]Order{}
	data["orders"] = orders
	respondWithJSON(w, http.StatusOK, data)
}

func (a *App) createOrder(w http.ResponseWriter, r *http.Request) {
	var u Order
	// decode the body of request and try to create order struct out of it
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		// json is not well formed or problem with data
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	// closing the body
	defer r.Body.Close()
	// userid must be provided in order to create the order
	if u.UserID < 1 {
		respondWithError(w, http.StatusBadRequest, "Userid not provided")
		return
	}
	// create order row
	result := a.DB.Create(&u)
	if result.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(result.Error))
		return
	}
	respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) updateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get orderid from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	// get userid from vars
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	var u Order
	//fetch single row with the orderid
	a.DB.First(&u, id)
	// order not fetched
	if uint(id) != u.ID {
		respondWithError(w, http.StatusNotFound, "order id is not found")
		return
	}
	// order is fetched but user id might be different
	if uint(userid) != u.UserID {
		respondWithError(w, http.StatusUnauthorized, "order does not belong to user")
		return
	}
	// decode the body and populated order struct
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	// making sure name is provided
	if len(u.Name) == 0 {
		respondWithError(w, http.StatusBadRequest, "name not provided")
		return
	}
	// updating order row in database
	res := a.DB.Save(&u)
	if res.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(res.Error))
		return
	}
	fmt.Println("res", res.RowsAffected, res.Error)

	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) deleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// get order id from vars
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// get userid from vars
	userid, err := strconv.Atoi(vars["userid"])
	var u Order
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	// query first order with userid and orderid
	result1 := a.DB.Where("user_id=? and id=?", userid, id).First(&u)
	fmt.Println(u)
	//check if order present
	if result1.RowsAffected == 0 {
		respondWithError(w, http.StatusUnauthorized, "order not found")
		return
	}
	// delete the order from db
	result := a.DB.Delete(&Order{}, id)
	if result.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(result.Error))
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) deleteUserOrders(w http.ResponseWriter, r *http.Request) {
	// deletes al orders of a user
	vars := mux.Vars(r)
	// get user id from vars
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	// delete the orders with userid
	result := a.DB.Where("user_id=?", userid).Delete(&Order{})
	if result.RowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "orders not found")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "deleted"})
}

func (a *App) initializeRoutes() {
	// register handlers for api paths
	a.Router.HandleFunc("/orders", a.getOrders).Methods("GET")
	a.Router.HandleFunc("/orders", a.createOrder).Methods("POST")
	a.Router.HandleFunc("/orders/{userid:[0-9]+}/{id:[0-9]+}", a.getOrder).Methods("GET")
	a.Router.HandleFunc("/orders/{userid:[0-9]+}/{id:[0-9]+}", a.updateOrder).Methods("PUT")
	a.Router.HandleFunc("/orders/{userid:[0-9]+}/{id:[0-9]+}", a.deleteOrder).Methods("DELETE")
	a.Router.HandleFunc("/orders/{userid:[0-9]+}", a.deleteUserOrders).Methods("DELETE")

}

func (a *App) initialMigration() {
	// Migrate the schema
	a.DB.AutoMigrate(&Order{})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	// common function to return error
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// common function to dump data to json and write it to the response
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
