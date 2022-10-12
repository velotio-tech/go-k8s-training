package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// const (
// 	host   = "localhost"
// 	port   = "5432"
// 	user   = "postgres"
// 	pass   = "myPassword"
// 	dbname = "company"
// )

type UserAndItsOrders struct {
	username     string
	orderdetails string
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:myPassword@localhost/company")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/alldata", dataIndex)
	http.HandleFunc("/alldata/show", usershow)
	http.HandleFunc("/alldata/create", usercreate)
	http.HandleFunc("/alldata/delete", user_delete)
	http.HandleFunc("/alldata/update", user_order_update)
	// http.HandleFunc("/books/show", booksShow)
	// http.HandleFunc("/books/create", booksCreate)
	fmt.Println("server running on port no : 8080")
	http.ListenAndServe(":8080", nil)
}

func dataIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Request type is not correct", 405)
		return
	}

	rows, err := db.Query("SELECT * FROM useranditsorders")
	if err != nil {
		http.Error(w, "DB Query Error", 500)
		return
	}
	defer rows.Close()

	uaio := make([]*UserAndItsOrders, 0)
	for rows.Next() {
		uo := new(UserAndItsOrders)
		err := rows.Scan(&uo.username, &uo.orderdetails)
		if err != nil {
			http.Error(w, "DB row has an error", 500)
			return
		}
		uaio = append(uaio, uo)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, "DB row has an error", 500)
		return
	}

	for _, uo := range uaio {
		//fmt.Println(uo.username, uo.orderdetails)
		fmt.Fprintf(w, "%s, %s\n", uo.username, uo.orderdetails)
	}
}

func usershow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Request type is not correct", 405)
		return
	}

	user := r.FormValue("username")
	if user == "" {
		http.Error(w, "username nil", 400)
		return
	}

	row := db.QueryRow("SELECT * FROM useranditsorders WHERE username = $1", user)

	uo := new(UserAndItsOrders)
	err := row.Scan(&uo.username, &uo.orderdetails)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "DB row has an error", 500)
		return
	}

	fmt.Fprintf(w, "%s, %s\n", uo.username, uo.orderdetails)
}

func usercreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	un := r.FormValue("username")
	od := r.FormValue("orderdetails")

	if un == "" || od == "" {
		http.Error(w, "Username and order details required", 400)
		return
	}

	result, err := db.Exec("INSERT INTO useranditsorders VALUES($1, $2)", un, od)
	if err != nil {
		http.Error(w, "DB Query Error", 500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "DB error", 500)
		return
	}

	fmt.Fprintf(w, "useranditsorders created successfully (%d row affected)\n", rowsAffected)
}

func user_order_update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		fmt.Errorf("Method not allowed")
		return
	}

	un := r.FormValue("username")
	od := r.FormValue("orderdetails")

	if un == "" || od == "" {
		http.Error(w, "check user and order", 400)
		fmt.Fprintln(w, "user and order name required")
		return
	}

	result, err := db.Exec("UPDATE useranditsorders SET orderdetails = $1 WHERE username = $2", od, un)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		fmt.Errorf("Internal Error", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		//	fmt.Errorf("Internal Error", err)
		return
	}

	fmt.Fprintf(w, "useranditsorders updated successfully (%d row affected)\n", rowsAffected)

	w.Write([]byte("updated the data\n"))

}

func user_delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		fmt.Errorf("Method not allowed")
		return
	}

	un := r.FormValue("username")

	if un == "" {
		http.Error(w, "check user and order", 400)
		fmt.Fprintln(w, "user and order name required")
		return
	}

	result, err := db.Exec("DELETE FROM useranditsorders WHERE username = $1", un)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		fmt.Errorf("Internal Error", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500)
		log.Println(err)
		//	fmt.Errorf("Internal Error", err)
		return
	}

	fmt.Fprintf(w, "useranditsorders deleted successfully (%d row affected)\n", rowsAffected)

	w.Write([]byte("updated the data\n"))

}
