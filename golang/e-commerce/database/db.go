package database

import(
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func OpenDbConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@e-commerce-db:5432/e_commerce?sslmode=disable")
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}

	return db
}

