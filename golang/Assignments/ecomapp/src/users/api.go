package main

import (
	"bytes"
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
	dbhost           = "DBHOST"
	dbport           = "DBPORT"
	dbuser           = "DBUSER"
	dbpass           = "DBPASS"
	dbname           = "DBNAME"
	orderServiceHost = "ORDERSERVICEHOST"
	orderServicePort = "ORDERSERVICEPORT"
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
	orderTableQuery := `CREATE TABLE IF NOT EXISTS users(user_id SERIAL PRIMARY KEY, user_name varchar NOT NULL, contact_no varchar UNIQUE NOT NULL)`
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
	//placeorder required json with following structure '{"item":"milk-shake", "user":"2"}'. user field is user_id. this can be fetched using /api/users
	myRouter.HandleFunc("/api/users/placeorder", createOrder).Methods("POST")
	//createUser requires json with following structure '{"username":"ritesh", "contact":"8282989812"}'
	myRouter.HandleFunc("/api/users", createUser).Methods("POST")
	myRouter.HandleFunc("/api/users", getUsers)
	myRouter.HandleFunc("/api/users/{user_id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/api/users/{user_id}/orders", getOrdersByUser)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

type orderSummary struct {
	OrderId  int
	ItemName string
	UserId   string
}

type orders struct {
	Orders []orderSummary
}

type userSummary struct {
	UserId        string
	UserName      string
	ContactNumber string
}

type users struct {
	Users []userSummary
}

func createUser(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	var placeOrder map[string]string
	json.Unmarshal(reqBody, &placeOrder)
	createUserStmt := `INSERT into users (user_name, contact_no) values ($1, $2)`
	rows, err := db.Query(createUserStmt, placeOrder["username"], placeOrder["contact"])
	if err != nil {
		log.Println(err)
	}

	getuserStmt := `SELECT user_id, user_name, contact_no FROM users where user_name=$1 and contact_no=$2`
	rows, err = db.Query(getuserStmt, placeOrder["username"], placeOrder["contact"])
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	user := userSummary{}
	for rows.Next() {
		rows.Scan(&user.UserId, &user.UserName, &user.ContactNumber)
	}
	out, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out))
}

func createOrder(w http.ResponseWriter, req *http.Request) {
	orderserviceHost, ok := os.LookupEnv(orderServiceHost)
	if !ok {
		panic("ORDERSERVICEHOST environment variable required but not set")
	}
	orderservicePort, ok := os.LookupEnv(orderServicePort)
	if !ok {
		panic("ORDERSERVICEPORT environment variable required but not set")
	}
	reqBody, _ := ioutil.ReadAll(req.Body)
	var placeOrder map[string]string
	json.Unmarshal(reqBody, &placeOrder)
	jsonbody, _ := json.Marshal(placeOrder)
	responseBody := bytes.NewBuffer(jsonbody)
	resp, err := http.Post("http://"+orderserviceHost+":"+orderservicePort+"/api/orders", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
	fmt.Fprintf(w, sb)
}

func getUsers(w http.ResponseWriter, req *http.Request) {
	userlist := users{}
	query := `SELECT user_id, user_name, contact_no FROM users`
	err := queryUsers(&userlist, query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(userlist)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
}

func queryUsers(userlist *users, query string) error {
	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		user := userSummary{}
		err = rows.Scan(&user.UserId, &user.UserName, &user.ContactNumber)
		if err != nil {
			return err
		}
		userlist.Users = append(userlist.Users, user)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func getOrdersByUser(w http.ResponseWriter, req *http.Request) {
	orderserviceHost, ok := os.LookupEnv(orderServiceHost)
	if !ok {
		panic("ORDERSERVICEHOST environment variable required but not set")
	}
	orderservicePort, ok := os.LookupEnv(orderServicePort)
	if !ok {
		panic("ORDERSERVICEPORT environment variable required but not set")
	}
	vars := mux.Vars(req)
	user_id := vars["user_id"]
	resp, err := http.Get("http://" + orderserviceHost + ":" + orderservicePort + "/api/orders/" + user_id)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
	fmt.Fprintf(w, sb)

}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user_id := vars["user_id"]
	deleteStmt := `DELETE FROM users WHERE user_id=$1`
	_, err := db.Exec(deleteStmt, user_id)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "User id: %v deleted", user_id)
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
