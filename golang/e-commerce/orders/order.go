package main

import (
	"e-commerce/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Order struct {
	Id            int    `json:"id"`
	Product_name  string `json:"product_name"`
	Product_price string `json:"product_price"`
	Ordered_by    string `json:"ordered_by"`
}

type TempOrder struct {
	Id         int `json:"id"`
	Product_id int `json:"product_id"`
	User_id    int `json:"user_id"`
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting request: GET /users/:user_id/orders")

	params := mux.Vars(r)

	db := database.OpenDbConnection()
	sqlQuery := `SELECT orders.id AS order_id, products.name AS product_name, products.price AS product_price, users.name AS ordered_by
				FROM orders INNER JOIN products ON orders.product_id = products.id
				INNER JOIN users ON orders.user_id = users.id
				WHERE user_id = $1`

	rows, err := db.Query(sqlQuery, params["user_id"])
	if err != nil {
		panic(err)
	}

	var orders []Order

	for rows.Next() {
		var id int
		var product_name string
		var product_price string
		var ordered_by string

		err = rows.Scan(&id, &product_name, &product_price, &ordered_by)
		if err != nil {
			panic(err)
		}

		orders = append(orders, Order{Id: id, Product_name: product_name, Product_price: product_price, Ordered_by: ordered_by})
	}

	json.NewEncoder(w).Encode(orders)

	db.Close()
	log.Print("Completed request: GET /users/:user_id/orders")
	log.Print("-------------------------------------------------")
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting request: POST /users/:user_id/orders")

	params := mux.Vars(r)
	var order TempOrder
	err := json.NewDecoder(r.Body).Decode(&order)

	db := database.OpenDbConnection()
	sqlQuery := "INSERT INTO orders (user_id, product_id) values ($1,$2)"
	_, err = db.Exec(sqlQuery, params["user_id"], order.Product_id)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "New order created!")

	db.Close()
	log.Print("Completed request: POST /users/:user_id/orders")
	log.Print("-------------------------------------------------")
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting request: PATCH /users/:user_id/orders/:id")

	params := mux.Vars(r)
	var order TempOrder
	err := json.NewDecoder(r.Body).Decode(&order)

	db := database.OpenDbConnection()
	sqlQuery := "UPDATE orders SET product_id = $1 WHERE id = $2"
	result, err := db.Exec(sqlQuery, order.Product_id, params["id"])
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	} else if rowsAffected == 0 {
		http.Error(w, "Failed to update order!", 412)
	} else {
		fmt.Fprintf(w, "Order updated successfully!")
	}

	db.Close()
	log.Print("Completed request: PATCH /users/:user_id/orders/:id")
	log.Print("-------------------------------------------------")
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	log.Print("Startting request: DELETE /users/:user_id/orders/:id")

	params := mux.Vars(r)

	db := database.OpenDbConnection()
	sqlQuery := "DELETE FROM orders WHERE id = $1"
	result, err := db.Exec(sqlQuery, params["id"])
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	} else if rowsAffected == 0 {
		http.Error(w, "Failed to delete order record!", 412)
	} else {
		fmt.Fprintf(w, "Order deleted successfully!")
	}

	db.Close()
	log.Print("Completed request: DELETE /users/:user_id/orders/:id")
	log.Print("-------------------------------------------------")
}

func DeleteAllOrdersForUser(w http.ResponseWriter, r *http.Request) {
	log.Print("Startting request: DELETE /users/:user_id/orders")

	params := mux.Vars(r)
	db := database.OpenDbConnection()
	sqlQuery := "DELETE FROM orders WHERE user_id = $1"
	_, err := db.Exec(sqlQuery, params["user_id"])
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Orders deleted succesfully!")

	db.Close()
	log.Print("Completed request: DELETE /users/:user_id/orders")
	log.Print("-------------------------------------------------")
}
