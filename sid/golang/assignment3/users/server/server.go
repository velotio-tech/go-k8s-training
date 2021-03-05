package server

import (
	"fmt"
	"net/http"

	"github.com/farkaskid/go-k8s-training/assignment3/users/handlers"
)

func New(port int) *http.Server {
	rootHandler := handlers.RootHandler{}
	rootHandler.Init()

	rootHandler.PathMapping["users"] = handlers.UserHandler
	rootHandler.PathMapping["orders"] = handlers.OrderHandler

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: rootHandler,
	}
}
