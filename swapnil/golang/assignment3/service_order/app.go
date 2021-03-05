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
	a.initialMigration()
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get single order hit\n")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	var u Order
	a.DB.First(&u, id)
	if uint(id) != u.ID {
		respondWithError(w, http.StatusNotFound, "order id is not found")
		return
	}
	if uint(userid) != u.UserID {
		respondWithError(w, http.StatusUnauthorized, "order does not belong to user")
		return
	}
	respondWithJSON(w, http.StatusOK, u)
}

func (a *App) getOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%+v\n", vars)
	userID := r.FormValue("userid")
	var orders []Order
	if len(userID) > 0 {
		fmt.Printf("user id is provided")
		a.DB.Where("user_id=?", userID).Find(&orders)
	} else {
		fmt.Printf("user id is not provided")
		a.DB.Find(&orders)
	}

	fmt.Println("{}", orders)
	data := map[string][]Order{}
	data["orders"] = orders
	respondWithJSON(w, http.StatusOK, data)
}

func (a *App) createOrder(w http.ResponseWriter, r *http.Request) {
	var u Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if u.UserID < 1 {
		respondWithError(w, http.StatusBadRequest, "Userid not provided")
		return
	}
	result := a.DB.Create(&u)
	if result.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(result.Error))
		return
	}
	respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) updateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	var u Order
	a.DB.First(&u, id)
	if uint(id) != u.ID {
		respondWithError(w, http.StatusNotFound, "order id is not found")
		return
	}
	if uint(userid) != u.UserID {
		respondWithError(w, http.StatusUnauthorized, "order does not belong to user")
		return
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

func (a *App) deleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	userid, err := strconv.Atoi(vars["userid"])
	var u Order
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	result1 := a.DB.Where("user_id=? and id=?", userid, id).First(&u)
	fmt.Println(u)
	if result1.RowsAffected == 0 {
		respondWithError(w, http.StatusUnauthorized, "order not found")
		return
	}

	result := a.DB.Delete(&Order{}, id)
	if result.Error != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint(result.Error))
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) deleteUserOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	result := a.DB.Where("user_id=?", userid).Delete(&Order{})
	if result.RowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "orders not found")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "deleted"})
}

func (a *App) initializeRoutes() {
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
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
