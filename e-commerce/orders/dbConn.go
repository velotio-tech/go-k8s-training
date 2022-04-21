package main

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Parav@123@tcp(127.0.0.1:3306)/ecommerce")

	if err != nil {
		log.Fatal("unable to make db connection")
	}
	return db
}

