package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Order struct {
	OrdeID      string `json:"order_id"`
	ProductName string `json:"p_name"`
	Description string `json:"p_desc"`
	Price       int    `json:"price"`
}

func GetUsers(db *sql.DB) []User {
	results, err := db.Query("select * from users")
	if err != nil {
		log.Fatal("error fetching data from users table")
	}
	var users []User
	for results.Next() {
		var (
			user_id  string
			username string
		)
		err = results.Scan(&user_id, &username)
		if err != nil {
			log.Fatal("unable to parse row, ", err)
		}
		users = append(users, User{user_id, username})
	}
	return users
}

func GetOrders(db *sql.DB, user_id string) []Order {
	results, err := db.Query("select o.order_id, p.Name, p.Description, p.Price from orders as o, product as p where p.productID = o.product_id and o.user_id = ?", user_id)
	if err != nil {
		log.Fatal("error fetching data from orders table")
	}

	var Orders []Order
	for results.Next() {
		var (
			Order_id string
			P_name   string
			P_desc   string
			Price    int
		)
		err = results.Scan(&Order_id, &P_name, &P_desc, &Price)
		if err != nil {
			log.Fatal("unable to parse row, ", err)
		}
		Orders = append(Orders, Order{Order_id, P_name, P_desc, Price})
	}
	return Orders
}

func DeleteOrder(db *sql.DB, order_id string) string {

	_, err := db.Query("DELETE FROM orders where order_id = ?", order_id)
	if err != nil {
		return fmt.Sprint("Order could not delete: ", err)
	} else {
		return fmt.Sprintln("Order successfully deleted")
	}
}
func UpdateOrderHelper(db *sql.DB, order_id, updated_product_id string) string {
	_, err := db.Query("UPDATE orders SET product_id = ? where order_id = ?", updated_product_id, order_id)
	if err != nil {
		return fmt.Sprint("Order could not update: ", err)
	} else {
		return fmt.Sprintln("Order successfully updated")
	}
}
