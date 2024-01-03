package db

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 64
const maxIdleDbConn = 64
const maxDbLifeTime = 5 * time.Minute

func ConnectDB(desc string) (*DB, error) {

	db, err := NewDatabase(desc)

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = db

	return dbConn, nil
}

func NewDatabase(desc string) (*sql.DB, error) {
	db, err := sql.Open("pgx", desc)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() *sql.DB {
	return dbConn.SQL
}
