package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/farkaskid/go-k8s-training/assignment3/db/helpers"
	dbserver "github.com/farkaskid/go-k8s-training/assignment3/db/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, err := os.Stat("./database/main.db")
	fmt.Println("Database exists", err)
	dbExists := err == nil

	db, err := sql.Open("sqlite3", "./database/main.db")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	if !dbExists {
		helpers.Bootstrap(db)
	}

	dbserver.New(2222, db).ListenAndServe()
}
