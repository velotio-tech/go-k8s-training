package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 20; i++ {
		err = db.Ping()
		if i == 20 {
			panic(err)
		}
		if err == nil {
			break
		} else if err != nil {
			fmt.Println(err)
			fmt.Println("DB Connection check. Retry count: ", i)
			time.Sleep(time.Second * 5)
		}
	}
	orderTableQuery := `CREATE TABLE IF NOT EXISTS orders(order_id SERIAL PRIMARY KEY, item_name varchar, user_id varchar)`
	_, err = db.Exec(orderTableQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Initialized Successfully!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

func main() {
	initDb()
	defer db.Close()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/orders", createOrder).Methods("POST")
	myRouter.HandleFunc("/api/orders", getOrders)
	myRouter.HandleFunc("/api/orders/{order_id}", deleteOrder).Methods("DELETE")
	myRouter.HandleFunc("/api/orders/{user_id}", getOrdersByUser)
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}

type orderSummary struct {
	OrderId  int
	ItemName string
	UserId   string
}

type orders struct {
	Orders []orderSummary
}

func getOrders(w http.ResponseWriter, req *http.Request) {
	orderlist := orders{}
	query := `SELECT order_id,item_name,user_id FROM orders`
	err := queryOrders(&orderlist, query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(orderlist)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func getOrdersByUser(w http.ResponseWriter, req *http.Request) {
	orderlist := orders{}
	vars := mux.Vars(req)
	user_id := vars["user_id"]
	err := queryOrdersByUser(&orderlist, user_id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(orderlist)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func deleteOrder(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	order_id := vars["order_id"]
	deleteStmt := `DELETE FROM orders WHERE order_id=$1`
	_, err := db.Exec(deleteStmt, order_id)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Order id: %v deleted", order_id)
}

func createOrder(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	var placeOrder map[string]string
	json.Unmarshal(reqBody, &placeOrder)
	createOrderStmt := `INSERT into orders (item_name, user_id) values ($1, $2)`
	rows, err := db.Query(createOrderStmt, placeOrder["item"], placeOrder["user"])
	if err != nil {
		log.Println(err)
	}

	getOrderStmt := `SELECT order_id,item_name,user_id FROM orders where user_id=$1 and item_name=$2`
	rows, _ = db.Query(getOrderStmt, placeOrder["user"], placeOrder["item"])

	defer rows.Close()
	order := orderSummary{}
	for rows.Next() {
		rows.Scan(&order.OrderId, &order.ItemName, &order.UserId)

	}
	out, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
}

func queryOrders(orderlist *orders, query string) error {
	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		order := orderSummary{}
		err = rows.Scan(&order.OrderId, &order.ItemName, &order.UserId)
		if err != nil {
			return err
		}
		orderlist.Orders = append(orderlist.Orders, order)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func queryOrdersByUser(orderlist *orders, userId string) error {
	rows, err := db.Query(`SELECT order_id,item_name,user_id FROM orders where user_id=$1`, userId)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		order := orderSummary{}
		err = rows.Scan(&order.OrderId, &order.ItemName, &order.UserId)
		if err != nil {
			return err
		}
		orderlist.Orders = append(orderlist.Orders, order)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
