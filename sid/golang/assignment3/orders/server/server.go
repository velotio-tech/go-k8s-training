package server

import (
	"fmt"
	"net/http"

	"github.com/farkaskid/go-k8s-training/assignment3/orders/handlers"
)

func New(port int) *http.Server {
	rootHandler := handlers.RootHandler{}
	rootHandler.Init()

	rootHandler.PathMapping["orders"] = handlers.OrderHandler

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: rootHandler,
	}
}
