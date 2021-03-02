package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/farkaskid/go-k8s-training/assignment3/db/handlers"
)

func New(port int, db *sql.DB) *http.Server {
	rootHandler := handlers.RootHandler{}
	rootHandler.Init(db)

	rootHandler.PathMapping["users"] = handlers.UserHandler

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: rootHandler,
	}
}
