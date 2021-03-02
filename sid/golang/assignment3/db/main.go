package main

import (
	"database/sql"
	"fmt"

	"github.com/farkaskid/go-k8s-training/assignment3/db/helpers"
	dbserver "github.com/farkaskid/go-k8s-training/assignment3/db/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	helpers.Bootstrap(db)
	dbserver.New(2222, db).ListenAndServe()
}
