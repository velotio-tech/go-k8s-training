package db

import (
	"context"
	"database/sql"
)

// Database ...
type Database interface {
	PingContext(ctx context.Context) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
