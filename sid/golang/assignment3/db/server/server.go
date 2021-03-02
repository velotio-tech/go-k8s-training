package server

import (
	"fmt"
	"net/http"

	"github.com/farkaskid/go-k8s-training/assignment3/db/handlers"
)

func New(port int) *http.Server {
	rootHandler := handlers.RootHandler{}
	rootHandler.Init()

	rootHandler.PathMapping["hello"] = handlers.Hello
	rootHandler.PathMapping["users"] = handlers.UserHandler

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: rootHandler,
	}
}
