package server

import (
	"net/http"
	"strconv"

	"github.com/farkaskid/go-k8s-training/assignment3/db/handlers"
)

func New(port int) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handlers.Hello)
	mux.HandleFunc("/users", handlers.CreateUserHandler)

	return &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}
}
