package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type RootHandler struct {
	PathMapping map[string]func(http.ResponseWriter, *http.Request, *sql.DB)
	DB          *sql.DB
}

func (handler *RootHandler) Init(db *sql.DB) {
	handler.PathMapping = make(map[string]func(http.ResponseWriter, *http.Request, *sql.DB))
	handler.DB = db
}

func (handler RootHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Request:", req.Method, "on", req.URL)

	pathChunks := strings.Split(strings.TrimSpace(req.URL.Path), "/")

	handlerFunc, present := handler.PathMapping[pathChunks[1]]

	if !present {
		fmt.Println("Path", req.URL.Path, "not found")
		http.NotFound(resp, req)
		return
	}

	handlerFunc(resp, req, handler.DB)
}
