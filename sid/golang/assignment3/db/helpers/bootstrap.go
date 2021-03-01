package helpers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Bootstrap() {
	os.Remove("./main.db")

	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	// create users table
	sqlStmt := `create table users (id integer not null primary key, name text); delete from users`

	_, err = db.Exec(sqlStmt)

	if err != nil {
		fmt.Println("Failed to create users table")
	}

	// create orders table
	sqlStmt = `create table orders (id integer not null primary key, details text); delete from orders`

	_, err = db.Exec(sqlStmt)

	if err != nil {
		fmt.Println("Failed to create orders table")
	}

	// create user_order_mapping table
	sqlStmt = `create table users_orders_mapping (
		id integer not null primary key, user_id integer not null, order_id integer not null); delete from users_orders_mapping`

	_, err = db.Exec(sqlStmt)

	if err != nil {
		fmt.Println("Failed to create user_order_mapping table")
	}
}
