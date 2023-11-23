package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// QuitRetry is the exit point for retrying from a infinite loop
var QuitRetry = make(chan bool)

// Connect2DB extablish the connection to the database
func Connect2DB(connStr string, dbRetry time.Duration) *sql.DB {
	for {
		select {
		case <-QuitRetry:
			return nil
		case <-time.After(dbRetry):
			DB, err := sql.Open("postgres", connStr)
			if err != nil {
				log.Println("failed to connect to database: ", err)
			} else {
				return DB
			}
			log.Printf("reconnecting to database, after: %v milliseconds\n", dbRetry)
		}
	}
}
